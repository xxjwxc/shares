package analy

import (
	"shares/internal/config"
)

func init() {
	// watchMyself()
	// maDaily()
	// tkgk() // 跳空提醒
	// updatePeg()
	// watchZTB()
	// dwfl()
	// hengpan()
	// watchFl()
	// fmt.Println(AnalyXielv("sh600104"))
	// hengpan()
	// initPinyin() // 初始化拼音
	// initMenu() // 初始化菜单
	// initHYZLJLR()
	// countHYZLJLR(tools.GetUtcDay0(time.Now()))
	if config.GetIsTools() > 0 {
		// getMinuteFangLiang()
		// initHY()      // 初始化板块信息
		// initHYZLJLR() // 更新所有行业净流入
		// initHYDaily() // 更新所有主力净流入
		startTools()
	}
}

func startTools() {
	if config.GetIsTools() == 4 {
		watchLocalFl()
		return
	}
	if config.GetIsTools() == 2 {
		// analyAllMsg() // 分析统计每日消息
		initMsg()
		return
	}
	if config.GetIsTools() == 3 {
		fangLiang() // 放量
		return
	}

	// initDayDailyLHB() //初始化龙虎榜
	// updateMacd()
	// initMacd()
	hengpan()     // 很盘
	dbSzx()       // 底部缩量十字星
	maLhb()       // 龙虎榜
	fangLiang()   // 放量
	chaoDie()     // 超跌
	initmaDaily() // 趋势

	// updateBS()    //  更新北上
	toolsTicket() // 监听
}
