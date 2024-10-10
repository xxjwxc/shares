#数字货币行情获取和指标计算演示
from  Ashare import *
from  MyTT import *
from MyTT_plus import *
import sys
from warnings import simplefilter
simplefilter(action='ignore', category=RuntimeWarning)

code = "sh000001"
############# 入口参数检测
if len(sys.argv) < 1 :
    sys.exit()

code = sys.argv[1]
#-----------------------end

#!/usr/bin/env python3
# -*- coding: utf-8 -*-

df=get_price(code,count=200,frequency='1d');      #日线数据获取  1d:1天  4h:4小时   60m: 60分钟    15m:15分钟
CLOSE=df.close.values;  OPEN=df.open.values;   HIGH=df.high.values;   LOW=df.low.values   #基础数据定义

# 支撑压力(boll 线)
UPPER = MA(CLOSE,20)+ 2*STD(CLOSE,20);
LOWER = MA(CLOSE,20)- 2*STD(CLOSE,20);
df['upper'] = UPPER
df['lower'] = LOWER

# 最后输出
# tmp = df[df['dwzq']>0]
df.index = df.index.strftime('%Y-%m-%d') # 日期转换

data = json.dumps(df.iloc[-92:].to_json(orient='index'))
print(json.loads(data))



