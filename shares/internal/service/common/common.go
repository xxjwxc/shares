package common

import (
	"fmt"
	"shares/internal/api"
	"shares/internal/core"
	"shares/internal/model"
	"strings"

	"github.com/xxjwxc/public/message"
)

// AddGroup 添加一个组织()
func AddGroup(openID, key string) (isAdd bool, _err error) {
	user, err := api.GetUserFromOpenID(openID)
	if err != nil {
		return false, err
	}

	orm := core.Dao.GetDBw()
	out, err := model.GroupListTblMgr(orm.DB).GetFromKey(key)
	if err != nil {
		return false, err
	}
	if out.ID == 0 {
		return false, message.GetError(message.NotFindError)
	}
	list := strings.Split(user.Info.Group, ",")
	for index, v := range list { // 操作是否删除
		if strings.EqualFold(v, out.GroupName) { // 已经存在
			list = append(list[:index], list[index+1:]...)
			model.WxUserinfoMgr(orm.Where("openid=?", user.Info.Openid)).Update("group", strings.Join(list, ","))
			api.DeleteCacheBody(user.Info.Openid) // 清除缓存
			return false, nil
		}
	}

	list = append(list, out.GroupName)
	tmp := list[:0]
	for _, v := range list {
		if len(v) > 0 {
			tmp = append(tmp, v)
		}
	}
	info, _ := model.WxUserinfoMgr(orm.DB).GetFromOpenid(user.Info.Openid)
	if info.ID == 0 {
		return false, fmt.Errorf("未找到用户信息，请先点击'首页'并且同意获取用户信息之后再重试。")
	}

	model.WxUserinfoMgr(orm.Where("openid=?", user.Info.Openid)).Update("group", strings.Join(tmp, ","))
	api.DeleteCacheBody(user.Info.Openid) // 清除缓存
	return true, nil
}
