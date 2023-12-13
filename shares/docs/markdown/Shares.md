

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Shares]
- [Waiting to write...]

--------------------

### AddMyCode

#### 简要描述：

- [给自己添加一个监听]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.add_my_code

#### 请求方式：

- post

#### 请求参数:

- ` AddMyCodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`price` | 否|float64|当前价格   |
|`kdj` | 否|bool|日线金叉提醒   |
|`kdj20` | 否|bool|20日线提醒   |
|`surge` | 否|bool|盘中急涨提醒   |
|`slump` | 否|bool|盘中急跌提醒   |
|`ai` | 否|bool|AI智能提醒   |
|`public` | 否|bool|是否推荐给组织   |
|`up` | 否|float64|估价涨到   |
|`down` | 否|float64|估价跌到   |
|`upPercent` | 否|float64|涨幅超   |
|`downPercent` | 否|float64|跌幅超   |
|`vaild` | 否|bool|是否有效   |
|`simpleCode` | 否|string|股票代码简写// 返回值用   |
|`ext` | 否|string|后缀   |
|`name` | 否|string|股票名字   |
|`percent` | 否|float64|百分比   |
|`color` | 否|string|颜色   |


#### 请求示例:
```
{
     "ai": false,
     "code": "",
     "color": "",
     "down": 0,
     "downPercent": 0,
     "ext": "",
     "kdj": false,
     "kdj20": false,
     "name": "",
     "percent": 0,
     "price": 0,
     "public": false,
     "simpleCode": "",
     "slump": false,
     "surge": false,
     "up": 0,
     "upPercent": 0,
     "vaild": false
}
```

#### 返回参数说明:

- ` AddMyCodeResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`status` | 否|bool|状态是否成功   |
|`msg` | 否|string|消息   |


#### 返回示例:
	
```
{
     "msg": "",
     "status": false
}
```

#### 备注:

- 给自己添加一个监听

--------------------

### GetMsg

#### 简要描述：

- [获取消息]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.get_msg

#### 请求方式：

- post

#### 请求参数:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetMsgResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`shares.Msg`|消息列表   |


#### 返回示例:
	
```
{
     "list": [
          {
               "code": "",
               "color": "",
               "desc": "",
               "ext": "",
               "key": "",
               "name": "",
               "now": "",
               "percent": 0,
               "price": 0,
               "simpleCode": "",
               "tag": ""
          }
     ]
}
```

#### 备注:

- 获取消息

--------------------

### GetMyCode

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.get_my_code

#### 请求方式：

- post

#### 请求参数:

- ` GetMyCodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |


#### 请求示例:
```
{
     "code": ""
}
```

#### 返回参数说明:

- ` GetMyCodeResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`shares.AddMyCodeReq`|   |


#### 返回示例:
	
```
{
     "list": [
          {
               "ai": false,
               "code": "",
               "color": "",
               "down": 0,
               "downPercent": 0,
               "ext": "",
               "kdj": false,
               "kdj20": false,
               "name": "",
               "percent": 0,
               "price": 0,
               "public": false,
               "simpleCode": "",
               "slump": false,
               "surge": false,
               "up": 0,
               "upPercent": 0,
               "vaild": false
          }
     ]
}
```

#### 备注:

- 

--------------------

### Gets

#### 简要描述：

- [精确查找代码]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.gets

#### 请求方式：

- post

#### 请求参数:

- ` GetsReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`codes` | 否|[]string|股票代码   |


#### 请求示例:
```
{
     "codes": [
          ""
     ]
}
```

#### 返回参数说明:

- ` GetsResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`shares.SimpleSharesInfo`|列表(只返回价格相关的信息(code,price,percent,color))   |


#### 返回示例:
	
```
{
     "list": [
          {
               "code": "",
               "color": "",
               "percent": 0,
               "price": 0
          }
     ]
}
```

#### 备注:

- 精确查找代码

--------------------

### UpsetGroupCode

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.upset_group_code

#### 请求方式：

- post

#### 请求参数:

- ` UpsetGroupCodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`groupName` | 否|string|分组名   |
|`userName` | 否|string|推荐人   |
|`isAdd` | 否|bool|是否添加,false:删除   |


#### 请求示例:
```
{
     "code": "",
     "groupName": "",
     "isAdd": false,
     "userName": ""
}
```

#### 返回参数说明:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 

--------------------

### Dayliy

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.dayliy

#### 请求方式：

- post

#### 请求参数:

- ` CodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |


#### 请求示例:
```
{
     "code": ""
}
```

#### 返回参数说明:


#### 返回示例:
	
```
{}
```

#### 备注:

- 

--------------------

### GetMyGroup

#### 简要描述：

- [获取我的组织]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.get_my_group

#### 请求方式：

- post

#### 请求参数:

- ` CodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |


#### 请求示例:
```
{
     "code": ""
}
```

#### 返回参数说明:

- ` GetMyGroupResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`groupName` | 否|string|分组名   |
|`userName` | 否|string|推荐人   |
|`wi` | 否|int32|权重   |
|`group` | 否|[]string|我的股票列表   |


#### 返回示例:
	
```
{
     "code": "",
     "group": [
          ""
     ],
     "groupName": "",
     "userName": "",
     "wi": 0
}
```

#### 备注:

- 获取我的组织

--------------------

### HaveNewMsg

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.have_new_msg

#### 请求方式：

- post

#### 请求参数:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` HaveNewMsgResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`badge` | 否|bool|是否有新消息   |


#### 返回示例:
	
