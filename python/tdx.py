#数字货币行情获取和指标计算演示
from  Ashare import *
from  MyTT import *
from MyTT_plus import *
import sys
from warnings import simplefilter
simplefilter(action='ignore', category=RuntimeWarning)

all = False
code = "sh000001"
############# 入口参数检测
if len(sys.argv) < 2 :
    sys.exit()

code = sys.argv[1]
if len(sys.argv) > 2  :
    all = sys.argv[2] == "all"
#-----------------------end

#!/usr/bin/env python3
# -*- coding: utf-8 -*-

df=get_price(code,count=200,frequency='1d');      #日线数据获取  1d:1天  4h:4小时   60m: 60分钟    15m:15分钟
CLOSE=df.close.values;  OPEN=df.open.values;   HIGH=df.high.values;   LOW=df.low.values   #基础数据定义

# 飞鹰优选
VAR1 = CLOSE / MA(CLOSE, 40) * 100 < 78;
VAR2 = CLOSE / MA(CLOSE, 60) * 100 < 74;
VAR3 = HIGH > LOW * 1.051;
VAR4 = VAR3 & (COUNT(VAR3, 5) > 1);
TYP = (HIGH + LOW + CLOSE) / 3;
CCI = (TYP - MA(TYP, 14)) / (0.015 * AVEDEV(TYP, 14));
T1 = (MA(CLOSE, 27) > 1.169*CLOSE) & (MA(CLOSE, 17) > 1.158*CLOSE);
T2 = (CLOSE < MA(CLOSE, 120)) & (MA(CLOSE, 60) < MA(CLOSE, 120)) & (MA(CLOSE, 60) > MA(CLOSE, 30)) & (CCI > -210);
FYYH = VAR4 & (VAR1 | VAR2) & T1 & T2;
XG = BARSLASTCOUNT(FYYH) == 1;
df['fyyx'] = XG 
# tmp = df[df['fyyx']]
# data = json.dumps({'飞鹰优选':tmp.to_json(orient='index')})
# print(data)
# print(tmp)

# 主力真吸货
VAR1 = REF(LOW,1);
VAR2 = SMA(ABS(LOW-VAR1),3,1)/SMA(MAX(LOW-VAR1,0),3,1)*100;
VAR3 = EMA(IF(CLOSE*1.2,VAR2*10,VAR2/10),3);
VAR4 = LLV(LOW,38);
VAR5 = HHV(VAR3,38);
VAR6 = IF(LLV(LOW,90),1,0);
VAR7 = EMA(IF(LOW<=VAR4,(VAR3+VAR5*2)/2,0),3)/618*VAR6;
VAR8 = ((CLOSE-LLV(LOW,21))/(HHV(HIGH,21)-LLV(LOW,21)))*100;
VAR9 = SMA(VAR8,13,8);
ZLZXH =  VAR7 # 主力吸货
df['zlzxh'] = ZLZXH 
# tmp = df[df['zlzxh']>0]
# data = json.dumps({'主力吸货':tmp.to_json(orient='index')})
# print(data)
# print(tmp)

FX = SMA(VAR9,13,8) # 风险
df['fx'] = FX 
#tmp = df[df['fx']>0]
# data = json.dumps({'风险':tmp.to_json(orient='index')})
# print(data)
# print(tmp)

# 极低涨指标
VA1 = HHV(HIGH,9)-LLV(LOW,9);
VA2 = HHV(HIGH,9)-CLOSE;
VA3 = CLOSE-LLV(LOW,9);
VA4 = VA2/VA1*100-70;
VA5 = (CLOSE-LLV(LOW,60))/(HHV(HIGH,60)-LLV(LOW,60))*100;
VA6 = (2*CLOSE+HIGH+LOW)/4;
VA7 = SMA(VA3/VA1*100,3,1);
VA8 = LLV(LOW,34);
VA9 = SMA(VA7,3,1)-SMA(VA4,9,1);
VARA = IF(VA9>100,VA9-100,0);
VARB = HHV(HIGH,34);
VARC = EMA((VA6-VA8)/(VARB-VA8)*100,13);
VARD = EMA(0.667*REF(VARC,1)+0.333*VARC,2);
smx = EMA(VARD,1); # 生命线
df['smx'] = smx 

