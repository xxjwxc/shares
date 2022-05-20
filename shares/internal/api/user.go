package api

import (
	"shares/internal/core"
	"shares/internal/model"
	"strconv"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
)

// UserInfo 用户信息
type UserInfo struct {
	AccessToken string // token
	UserName    string // 用户名
	ExpireTime  int    // 超时时间
	Info        model.WxUserinfo
}

// MapMessageBody req message body
type MapMessageBody struct {
	message.MessageBody
	Data map[string]string `json:"data,omitempty"`
}

// GetUserFromOpenID 通过openid获取用户信息
func GetUserFromOpenID(openid string) (userInfo *UserInfo, err error) {
	// 先从缓存中获取
	if tmp, ok := GetCacheBody(openid); ok {
		return &UserInfo{
			AccessToken: openid,
			UserName:    tmp.UserName,
			ExpireTime:  tmp.ExpireTime,
			Info:        tmp.Info,
		}, nil
	}

	// 从数据库里面拿
	orm := core.Dao.GetDBr()
	info, err := model.WxUserinfoMgr(orm.DB).GetFromOpenid(openid)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}

	userInfo = &UserInfo{
		AccessToken: openid,
		UserName:    info.NickName,
		ExpireTime:  (2 * 60 * 60),
		Info:        info}

	SaveCacheBody(userInfo, openid, (2 * 60 * 60)) // 保存缓存
	return userInfo, nil                           // 返回结果
}

// SaveCacheBody 保存缓存
func SaveCacheBody(userInfo *UserInfo, openid string, expireTime int) {
	// 保存缓存
	cache := mycache.NewCache("oauth2")
	cache.Add(openid, userInfo, time.Duration(userInfo.ExpireTime)*time.Second)
	//------------------end
}

func DeleteCacheBody(openid string) { // 删除缓存
	cache := mycache.NewCache("oauth2")
	cache.Delete(openid)
}

// GetCacheBody 获取缓存
func GetCacheBody(token string) (value *UserInfo, b bool) {
	value = new(UserInfo)
	cache := mycache.NewCache("oauth2")
	err := cache.Value(token, &value)
	if err != nil {
		return nil, false
	}
	return value, true
}

func strToInt(src string) int {
	i, _ := strconv.Atoi(src)
	return i
}
