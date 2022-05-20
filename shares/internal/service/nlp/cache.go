package nlp

/**
缓存：支持redis，内存缓存
*/
import (
	"fmt"
	"time"

	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
)

const (
	_wxOpenapiSign = "wxOpenapiSign:"
	_motionData    = "motionData:"
)

// var _once sync.Once
var _cache mycache.CacheIFS

func init() {
	_cache = mycache.NewCache("wx_nlp")
}

// GetCacheWxOpenapiSign 获取微信sing信息
func GetCacheWxOpenapiSign(avatarID string) string {
	key := fmt.Sprintf("%v%v", _wxOpenapiSign, avatarID)
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

// SetCacheWxOpenapiSign 设置微信sign信息
func SetCacheWxOpenapiSign(avatarID, value string, tm time.Duration) error {
	key := fmt.Sprintf("%v%v", _wxOpenapiSign, avatarID)
	return _cache.Add(key, value, tm)
}

// DeleteCacheWxOpenapiSign 删除微信sign信息
func DeleteCacheWxOpenapiSign(username, userid string) error {
	key := fmt.Sprintf("%v%v-%v", _wxOpenapiSign, username, userid)
	return _cache.Delete(key)
}