# VARE = SMA(MAX(CLOSE-REF(CLOSE,1),0)/CLOSE,8,1)/SMA(ABS(INDEXC-REF(INDEXC,1))/INDEXC,8,1)*100-25;
# VARF = MA(VARE,3);
# VAR1 = HHV(HIGH,9)-LLV(LOW,9);
# VAR2 = HHV(HIGH,9)-CLOSE;
# VAR3 = CLOSE-LLV(LOW,9);
# VAR4 = ((VAR2)/(VAR1))*(100)-70;
# VAR5 = ((CLOSE-LLV(LOW,60))/(HHV(HIGH,60)-LLV(LOW,60)))*(100);
# VAR6 = ((2)*(CLOSE)+HIGH+LOW)/(4);
# VAR7 = SMA(((VAR3)/(VAR1))*(100),3,1);
# VAR8 = LLV(LOW,34);
# VAR9 = SMA(VAR7,3,1)-SMA(VAR4,9,1);
# VAR10 = IF((VAR9>100),VAR9-100,0);
# VAR11 = HHV(HIGH,34);
# B1 = EMA(((VAR6-VAR8)/(VAR11-VAR8))*(100),8);
# B2 = EMA(B1,5),COLORFF7000;
# V1 = MA(CLOSE,5) = HHV(MA(CLOSE,5),20);
# V2:=MA(C,5)>MA(C,10);
# V3:=MA(V,5)>MA(V,40)*1.01;
# V4:=COUNT(MA(C,1)>REF(C,1),2)=2;
# VV:=V1 AND V2 AND V3 AND V4;
# STICKLINE(V1 AND V2 AND V3 AND V4,100,110,2,0),COLOR00FF00;
A1 = EMA(CLOSE,8);
A2 = EMA(A1,20);
A3 = CROSS(A1,A2);
A4 = A1<EMA(CLOSE,120);
# STICKLINE(A3 AND A4,100,110,2,0),COLOR00FF00;
A5 = 3*SMA((CLOSE-LLV(LOW,18))/(HHV(HIGH,18)-LLV(LOW,18))*100,21,1)-2*SMA(SMA((CLOSE-LLV(LOW,18))/(HHV(HIGH,18)-LLV(LOW,18))*100,20,1),8,1);
dwzq = IF(CROSS(A5,20),50,0);# 低位转强
df['dwzq'] = dwzq 

ksls = IF(CROSS(VARC,smx) & (VARC>VARD) & (VARC<55.0) & (CLOSE>OPEN),50,0);# 开始拉升
df['ksls'] = ksls 
A6 = CLOSE/REF(CLOSE,1)>(1+9.8/100);
jdz = IF(A6,50,0);# 极低涨
df['jdz'] = jdz 

# 唐奇安指标
M1 = MA(CLOSE,20);
HPB = REF(HHV(HIGH,20),1);
LPB = REF(LLV(LOW,10),1);
BUYIN = (HIGH>HPB) & (CLOSE>M1);
BUYEXIT = LOW<LPB;
SINGAL = IF(BUYIN,1,IF(BUYEXIT,-1,0));
tqa_tag = 0
for index in range(len(SINGAL)):
    if tqa_tag == SINGAL[index] : SINGAL[index] = 0
    elif SINGAL[index] != 0:tqa_tag = SINGAL[index] 
df['tqa'] = SINGAL
# print(CURRSIGN)

# 主力短线启动
zldxqd = ((CLOSE>MA(CLOSE,13)) & (MA(CLOSE,55)>REF(MA(CLOSE,55),1)) & (CLOSE>OPEN) & (CLOSE>REF(CLOSE,1)) & (MA(CLOSE,13)>REF(MA(CLOSE,13),1)) & (MA(CLOSE,13)>MA(CLOSE,55)));
df['zldq'] = zldxqd

# 主力清洗
VAR1 =REF((LOW+OPEN+CLOSE+HIGH)/4,1);
VAR2 =SMA(ABS(LOW-VAR1),13,1)/SMA(MAX(LOW-VAR1,0),10,1);
VAR3 =EMA(VAR2,10);
VAR4 =LLV(LOW,33);
VAR5 =EMA(IF(LOW<=VAR4,VAR3,0),3);
zlqx_jc = IF(VAR5>REF(VAR5,1),VAR5,0);# 主力进场
zlqx_xp = IF(VAR5<REF(VAR5,1),VAR5,0);# 洗盘
VAR21 =SMA(ABS(HIGH-VAR1),13,1)/SMA(MIN(HIGH-VAR1,0),10,1);
VAR31 =EMA(VAR21,10);
VAR41 =HHV(HIGH,33);
VAR51 =EMA(IF(HIGH>=VAR41,VAR31,0),3);
zlqx_lg = IF(VAR51<REF(VAR51,1),VAR51,0);# 主力拉高
# STICKLINE(VAR51<REF(VAR51,1),0,VAR51,3,0),coloryellow;
zlqx_ch = IF(VAR51>REF(VAR51,1),VAR51,0);# 出货
# STICKLINE(VAR51>REF(VAR51,1),0,VAR51,3,0 ),colorcyan;
df['zlqx_jc'] = zlqx_jc# 主力进场
df['zlqx_xp'] = zlqx_xp# 洗盘
df['zlqx_lg'] = zlqx_lg# 主力拉高
df['zlqx_ch'] = zlqx_ch# 出货

# RSI
LC = REF(CLOSE,1)
RSI6 = SMA(MAX(CLOSE-LC,0),6,1)/SMA(ABS(CLOSE-LC),6,1)*100
df['rsi6'] = RSI6# 出货

# 支撑压力(boll 线)
# UPPER = MA(CLOSE,20)+ 2*STD(CLOSE,20);
# LOWER = MA(CLOSE,20)- 2*STD(CLOSE,20);
# df['upper'] = UPPER
# df['lower'] = LOWER

# 筹码集中度:(COST(90+(100-90)/2)-COST((100-90)/2))/(COST(90+(100-90)/2)+COST((100-90)/2))*100/2;

# 最后输出
# tmp = df[df['dwzq']>0]
df.index = df.index.strftime('%Y-%m-%d') # 日期转换

if all :
    data = json.dumps(df.iloc[-30:].to_json(orient='index'))
    print(json.loads(data))
else :
    data = json.dumps(df.iloc[-1:].to_json(orient='index'))
    print(json.loads(data))
