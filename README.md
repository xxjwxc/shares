# [shares](https://github.com/xxjwxc/shares)

### 功能

- A股量化交易系统
- 后台开发语言 Go  [gmsec](https://github.com/gmsec/gmsec)
- gormt 嵌入，自动数据库代码生成 [gorm 自动构建(gormt)](https://github.com/xxjwxc/gormt)
- 分时任务,每日监控,微信提醒功能，公众号
- uniapp 小程序端

### 欢迎扫码使用

![show](/image/0.png)

## 安装

- 进入到 server 目录
- 安装 cmake 工具
- 安装服务器
```
git clone git@github.com:xxjwxc/shares.git

cd shares
git submodule update --init --recursive

cd server
make run
```
- 客户端运行(hbuilder 直接导入 uniapp 即可)
  
## 部署运行

- 可直接运行程序
- 安装服务方式
```
sudo ./shares install
sudo ./shares start
```
or 
```
sudo ./shares stop
sudo ./shares run
```


## proto配置新加接口
- 修改目录`apidoc/proto/shares/`目录下相关proto文件
- 进入到`server`目录 使用`make gen`生成相关接口

## 配置说明
- 服务配置
```yaml
base:
    serial_number: "v1" # 版本号
    service_name: "shares" # 服务名
    service_displayname: "shares" # 服务显示名
    sercice_desc: "shares" # 描述
    is_dev: true
mysql_info:
    port : 3306 # 端口号
    username : root # 用户名
    host :  localhost # 地址
    password : 123456 # 密码
    # host : localhost
    # password : qwer
    database : caoguo_dev # 数据库名
kdniao: # 快递鸟配置
    business_id : 1317777
    app_key : 111111-2222-3333-4444-555555555
email: # 发邮件配置
    user: xie1xiao1jun@126.com
    password: pppppppppppppp
    host: smtp.126.com:25
wx_info: # 微信相关配置
    app_id : wxc111111111111
    app_secret : 111111111111111111111
    api_key : 1111111111111111111111111
    mch_id : 1111111111111111
    notify_url : http://www.xxjwxc.cn
    shear_url : ""
file_host: https://localhost/commcn/api/v1
oauth2_url: http://localhost/oauth2/api/v1
register_url: http://localhost/register/api/v1
token_type: nomal
app_id: wwwthings
app_secret: 4EE0A9A43B9B911C067BEE5CC50A9972
port : 8001
```
- uniapp 配置
 修改`shares\uniapp\commcn\utils\server\def.js` 中 `server.Host`进行服务器配置

 - 数据库说明
  详细请看`mysql`目录

## [传送门](https://github.com/xxjwxc/shares)

### 实际效果图

![show](/image/1.jpg)

![show](/image/2.jpg)

![show](/image/3.jpg)

![show](/image/4.jpg)

![show](/image/5.jpg)

![show](/image/6.jpg)

![show](/image/7.jpg)

![show](/image/8.jpg)



