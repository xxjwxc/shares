package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"

	"github.com/mozillazg/go-pinyin"
)

var _pinyin = pinyin.NewArgs()

func initPinyin() error {
	orm := core.Dao.GetDBw()
	list, err := model.SharesInfoTblMgr(orm.DB).Gets()
	if err != nil {
		return err
	}
	for _, v := range list {
		sampleName := FirstLetterOfPinYin(v.Name)
		model.SharesInfoTblMgr(orm.Where("id = ?", v.ID)).Update(model.SharesInfoTblColumns.SimpleName, sampleName)
	}

	return nil
}

func FirstLetterOfPinYin(name string) string {
	result := pinyin.Pinyin(name, _pinyin)
	if len(result) == 0 {
		return ""
	}
	var out string
	for _, v := range result {
		out += fmt.Sprintf("%s", string(v[0][0]))
	}
	return out
}
