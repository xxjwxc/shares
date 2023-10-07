package analy

import (
	"math"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 横盘 (线性回归)
func hengpan() {
	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())

	malist, err := model.AnalyHpTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyHpTblMgr(orm.DB).Delete(&v)
		} else {
			v.Percent = sharesInfo.Percent
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyHpTblMgr(orm.DB).Save(&v) // 保留
		}
	}

	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	// codes = append(codes, "sz002085", "sh600777")
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 <= ?", day0).Order("day0 desc").Limit(30)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 30 { // 不需要
			continue
		}

		var maxNum, minNum float64
		_len := len(list)
		var points []Point
		for i := len(list) - 1; i >= 0; i-- {
			points = append(points, Point{
				X: float64(_len - i),
				Y: list[i].Price,
			})

			if minNum == 0 || list[i].Price < minNum {
				minNum = list[i].Price
			}
			if maxNum == 0 || list[i].Price > maxNum {
				maxNum = list[i].Price
			}
		}

		_, p0 := linearRegressionLSE(points)

		var jp0 float64
		for _, v := range points { // 计算波动率
			jp0 += math.Sqrt(math.Pow((p0.X*v.X)-(1*v.Y)+p0.Y, 2) / (math.Pow(v.X, 2) + math.Pow(v.Y, 2)))
		}
		jp0 = (jp0 / float64(len(points))) * 100

		// a = math.Pi / (math.Atan(a) * 180)
		if p0.X < 0.001 && p0.X > -0.001 && jp0 < 0.5 && (maxNum-minNum) < (minNum*0.1) {
			sharesInfo, err := event.GetShare(code, true)
			if err != nil {
				continue
			}
			tmp, _ := model.AnalyHpTblMgr(orm.DB).GetFromCode(code)
			if tmp.ID == 0 {
				tmp.Code = code
				tmp.Day = datatypes.Date(time.Now())
				tmp.DayStr = list[0].Day0Str
				tmp.CreatedAt = time.Now()
				tmp.Price = sharesInfo.Price
			}
			tmp.Name = sharesInfo.Name
			tmp.Percent = sharesInfo.Percent
			tmp.TurnoverRate = list[0].TurnoverRate
			tmp.Score = p0.X * 1000 // 100分
			model.AnalyHpTblMgr(orm.DB).Save(&tmp)
			// fmt.Println(code, a, b)
		}
	}
}

func LeastSquares(x []float64, y []float64) (a float64, b float64) {
	// x是横坐标数据,y是纵坐标数据
	// a是斜率，b是截距
	xi := float64(0)
	x2 := float64(0)
	yi := float64(0)
	xy := float64(0)

	if len(x) != len(y) {
		mylog.Debug("最小二乘时，两数组长度不一致!")
	} else {
		length := float64(len(x))
		for i := 0; i < len(x); i++ {
			xi += x[i]
			x2 += x[i] * x[i]
			yi += y[i]
			xy += x[i] * y[i]
		}
		a = (yi*xi - xy*length) / (xi*xi - x2*length)
		b = (yi*x2 - xy*xi) / (x2*length - xi*xi)
	}
	return
}

type Point struct {
	X float64
	Y float64
}

func linearRegressionLSE(series []Point) ([]Point, Point) {

	q := len(series)

	if q == 0 {
		return make([]Point, 0, 0), Point{}
	}

	p := float64(q)

	sum_x, sum_y, sum_xx, sum_xy := 0.0, 0.0, 0.0, 0.0

	for _, p := range series {
		sum_x += p.X
		sum_y += p.Y
		sum_xx += p.X * p.X
		sum_xy += p.X * p.Y
	}

	m := (p*sum_xy - sum_x*sum_y) / (p*sum_xx - sum_x*sum_x)
	b := (sum_y / p) - (m * sum_x / p)

	r := make([]Point, q, q)

	for i, p := range series {
		r[i] = Point{p.X, (p.X*m + b)}
	}

	return r, Point{X: m, Y: b}
}
