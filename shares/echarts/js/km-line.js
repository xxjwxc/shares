var bgColor = "#100C2A";//背景
var upColor= "#00aa3b";//"#F9293E";//涨颜色
var downColor= "#F9293E"; // "#00aa3b";//跌颜色

// ma  颜色
// var ma5Color = "#39afe6";
// var ma10Color = "#da6ee8";
// var ma20Color = "#ffab42";
// var ma30Color = "#00940b";

var ma5Color = "#ffffff";
var ma10Color = "#feb71d";
var ma20Color = "#da6ee8";
var ma30Color = "#00940b";

var txtColor = "#5b91eb"; // 文本颜色

function setColor(up,down){ 
	this.upColor = up;
	this.downColor = down;
}
/**
 * 15:20 时:分 格式时间增加num分钟
 * @param {Object} time 起始时间
 * @param {Object} num
 */
function addTimeStr(time,num){ 
	var hour=time.split(':')[0];
	var mins=Number(time.split(':')[1]);
	var mins_un=parseInt((mins+num)/60);
	var hour_un=parseInt((Number(hour)+mins_un)/24);
	if(mins_un>0){
		if(hour_un>0){
			var tmpVal=((Number(hour)+mins_un)%24)+"";
			hour=tmpVal.length>1? tmpVal:'0'+tmpVal;//判断是否是一位
		}else{
			var tmpVal=Number(hour)+mins_un+"";
			hour=tmpVal.length>1? tmpVal:'0'+tmpVal;
		}
		var tmpMinsVal=((mins+num)%60)+"";
		mins=tmpMinsVal.length>1? tmpMinsVal:0+tmpMinsVal;//分钟数为 取余60的数
	}else{
		var tmpMinsVal=mins+num+"";
		mins=tmpMinsVal.length>1? tmpMinsVal:'0'+tmpMinsVal;//不大于整除60
	} 
	return hour+":"+mins;
}

//获取增加指定分钟数的 数组  如 09:30增加2分钟  结果为 ['09:31','09:32'] 
function getNextTime(startTime, endTIme, offset, resultArr) {
	var result = addTimeStr(startTime, offset);
	resultArr.push(result);
	if (result == endTIme) {
		return resultArr;
	} else {
		return getNextTime(result, endTIme, offset, resultArr);
	}
}


/**
 * 不同类型的股票的交易时间会不同  
 * @param {Object} type   hs=沪深  us=美股  hk=港股
 */
var time_arr = function(type) { 
	if(type.indexOf('us')!=-1){//生成美股时间段
		var timeArr = new Array();
		timeArr.push('09:30')
		return getNextTime('09:30', '16:00', 1, timeArr);
	}
	if(type.indexOf('hs')!=-1){//生成沪深时间段
		var timeArr = new Array();
			timeArr.push('09:30');
			timeArr.concat(getNextTime('09:30', '11:30', 1, timeArr)); 
			timeArr.push('13:00');
			timeArr.concat(getNextTime('13:00', '15:00', 1, timeArr)); 
		return timeArr;
	}
	if(type.indexOf('hk')!=-1){//生成港股时间段
		var timeArr = new Array();
			timeArr.push('09:30');
			timeArr.concat(getNextTime('09:30', '11:59', 1, timeArr)); 
			timeArr.push('13:00');
			timeArr.concat(getNextTime('13:00', '16:00', 1, timeArr)); 
		return timeArr;
	}
	
}


var get_m_data = function(m_data,type) {
	var priceArr = new Array();
	var avgPrice = new Array();
	var vol = new Array();
	var times = time_arr(type); 
	m_data.data.forEach(function(v,i) {
		priceArr.push(v[1]);
		avgPrice.push(v[2]);
		vol.push(v[3]); 
	})
	return {
		priceArr: priceArr,
		avgPrice: avgPrice,
		vol: vol,
		times: times
	}
}



//==========================================分时表 option

/**
 * 生成分时option 
 * @param {Object} m_data 分时数据
 * @param {Object} type 股票类型  us-美股  hs-沪深  hk-港股
 */
