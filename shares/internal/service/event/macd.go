package event

type Macd struct {
	DayStr string
	Dif    float64
	Dea    float64
	Macd   float64
}

// var macd=calcMACD(12,26,9,data.datas,1);

/*
 * 计算MACD
 * @param {number} short 快速EMA时间窗口
 * @param {number} long 慢速EMA时间窗口
 * @param {number} mid dea时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
func calcMACD(short, long, mid int, data [][]interface{}, field int) (macds []Macd) {
	dif := calcDIF(short, long, data, field)
	dea := calcDEA(mid, dif)
	for i := 0; i < len(data); i++ {
		macds = append(macds, Macd{
			Dif:    dif[i],
			Dea:    dea[i],
			DayStr: data[i][0].(string),
			Macd:   toFixed((dif[i] - dea[i]) * 2),
		})
	}

	return
}

/*
 * 计算DIF快线，用于MACD
 * @param {number} short 快速EMA时间窗口
 * @param {number} long 慢速EMA时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
func calcDIF(short, long int, data [][]interface{}, field int) (dif []float64) {
	emaShort := calcEMA(short, data, field)
	emaLong := calcEMA(long, data, field)
	for i := 0; i < len(data); i++ {
		dif = append(dif, toFixed(emaShort[i]-emaLong[i])) // toFixed(3));
	}
	return dif
}

/*
 * 计算DEA慢线，用于MACD
 * @param {number} mid 对dif的时间窗口
 * @param {array} dif 输入数据
 */
func calcDEA(mid int, dif []float64) []float64 {
	return calcEMAEx(mid, dif)
}

/*
 * 计算EMA指数平滑移动平均线，用于MACD
 * @param {number} n 时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
func calcEMA(n int, data [][]interface{}, field int) []float64 {
	var ema []float64
	a := (2.0 / (float64(n) + 1.0))
	//二维数组
	ema = append(ema, data[0][field].(float64))
	for i := 1; i < len(data); i++ {
		ema = append(ema, toFixed(((a * data[i][field].(float64)) + ((1 - a) * ema[i-1]))))
	}
	return ema
}

/*
 * 计算EMA指数平滑移动平均线，用于MACD
 * @param {number} n 时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
func calcEMAEx(n int, data []float64) []float64 {
	var ema []float64
	a := (2.0 / (float64(n) + 1.0))
	//普通一维数组
	ema = append(ema, data[0])
	for i := 1; i < len(data); i++ {
		ema = append(ema, toFixed(a*data[i]+(1-a)*ema[i-1]))
	}

	return ema
}

func toFixed(data float64) float64 {
	var tmp int64 = int64((data + 0.0004) * 1000)
	return float64(tmp) * 0.001
}
