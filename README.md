# [shares](https://github.com/xxjwxc/shares)

### 功能

- A股量化交易系统
- 后台开发语言 Go/Python  [gmsec](https://github.com/gmsec/gmsec)
- gormt 嵌入，自动数据库代码生成 [gorm 自动构建(gormt)](https://github.com/xxjwxc/gormt)
- 分时任务,盯盘助手,研报股评,每日监控,微信提醒，玩转组织，AI智能
- uniapp 小程序端
<!-- ### 如果喜欢的朋友可以加作者微信(jnnpher)，作者建立了一个群，答疑解惑，方便交流 -->

### 欢迎微信扫码体验

![show](/image/0.png) 

### Vscode插件地址 [复利备忘录](https://marketplace.visualstudio.com/items?itemName=a-shares.a-shares)
(插件市场搜索:复利备忘录)

## 安装

- 进入到 shares 目录
- 安装 cmake 工具
- 安装服务器
```
git clone git@github.com:xxjwxc/shares.git

cd shares
git submodule update --init --recursive

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
- mysql
### 目录位置 [mysql](mysql)

[shares_tmp_db.sql（表结构&&数据）](mysql%2Fshares_tmp_db.sql)<br>
[shares_tmp_db_views.sql（views视图）](mysql%2Fshares_tmp_db_views.sql)<br>
[shares_tmp_db.sql.zip（两个sql的综合压缩文件）](mysql%2Fshares_tmp_db.sql.zip)<br>

- 服务配置
```text
    # 修改配置文件
    vim shares/conf/config.yml
```
```yaml
base:
    is_dev : true
    serial_number : 1.0.0
    service_name : shares
    service_displayname : sharesserver
    sercice_desc : shares微服务
tools_type: 4 # 0:正式环境,1:日分析,2:抓取消息,3:放量,4:放量监听
db_info:
    port : 3306 # 端口号
    username : root # 用户名
    host :  localhost # 地址
    password : 123456 # 密码
    # host : localhost
    # password : qwer
    database : caoguo_dev # 数据库名
redis_info:
    addrs: ["127.0.0.1:6379"]
    password: 123456
    group_name: oauth2
    db: 0
etcd_info:
    addrs: ["127.0.0.1:2379"]
    timeout: 3
wx_info:
    app_id : wxxxxxxxxxxxx31a
    app_secret : xxxxxxxxxxxxxxxxxxxxxxxx
    api_key : xxxxxxxxxxxxxxxxxxx
    mch_id : xxxxxxxxx
    notify_url : http://www.localhost.com
    shear_url : 
port: 82
file_host: https://www.localhost.com/shares/api/v1
max_capacity : 5
def_group: 默认指标
ext : [sh,sz,hk]
```
- uniapp 配置
 修改`shares\uniapp\commcn\utils\server\def.js` 中 `server.Host`进行服务器配置

 - 数据库说明
  详细请看`mysql`目录

## [传送门](https://github.com/xxjwxc/shares)

### 实际效果图

<p align="center" margin: 0 auto;>
<img src="/image/v_1.jpg" width=90%>
</p>
<p align="center" margin: 0 auto;>
<img src="/image/v_2.jpg" width=90%>
</p>


<p align="center" margin: 0 auto;>
<img src="/image/1.jpg" width=45%>
<img src="/image/2.jpg" width=45%>
</p>
<p align="center" margin: 0 auto;>
<img src="/image/3.jpg" width=45%>
<img src="/image/4.jpg" width=45%>
</p>

<p align="center" margin: 0 auto;>
<img src="/image/5.jpg" width=45%>
<img src="/image/6.jpg" width=45%>
</p>

<p align="center" margin: 0 auto;>
<img src="/image/7.jpg" width=45%>
<img src="/image/8.jpg" width=45%>
</p>




