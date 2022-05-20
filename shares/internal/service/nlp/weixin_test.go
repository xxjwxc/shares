package nlp

import (
	"fmt"
	"testing"

	wx "github.com/xxjwxc/public/weixin"
)

func TestWXMain(t *testing.T) {
	SendCustomMsg(wx.CustomMsg{
		Touser:  "oxFCP6hcaZTReezFSs80ZY6qNAv8",
		Msgtype: "text",
		Text: &wx.CustomText{
			Content: "刘德华的歌WW",
		},
	})
	GetVoiceAsr("aaaa", "s4tLZ50qG05plLwL2kR0YXCDKgMLJuHfe003WumR-ySFdYz1v1J8FZiDhgi4AoAz")
	out, _ := GetAnswer("aaa", "刘德华的歌")
	if out.MsgType == "music" {
	}

	fmt.Println(out)
}