function initMOption(m_data,type){
	var m_datas = get_m_data(m_data,type); 
	return {
		animation: false,
		tooltip: { //弹框指示器
			trigger: 'axis',
			axisPointer: {
				type: 'cross'
			},
			formatter: function(params, ticket, callback) {
				var i = params[0].dataIndex;
	
				var color;
				if (m_datas.priceArr[i] > m_data.yestclose) {
					color ='style="color:'+upColor+'"';
				} else {
					color ='style="color:'+downColor+'"'; 
				}
	
				var html = '<div class="commColor" style="width:100px;"><div>当前价 <span  '+color+' >' + m_datas.priceArr[i] + '</span></div>';
				html += '<div>均价 <span  '+color+' >' + m_datas.avgPrice[i] + '</span></div>';
				html += '<div>涨幅 <span  '+color+' >' + ratioCalculate(m_datas.priceArr[i],m_data.yestclose)+ '%</span></div>';
				html += '<div>成交量 <span  '+color+' >' + m_datas.vol[i] + '</span></div></div>'
				return html;
			}
		},
		legend: { //图例控件,点击图例控制哪些系列不显示
			icon: 'rect',
			type: 'scroll',
			itemWidth: 14,
			itemHeight: 2,
			left: 0,
			top: '0%', // 控件位置 
			textStyle: {
				fontSize: 12,
				color: '#0e99e2'
			}
		},
		axisPointer: {
			link: [
				{
				  xAxisIndex: 'all'
				}
			  ],
			  label: {
				backgroundColor: '#777'
			  }
		},
		color: [ma5Color, ma10Color],
		grid: [{
				id: 'gd1',
				left: '0%',
				right: '1%',
				height: '67.5%', //主K线的高度,
				top: '5%'
			},
			{
				id: 'gd2',
				left: '0%',
				right: '1%',
				height: '67.5%', //主K线的高度,
				top: '5%'
			}, {
				id: 'gd3',
				left: '0%',
				right: '1%',
				top: '75%',
				height: '19%' //交易量图的高度
			}
		],
		xAxis: [ //==== x轴
			{ //主图 
				gridIndex: 0,
				data: m_datas.times,
				axisLabel: { //label文字设置
					show: false
				},
				splitLine: {
					show: false,
				}
			},
			{
				show:false,
				gridIndex: 1,
				data: m_datas.times,
				axisLabel: { //label文字设置
					show: false
				},
				splitLine: {
					show: false,
				}
			}, { //交易量图
				splitNumber: 2,
				type: 'category',
				gridIndex: 2,
				data: m_datas.times,
				axisLabel: { //label文字设置
					color: '#9b9da9',
					fontSize: 10
				},
			}
		],
		yAxis: [ //y轴
			{
				gridIndex: 0,
				scale: true, 
				splitNumber: 3,  
				axisLabel: { //label文字设置 
					inside: true, //label文字朝内对齐 
					fontWeight:'bold',
					color:function(val){ 
						if(val==m_data.yestclose){
							return '#aaa'
						}
						return val>m_data.yestclose? upColor:downColor;
					}
				},z:4,splitLine: { //分割线设置
					show: false,
					lineStyle: {
						color: '#181a23'
					}
				},  
			}, { 
				scale: true,  gridIndex: 1, splitNumber: 3,  
				position: 'right', z:4,
				axisLabel: { //label文字设置
					color:function(val){ 
						if(val==m_data.yestclose){
							return '#aaa'
						}
						return val>m_data.yestclose? upColor:downColor; 
					},
					inside: true, //label文字朝内对齐 
					fontWeight:'bold',
					formatter:function(val){
						var resul=ratioCalculate(val,m_data.yestclose);
						return  Number(resul).toFixed(2)+' %'}
				},
				splitLine: { //分割线设置
					show: false,
					lineStyle: {
						color: '#181a23'
					}
				},
				axisPointer:{show:true,
					label:{
						formatter:function(params){ //计算右边Y轴对应的当前价的涨幅比例
							return  ratioCalculate(params.value,m_data.yestclose)+'%';
						}
					}
				}
			} 
			, { //交易图
				gridIndex: 2,z:4,
				splitNumber: 3,
				axisLine: {
					onZero: false
				},
				axisTick: {
					show: false
				},
				splitLine: {
					show: false
				},
				axisLabel: { //label文字设置
					color: '#c7c7c7',
					inside: true, //label文字朝内对齐 
					fontSize: 8
				},
			}
		],
		dataZoom: [
	
		],
		animation:true,//禁止动画效果
		backgroundColor: bgColor,
		blendMode: 'source-over',
		series: [
			 {
				name: '当前价',
				type: 'line',
				data: m_datas.priceArr,
				smooth: true,
				symbol: "circle", //中时有小圆点 
				lineStyle: {
					normal: {
						opacity: 0.8,
						color: '#39afe6',
						width: 1,
						margin:2,
					}
				},
				areaStyle: {
					normal: {
						color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
							offset: 0,
							color: 'rgba(0, 136, 212, 0.7)'
						}, {
							offset: 0.8,
							color: 'rgba(0, 136, 212, 0.02)'
						}], false),
						shadowColor: 'rgba(0, 0, 0, 0.1)',
						shadowBlur: 10
					}
				},markLine: {// 昨日收盘价
					label:{ normal:{position:'end',padding: [0, 0, 0, -50] } },
					symbol: ['none', 'none'],
					itemStyle:{
						normal:{
							lineStyle:{
								// width:2,
								type:'dotted'  //'dotted'虚线 'solid'实线
							}
						}
					},
					data: [
					  {
						yAxis: m_data.yestclose,
						itemStyle: {
							color: '#aaa'
						  }
					  }
					]
				  }
			},
			{
				name: '均价',
				type: 'line',
				data: m_datas.avgPrice,
				smooth: true,
				symbol: "circle",
				lineStyle: { //标线的样式
					normal: {
						opacity: 0.8,
						color: '#da6ee8',
						width: 1
					}
				}
			},{  
				type: 'line', 
				data: m_datas.priceArr,
				smooth: true,
				symbol: "none",
				gridIndex: 1,
				xAxisIndex: 1,
				yAxisIndex: 1,
				lineStyle: { //标线的样式 
					normal: { 
						width: 0
					}
				}
			},
			{
				name: '成交量',
				type: 'bar',
				gridIndex: 2,
				xAxisIndex: 2,
				yAxisIndex: 2,
				data: m_datas.vol,
				barWidth: '60%',
				itemStyle: {
					normal: {
						color: function(params) {
							var colorList;
							if (m_datas.priceArr[params.dataIndex] > m_datas.priceArr[params.dataIndex-1]) {
								colorList = upColor;
							} else {
								colorList = downColor;
							}
							return colorList;
						},
					}
				}
			}
		]
	};
}