```
{
     "badge": false
}
```

#### 备注:

- 

--------------------

### Minute

#### 简要描述：

- [获取分时图]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.minute

#### 请求方式：

- post

#### 请求参数:

- ` CodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |


#### 请求示例:
```
{
     "code": ""
}
```

#### 返回参数说明:

- ` MinuteOut ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`yestclose` | 否|float64|   |
|`data` | 否|[]Array|   |
|`ref` | 否|bool|是否继续刷新   |


#### 返回示例:
	
```
{
     "data": "",
     "ref": false,
     "yestclose": 0
}
```

#### 备注:

- 获取分时图

--------------------

### Search

#### 简要描述：

- [搜索]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.search

#### 请求方式：

- post

#### 请求参数:

- ` SearchReq ` : 搜索请求

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`tag` | 否|string|标记(min,daily)   |


#### 请求示例:
```
{
     "code": "",
     "tag": ""
}
```

#### 返回参数说明:

- ` SearchResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`info` | 否|`shares.SharesInfo`|返回信息   |


#### 返回示例:
	
```
{
     "info": {
          "attach": "",
          "code": "",
          "color": "",
          "ext": "",
          "hy": "",
          "img": "",
          "name": "",
          "peg": "",
          "percent": 0,
          "price": 0,
          "simpleCode": ""
     }
}
```

#### 备注:

- 搜索

--------------------

### AddGroup

#### 简要描述：

- [添加一个组织]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.add_group

#### 请求方式：

- post

#### 请求参数:

- ` AddGroupReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`key` | 否|string|分组key   |


#### 请求示例:
```
{
     "key": ""
}
```

#### 返回参数说明:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 添加一个组织

--------------------

### DeleteMyCode

#### 简要描述：

- [删除一个监听]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.delete_my_code

#### 请求方式：

- post

#### 请求参数:

- ` DeleteMyCodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |


#### 请求示例:
```
{
     "code": ""
}
```

#### 返回参数说明:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 删除一个监听

--------------------

### GetAllCodeName

#### 简要描述：

- [获取所有中文]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.get_all_code_name

#### 请求方式：

- post

#### 请求参数:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetAllCodeNameResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`shares.CodeNameInfo`|代码列表   |


#### 返回示例:
	
```
{
     "list": [
          {
               "code": "",
               "name": "",
               "sName": ""
          }
     ]
}
```

#### 备注:

- 获取所有中文

--------------------

### GetGroup

#### 简要描述：

- [获取分组]

#### 请求URL:

- http://localhost:82/shares/api/v1/shares.get_group

#### 请求方式：

- post

#### 请求参数:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetGroupResp ` : 请求结构

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`shares.Group`|列表   |


#### 返回示例:
	
```
{
     "list": [
          {
               "list": [
                    {
                         "attach": "",
                         "code": "",
                         "color": "",
                         "ext": "",
                         "hy": "",
                         "img": "",
                         "name": "",
                         "peg": "",
                         "percent": 0,
                         "price": 0,
                         "simpleCode": ""
                    }
               ],
               "name": ""
          }
     ]
}
```

#### 备注:

- 获取分组
	

--------------------
--------------------

#### 自定义类型:

#### ` shares `


- ` Msg ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`simpleCode` | 否|string|股票代码简写   |
|`ext` | 否|string|后缀   |
|`name` | 否|string|股票名字   |
|`price` | 否|float64|当前价格   |
|`key` | 否|string|标签   |
|`desc` | 否|string|描述   |
|`percent` | 否|float64|百分比   |
|`color` | 否|string|颜色   |
|`now` | 否|string|当前时间   |
|`tag` | 否|string|标记(min,daily)   |



- ` AddMyCodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`price` | 否|float64|当前价格   |
|`kdj` | 否|bool|日线金叉提醒   |
|`kdj20` | 否|bool|20日线提醒   |
|`surge` | 否|bool|盘中急涨提醒   |
|`slump` | 否|bool|盘中急跌提醒   |
|`ai` | 否|bool|AI智能提醒   |
|`public` | 否|bool|是否推荐给组织   |
|`up` | 否|float64|估价涨到   |
|`down` | 否|float64|估价跌到   |
|`upPercent` | 否|float64|涨幅超   |
|`downPercent` | 否|float64|跌幅超   |
|`vaild` | 否|bool|是否有效   |
|`simpleCode` | 否|string|股票代码简写// 返回值用   |
|`ext` | 否|string|后缀   |
|`name` | 否|string|股票名字   |
|`percent` | 否|float64|百分比   |
|`color` | 否|string|颜色   |



- ` SimpleSharesInfo ` : 股票信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`price` | 否|float64|当前价格   |
|`percent` | 否|float64|百分比   |
|`color` | 否|string|颜色   |



- ` SharesInfo ` : 股票信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`simpleCode` | 否|string|股票代码简写   |
|`ext` | 否|string|后缀   |
|`name` | 否|string|股票名字   |
|`price` | 否|float64|当前价格   |
|`percent` | 否|float64|百分比   |
|`color` | 否|string|颜色   |
|`img` | 否|string|图片地址   |
|`hy` | 否|string|行业板块   |
|`attach` | 否|string|附加   |
|`peg` | 否|string|peg信息   |



- ` CodeNameInfo ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`code` | 否|string|股票代码   |
|`name` | 否|string|股票名字   |
|`sName` | 否|string|股票简写   |



- ` Group ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`name` | 否|string|分组名   |
|`list` | 否|[]`shares.SharesInfo`|列表   |




