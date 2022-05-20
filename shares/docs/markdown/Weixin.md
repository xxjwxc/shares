

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Weixin]
- [Waiting to write...]

--------------------

### ReLogin

#### 简要描述：

- [是否要重新登录]

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.re_login

#### 请求方式：

- post

#### 请求参数:


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` ReLoginResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`relogin` | 否|bool|是否重新登录:true 重新登录   |
|`sessionId` | 否|string|用户sessionid 用于当前交互使用   |
|`openId` | 否|string|openid   |
|`overdueTime` | 否|int64|逾期时间点(时间戳)   |


#### 返回示例:
	
```
{
     "openId": "",
     "overdueTime": 0,
     "relogin": false,
     "sessionId": ""
}
```

#### 备注:

- 是否要重新登录

--------------------

### UpdateUserInfo

#### 简要描述：

- [更新用户信息]

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.update_user_info

#### 请求方式：

- post

#### 请求参数:

- ` WxUserinfo ` : 用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`sessionID` | 否|string|用户sessionid 用于当前交互使用   |
|`nickName` | 否|string|用户昵称   |
|`avatarURL` | 否|string|用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。   |
|`gender` | 否|string|用户的性别，值为1时是男性，值为2时是女性，值为0时是未知   |
|`city` | 否|string|用户所在城市   |
|`province` | 否|string|用户所在省份   |
|`country` | 否|string|用户所在国家   |
|`language` | 否|string|用户的语言，简体中文为zh_CN   |


#### 请求示例:
```
{
     "avatarURL": "",
     "city": "",
     "country": "",
     "gender": "",
     "language": "",
     "nickName": "",
     "province": "",
     "sessionID": ""
}
```

#### 返回参数说明:


#### 返回示例:
	
```
{}
```

#### 备注:

- 更新用户信息

--------------------

### UpsetUserInfo

#### 简要描述：

- [更新用户信息]

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.upset_user_info

#### 请求方式：

- post

#### 请求参数:

- ` UpsetUserInfoReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`key` | 否|string|   |
|`value` | 否|string|   |


#### 请求示例:
```
{
     "key": "",
     "value": ""
}
```

#### 返回参数说明:


#### 返回示例:
	
```
{}
```

#### 备注:

- 更新用户信息

--------------------

### GetQrcode

#### 简要描述：

- [获取微信二维码]

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.get_qrcode

#### 请求方式：

- post

#### 请求参数:

- ` GetQrcodeReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`page` | 否|string|小程序页面头部   |
|`scene` | 否|string|附带参数   |


#### 请求示例:
```
{
     "page": "",
     "scene": ""
}
```

#### 返回参数说明:

- ` GetQrcodeResp ` : 获取二维码

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`url` | 否|string|二维码图片地址   |


#### 返回示例:
	
```
{
     "url": ""
}
```

#### 备注:

- 获取微信二维码

--------------------

### GetUserInfo

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.get_user_info

#### 请求方式：

- post

#### 请求参数:


#### 请求示例:
```
{}
```

#### 返回参数说明:

- ` GetUserInfoResp ` : 用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`openid` | 否|string|用户openid   |
|`nickName` | 否|string|用户昵称   |
|`avatarURL` | 否|string|用户头像地址   |
|`gender` | 否|string|用户的性别   |
|`city` | 否|string|用户所在城市   |
|`province` | 否|string|用户所在省份   |
|`country` | 否|string|用户所在国家   |
|`phone` | 否|string|用户手机号   |
|`group` | 否|string|用户所在分组列表   |
|`rg` | 否|bool|是否涨绿跌红   |
|`only20` | 否|bool|是否20日线   |
|`capacity` | 否|int32|可用容量   |


#### 返回示例:
	
```
{
     "avatarURL": "",
     "capacity": 0,
     "city": "",
     "country": "",
     "gender": "",
     "group": "",
     "nickName": "",
     "only20": false,
     "openid": "",
     "phone": "",
     "province": "",
     "rg": false
}
```

#### 备注:

- 

--------------------

### Oauth

#### 简要描述：

- [微信授权获取登录信息]

#### 请求URL:

- http://localhost:82/shares/api/v1/weixin.oauth

#### 请求方式：

- post

#### 请求参数:

- ` OauthReq ` : 请求结构

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`jscode` | 否|string|微信端jscode   |
|`isUpdateUser` | 否|bool|是否更新用户   |


#### 请求示例:
```
{
     "isUpdateUser": false,
     "jscode": ""
}
```

#### 返回参数说明:

- ` OauthResp ` : 微信授权返回信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`sessionId` | 否|string|用户sessionid 用于当前交互使用   |
|`openId` | 否|string|openid   |
|`overdueTime` | 否|int64|逾期时间点(时间戳)   |


#### 返回示例:
	
```
{
     "openId": "",
     "overdueTime": 0,
     "sessionId": ""
}
```

#### 备注:

- 微信授权获取登录信息
	

--------------------
--------------------

#### 自定义类型:


