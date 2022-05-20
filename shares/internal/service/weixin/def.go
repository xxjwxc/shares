package weixin

const (
	_cacheSessionkey = "session_key"
)

// WxSessionKey ...
type WxSessionKey struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}
