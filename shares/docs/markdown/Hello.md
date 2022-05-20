

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Hello]
- [Waiting to write...]

--------------------

#### 简要描述：

- []

#### 请求URL:

- http://localhost:82/shares/api/v1/hello.say_hello

#### 请求方式：

- post

#### 请求参数:

- ` HelloRequest ` : 请求结构

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`name` | 否|string|名字   |


#### 请求示例:
```
{
     "name": ""
}
```

#### 返回参数说明:

- ` HelloReply ` : 响应结构

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`message` | 否|string|消息   |
|`result` | 否|`common.ResultResp`|多proto 文件样例   |


#### 返回示例:
	
```
{
     "message": "",
     "result": {
          "msg": "",
          "result": false
     }
}
```

#### 备注:

- 
	

--------------------
--------------------

#### 自定义类型:

#### ` common `


- ` ResultResp ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`result` | 否|bool|   |
|`msg` | 否|string|   |




