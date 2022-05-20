package analy

import (
	"fmt"
	"time"

	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
)

const (
	_wxMsgCode = "wxOpenapiSign:"
)

// var _once sync.Once
var _cache mycache.CacheIFS

func init() {
	_cache = mycache.NewCache("wx_msg")
}

// GetCacheWxMsgCode 获取
func GetCacheWxMsgCode(openID string) string {
	key := fmt.Sprintf("%v%v", _wxMsgCode, openID)
	if _cache.IsExist(key) {
		var tmp string
		err := _cache.Value(key, &tmp)
		if err == nil { // 没有错误
			return tmp
		}
		mylog.Error(err)
	}

	return ""
}

// SetCacheWxMsgCode 设置
func SetCacheWxMsgCode(openID, value string, tm time.Duration) error {
	key := fmt.Sprintf("%v%v", _wxMsgCode, openID)
	return _cache.Add(key, value, tm)
}

// DeleteCacheWxMsgCode 删除
func DeleteCacheWxMsgCode(openID string) error {
	key := fmt.Sprintf("%v%v", _wxMsgCode, openID)
	return _cache.Delete(key)
}