/**
 * 计算价格涨跌幅百分比
 * @param {Object} price 当前价
 * @param {Object} yclose 昨收价
 */
function ratioCalculate(price,yclose){
	return ((price-yclose)/yclose*100).toFixed(3);
}









//数组处理
function splitData(rawData) {
	var datas = []; var times = [];var vols = []; 
	for (var i = 0; i < rawData.length; i++) {
		datas.push(rawData[i]);
		times.push(rawData[i].splice(0, 1)[0]);
		vols.push(rawData[i][4]); 
	}
	return {datas:datas,times:times,vols:vols};
}

//================================MA计算公式
function getKeyLegend(only20) {
	if (only20 ){// 只显示20日线
		return ['日K', 'MA20','成交量','MACD','DIF','DEA']
	}

	return ['日K', 'MA5', 'MA10', 'MA20', 'MA30','成交量','MACD','DIF','DEA']
}


//================================MA计算公式
function calculateMA(dayCount,only20,data) {
	var result = [];
	if (only20 && dayCount != 20){// 只显示20日线
		return result
	}

	for (var i = 0, len = data.times.length; i < len; i++) {
		if (i < dayCount) {
			result.push('-');
			continue;
		}
		var sum = 0;
		for (var j = 0; j < dayCount; j++) {
			sum += data.datas[i - j][1];
		}
		result.push((sum / dayCount).toFixed(2));
	}
	return result;
}


