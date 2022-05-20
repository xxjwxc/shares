package shares

const (
	_ChineseCode = "ChineseCode_cache"
)

type MinuteOut struct {
	Yestclose float64         `json:"yestclose"`
	Data      [][]interface{} `json:"data"`
	Ref       bool            `json:"ref"` // 是否继续刷新
}
