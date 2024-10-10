import numpy as np
import pandas as pd

# 生成模拟数据
data = pd.DataFrame({'close': np.random.rand(100)})

# 计算移动平均线
data['ma'] = data['close'].rolling(window=10).mean()

# 计算COST指标
data['cost'] = (data['close'] - data['ma']).abs() / data['ma'].abs() * 100

# 计算百分比差
data['diff'] = data['close'] - data['ma']
data['percent_diff'] = (data['diff'] / data['ma']).abs * 100

# 计算筹码分布
data['cost_ratio'] = data['cost'] / data['percent_diff'] * 100

# 绘制筹码分布图
data.plot(kind='bar', color='blue', alpha=0.5, title='COST Chart')