//=================================================MADC计算公式

var calcEMA,calcDIF,calcDEA,calcMACD;

/*
 * 计算EMA指数平滑移动平均线，用于MACD
 * @param {number} n 时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
calcEMA=function(n,data,field){
    var i,l,ema,a;
    a=2/(n+1);
    if(field){
        //二维数组
        ema=[data[0][field]];  
        for(i=1,l=data.length;i<l;i++){
            ema.push((a*data[i][field]+(1-a)*ema[i-1]).toFixed(2));
        }
    }else{
        //普通一维数组
        ema=[data[0]];
        for(i=1,l=data.length;i<l;i++){
            ema.push((a*data[i]+(1-a)*ema[i-1]).toFixed(3) );
        }
    } 
    return ema;
};

/*
 * 计算DIF快线，用于MACD
 * @param {number} short 快速EMA时间窗口
 * @param {number} long 慢速EMA时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
calcDIF=function(short,long,data,field){
    var i,l,dif,emaShort,emaLong;
    dif=[];
    emaShort=calcEMA(short,data,field);
    emaLong=calcEMA(long,data,field);
    for(i=0,l=data.length;i<l;i++){
        dif.push((emaShort[i]-emaLong[i]).toFixed(3));
    }
    return dif;
};

/*
 * 计算DEA慢线，用于MACD
 * @param {number} mid 对dif的时间窗口
 * @param {array} dif 输入数据
 */
calcDEA=function(mid,dif){
    return calcEMA(mid,dif);
};

/*
 * 计算MACD
 * @param {number} short 快速EMA时间窗口
 * @param {number} long 慢速EMA时间窗口
 * @param {number} mid dea时间窗口
 * @param {array} data 输入数据
 * @param {string} field 计算字段配置
 */
calcMACD=function(short,long,mid,data,field){
    var i,l,dif,dea,macd,result;
    result={};
    macd=[];
    dif=calcDIF(short,long,data,field);
    dea=calcDEA(mid,dif);
    for(i=0,l=data.length;i<l;i++){
        macd.push(((dif[i]-dea[i])*2).toFixed(3));
    }
    result.dif=dif;
    result.dea=dea;
    result.macd=macd;
    return result;
};
 
 
 //=================================================MADC计算公式 end


