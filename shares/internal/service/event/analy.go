package event

import (
	proto "shares/rpc/shares"
)

const (
	_maxSpan        = 3 * (60 / _timeoutTicker) // 三分钟内涨跌幅
	_maxUpPercent   = 3                         // 快速涨幅 3%
	_maxDownPercent = 4                         // 快速跌幅 4%
)

var gUpDownTicket int
var gmaTicket int
var _analy Analy

type priceInfo struct {
	now   int64
	price float64
}

type Analy struct {
	mp   map[string][]priceInfo
	span int //  跨度
	cap  int // 报警容量
}

func (a *Analy) Init() { //
	a.mp = make(map[string][]priceInfo)
	a.span = _maxSpan
	a.cap = a.span * 2
}

func (a *Analy) Add(info *proto.SharesInfo, now int64) {
	a.mp[info.Code] = append(a.mp[info.Code], priceInfo{
		now:   now,
		price: info.Price,
	})

	if len(a.mp[info.Code]) > a.cap { // 去掉不用的
		a.mp[info.Code] = a.mp[info.Code][(len(a.mp[info.Code]) - a.span):]
	}
}

func (a *Analy) StartAnaly(list []*proto.SharesInfo) { // 开始分析
	gUpDownTicket++
	gmaTicket++

	for _, v := range list {
		if len(a.mp[v.Code]) > a.span { // 找最大最小的
			offset := len(a.mp[v.Code]) - a.span
			minV := a.mp[v.Code][offset]
			maxV := minV
			for i := offset; i < len(a.mp[v.Code]); i++ {
				tmp := a.mp[v.Code][i]
				if minV.price > tmp.price {
					minV = tmp
				}
				if maxV.price < tmp.price {
					maxV = tmp
				}
			}

			// 开始比较
			percent := ((maxV.price - minV.price) / minV.price) * 100 // 这是个正数
			if percent >= _maxUpPercent {                             // 快速涨跌提醒
				if maxV.now < minV.now { // 说明是跌幅
					if percent >= _maxDownPercent {
						surgeSlump(v.Code, maxV.price, minV.price, -percent, minV.now-maxV.now)
					}
				} else { // 涨幅
					surgeSlump(v.Code, minV.price, maxV.price, percent, maxV.now-minV.now)
				}
			}
		}

		if gUpDownTicket > 10 {
			upDown(v.Code, v.Price, v.Percent) // 估价涨跌提醒
		}

		if gmaTicket > 200 {
			maRealTime(v.Code, v.Price, v.Percent) // 日均线实时提醒
		}
	}

	if gUpDownTicket > 10 {
		gUpDownTicket = 0
	}
	if gmaTicket > 200 {
		gmaTicket = 0
	}

}
