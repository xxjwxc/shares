package config

import (
	"fmt"
)

// Config custom config struct
type Config struct {
	CfgBase     `yaml:"base"`
	MySQLInfo   MysqlDbInfo `yaml:"db_info"`
	EtcdInfo    EtcdInfo    `yaml:"etcd_info"`
	RedisDbInfo RedisDbInfo `yaml:"redis_info"`
	Oauth2Url   string      `yaml:"oauth2_url"`
	Port        string      `yaml:"port"` // 端口号
	WxInfo      WxInfo      `yaml:"wx_info"`
	FileHost    string      `yaml:"file_host"`
	MaxCapacity int         `yaml:"max_capacity"` // 最大容量
	DefGroup    string      `yaml:"def_group"`    // 默认分组
	Ext         []string    `yaml:"ext"`
	ToolsType   int         `yaml:"tools_type"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
	Database string `validate:"required"` // Database 数据库名
	Type     int    // 数据库类型: 0:mysql , 1:sqlite , 2:mssql
}

// RedisDbInfo redis database information. redis 数据库信息
type RedisDbInfo struct {
	Addrs     []string `yaml:"addrs"` // Host. 地址
	Password  string   // Password 密码
	GroupName string   `yaml:"group_name"` // 分组名字
	DB        int      `yaml:"db"`         // 数据库序号
}

// EtcdInfo etcd config info
type EtcdInfo struct {
	Addrs   []string `yaml:"addrs"`   // Host. 地址
	Timeout int      `yaml:"timeout"` // 超时时间(秒)
}

// WxInfo 微信相关配置
type WxInfo struct {
	AppID     string `yaml:"app_id"`     // 微信公众平台应用ID
	AppSecret string `yaml:"app_secret"` // 微信支付商户平台商户号
	APIKey    string `yaml:"api_key"`    // 微信支付商户平台API密钥
	MchID     string `yaml:"mch_id"`     // 商户号
	NotifyURL string `yaml:"notify_url"` // 通知地址
	ShearURL  string `yaml:"shear_url"`  // 微信分享url
}

// GetWxInfo 获取微信配置
func GetWxInfo() WxInfo {
	return _map.WxInfo
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

// GetMysqlDbInfo Get MySQL configuration information .获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}

// GetCheckTokenURL checktoken
func GetCheckTokenURL() string {
	return _map.Oauth2Url + "/check_token"
}

// GetLoginURL 登陆用的url
func GetLoginURL() string {
	return _map.Oauth2Url + "/authorize"
}

// GetLoginNoPwdURL token 授权模式登陆
func GetLoginNoPwdURL() string {
	return _map.Oauth2Url + "/authorize_nopwd"
}

// GetPort 获取端口号
func GetPort() string {
	return _map.Port
}

// GetRedisDbInfo Get redis configuration information .获取redis配置信息
func GetRedisDbInfo() RedisDbInfo {
	return _map.RedisDbInfo
}

// GetEtcdInfo get etcd configuration information. 获取etcd 配置信息
func GetEtcdInfo() EtcdInfo {
	return _map.EtcdInfo
}

// GetFileHost 获取文件host
func GetFileHost() string {
	return _map.FileHost
}

func GetMaxCapacity() int {
	return _map.MaxCapacity
}

func GetDefGroup() string {
	return _map.DefGroup
}

func GetExt() []string {
	return _map.Ext
}

func GetIsTools() int {
	return _map.ToolsType
}