function initKOption(cdata,only20){
	var data = splitData(cdata);
	var macd=calcMACD(12,26,9,data.datas,1);  
	return {
			tooltip: { //弹框指示器
				trigger: 'axis',
				backgroundColor:'rgba(255,255,255,0.8)',//通过设置rgba调节背景颜色与透明度
				axisPointer: {
					type: 'cross'
				},
				formatter: function (params) {//修改鼠标划过显示为中文
					var out = "";
					var index = 0;
					if (params[0].componentSubType === 'candlestick'){
						out += params[0].name + '<br>' +
						params[0].marker +  '开盘:&nbsp' + params[0].data[1] + '<br>' +
						params[0].marker +  '收盘:&nbsp' + params[0].data[2] + '<br>' +
						params[0].marker + 	'最低:&nbsp' + params[0].data[3] + '<br>' +
						params[0].marker + 	'最高:&nbsp' + params[0].data[4] + '<br>' 

						index = 1;
						// params.forEach(function(v,i) {
						// 	if (i > 0){
						// 		out += v.marker + v.seriesName +'&nbsp '+ v.data + '</br>';
						// 	}	
						// })
					}else{
						out += params[0].marker + params[0].seriesName +'&nbsp '+ params[0].data + '</br>';
						index = 1;
					}

					for(var i=index;i<params.length;i++){
						if (params[i].componentSubType !== 'line'){
							break
						}
						out += params[i].marker + params[i].seriesName +'&nbsp '+ params[i].data + '</br>';
					}
					
					return out;
				}
			},
			legend: { // 图例控件,点击图例控制哪些系列不显示
				data: getKeyLegend(only20),
				icon: 'rect', 
				type:'scroll',
				itemWidth: 14,
				itemHeight: 2,
				left: 0,
				top: '0%',  
				animation:true,
				textStyle: {
					fontSize: 12,
					color: '#0e99e2'
				},
				pageIconColor:'#0e99e2'
			},
			axisPointer: {
				link: [
					{
					  xAxisIndex: 'all'
					}
				  ],
				  label: {
					backgroundColor: '#777'
				  }
			},
			color: [ma5Color, ma10Color, ma20Color, ma30Color],
			grid: [{
				id: 'gd1',
				left: '0%',
				right: '1%',
				height: '60%', //主K线的高度,
				top: '5%'
			}, {
				left: '0%',
				right: '1%',
				top: '66.5%',
				height: '10%' //交易量图的高度
			}, {
				left: '0%',
				right: '1%',
				top: '80%', //MACD 指标
				height: '12%'
			}],
			xAxis: [ //==== x轴
				{ //主图
					type: 'category',
					data: data.times,
					scale: true,
					boundaryGap: false,
					axisLine: {
						onZero: false
					},
					axisLabel: { //label文字设置
						show: false
					},
					splitLine: {
						show: false,
						lineStyle: {
							color: '#3a3a3e'
						}
					},
					splitNumber: 20,
					min: 'dataMin',
					max: 'dataMax'
				}, { //交易量图
					type: 'category',
					gridIndex: 1,
					data: data.times,
					axisLabel: { //label文字设置
						color: '#9b9da9',
						fontSize: 10
					},
				}, { //幅图
					type: 'category',
					gridIndex: 2,
					data: data.times,
					axisLabel: {
						show: false
					}
				}
			],
			yAxis: [ //y轴
				{ //==主图
					scale: true,
					z:4,
					axisLabel: { //label文字设置
						color: '#c7c7c7',
						inside: true, //label文字朝内对齐
					},
					splitLine: { //分割线设置
						show: false,
						lineStyle: {
							color: '#181a23'
						}
					},
				}, { //交易图
					gridIndex: 1, splitNumber: 3, z:4,
					axisLine: {
						onZero: false
					},
					axisTick: {
						show: false
					},
					splitLine: {
						show: false
					},
					axisLabel: { //label文字设置
						color: '#c7c7c7',
						inside: true, //label文字朝内对齐 
						fontSize: 8
					},
				}, { //幅图
					z:4, gridIndex: 2,splitNumber: 4,
					axisLine: {
						onZero: false
					},
					axisTick: {
						show: false
					},
					splitLine: {
						show: false
					},
					axisLabel: { //label文字设置
						color: '#c7c7c7',
						inside: true, //label文字朝内对齐 
						fontSize: 8
					},
				}
			],
			dataZoom: [{
					type: 'slider',
					xAxisIndex: [0, 1, 2], //控件联动
					start: 100,
					end: 80,
					throttle: 10,
					top: '92%',
					height: '8%',
					borderColor: '#696969',
					textStyle: {
						color: '#dcdcdc'
					},
					handleSize: '90%', //滑块图标
					// handleIcon: 'M10.7,11.9v-1.3H9.3v1.3c-4.9,0.3-8.8,4.4-8.8,9.4c0,5,3.9,9.1,8.8,9.4v1.3h1.3v-1.3c4.9-0.3,8.8-4.4,8.8-9.4C19.5,16.3,15.6,12.2,10.7,11.9z M13.3,24.4H6.7V23h6.6V24.4z M13.3,19.6H6.7v-1.4h6.6V19.6z',
					handleIcon:'M0,0 v9.7h5 v-9.7h-5 Z',
					dataBackground: {
						lineStyle: {
							color: '#fff'
						}, //数据边界线样式
						areaStyle: {
							color: '#696969'
						} //数据域填充样式
					}
				},
				// {
				// 	type: 'inside',
				// 	xAxisIndex: [0,1,2],//控件联动
				// },
			],
			// animation: false, //禁止动画效果
			backgroundColor: bgColor,
			blendMode: 'source-over',
			series: [{
					name: '日K',
					type: 'candlestick',
					data: data.datas,
					barWidth: '55%',
					large: true,
					largeThreshold: 100,
					itemStyle: {
						normal: {
							color: upColor, //fd2e2e  ff4242
							color0: downColor,
							borderColor: upColor,
							borderColor0: downColor,
							opacity:0.8
						}
					},
					markPoint: {
						itemStyle:{
							normal:{
						   label:{ 
									show: true,  
									color: txtColor,//气泡中字体颜色
								}
							}
						},
						data: [
						  {
							symbolSize: [0, 0],
							name: 'highest value',
							type: 'max',
							valueDim: 'highest',
						  },
						  {
							symbolSize: [0, 0],
							name: 'lowest value',
							type: 'min',
							valueDim: 'lowest',
						  }
						]
					  },
					//   markLine: {// 最大最小线
					// 	label:{ normal:{ position:'end',padding: [0, 0, 0, -50] } },
					// 	symbol: ['none', 'none'],
					// 	data: [
					// 	  {
					// 		name: 'min line on close',
					// 		type: 'min',
					// 		valueDim: 'close',
					// 		itemStyle: {
					// 			color: upColor
					// 		  }
					// 	  },
					// 	  {
					// 		name: 'max line on close',
					// 		type: 'max',
					// 		valueDim: 'close',
					// 		itemStyle: {
					// 			color: upColor
					// 		  }
					// 	  }
					// 	]
					//   }
				}, {
					name: 'MA5',
					type: 'line',
					data: calculateMA(5,only20,data),
					smooth: true,
					symbol: "none", //隐藏选中时有小圆点
					lineStyle: {
						normal: {
							opacity: 0.8,
							color: ma5Color,
							width: 1
						}
					},
				},
				{
					name: 'MA10',
					type: 'line',
					data: calculateMA(10,only20,data),
					smooth: true,
					symbol: "none",
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.8,
							color: ma10Color,
							width: 1
						}
					}
				},
				{
					name: 'MA20',
					type: 'line',
					data: calculateMA(20,only20,data),
					smooth: true,
					symbol: "none",
					lineStyle: {
						opacity: 0.8,
						width: 1,
						color: ma20Color
					}
				},
				{
					name: 'MA30',
					type: 'line',
					data: calculateMA(30,only20,data),
					smooth: true,
					symbol: "none",
					lineStyle: {
						normal: {
							opacity: 0.8,
							width: 1,
							color: ma30Color
						}
					}
				}, {
					name: '成交量',
					type: 'bar',
					xAxisIndex: 1,
					yAxisIndex: 1,
					data: data.vols,
					barWidth: '60%',
					itemStyle: {
						normal: {
							color: function(params) {
								var colorList;
								if (data.datas[params.dataIndex][1] > data.datas[params.dataIndex][0]) {
									colorList = upColor;
								} else {
									colorList = downColor;
								}
								return colorList;
							},
						}
					}
				}, {
					name: 'MACD',
					type: 'bar',
					xAxisIndex: 2,
					yAxisIndex: 2,
					data: macd.macd,
					barWidth: '40%',
					itemStyle: {
						normal: {
							color: function(params) {
								var colorList;
								if (params.data >= 0) {
									colorList = upColor;
								} else {
									colorList = downColor;
								}
								return colorList;
							},
						}
					}
				}, {
					name: 'DIF',
					type: 'line',
					symbol: "none",
					xAxisIndex: 2,
					yAxisIndex: 2,
					data: macd.dif,
					lineStyle: {
						normal: {
							color: '#ffffff',
							width: 1
						}
					}
				}, {
					name: 'DEA',
					type: 'line',
					symbol: "none",
					xAxisIndex: 2,
					yAxisIndex: 2,
					data: macd.dea,
					lineStyle: {
						normal: {
							opacity: 0.8,
							color: '#fd8008',
							width: 1
						}
					}
				}
			]
		};
		
		
}
 


 