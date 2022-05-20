package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

type pegResp struct {
	Status  int    `json"status"`
	Message string `json:"message"`
	Gzbj    Gzbj   `json:"gzbj"` // 个股比较
}

type Gzbj struct {
	BaseYear int       `json:"baseYear"`
	Data     []pegData `json:"data"`
}

type pegData struct {
	Dm  string `json:"dm"`
	Jc  string `json:"jc"`
	Peg string `json:"peg"`
}

// updatePeg 更新peg信息
func updatePeg() {
	// https://emweb.securities.eastmoney.com/PC_HSF10/IndustryAnalysis/Index?type=web&code=sz300750
	day0 := tools.GetUtcDay0(time.Now())
	var codes []string
	orm := core.Dao.GetDBw()
	model.SharesInfoTblMgr(orm.DB).Select("shares_info_tbl.code as code").Joins("left join no_peg_tbl on (no_peg_tbl.code = shares_info_tbl.code) where no_peg_tbl.code IS NULL").Scan(&codes)
	for _, code := range codes {
		info, err := getPeg(code)
		if info != nil && len(info.Gzbj.Data) >= 3 {
			var srcs []string
			for i := 3; i < len(info.Gzbj.Data); i++ {
				srcs = append(srcs, fmt.Sprintf("%v:%v", info.Gzbj.Data[i].Jc, info.Gzbj.Data[i].Peg))
			}
			model.SharesInfoTblMgr(orm.Where("code = ?", code)).Updates(
				map[string]interface{}{
					model.SharesInfoTblColumns.Peg:       stringToFloat(info.Gzbj.Data[0].Peg),
					model.SharesInfoTblColumns.PegAvg:    stringToFloat(info.Gzbj.Data[1].Peg),
					model.SharesInfoTblColumns.PegAve:    stringToFloat(info.Gzbj.Data[2].Peg),
					model.SharesInfoTblColumns.PegShares: strings.Join(srcs, ";"),
				})
			model.SharesDailyTblMgr(orm.Where("code = ? and day0 = ?", code, day0)).Update(model.SharesDailyTblColumns.Peg, stringToFloat(info.Gzbj.Data[0].Peg))
		} else {
			mylog.Error(err)
		}
	}
}

// getPeg 获取peg信息
func getPeg(code string) (*pegResp, error) {
	url := fmt.Sprintf("https://emweb.securities.eastmoney.com/PC_HSF10/IndustryAnalysis/IndustryAnalysisAjax?code=%v&icode=1035", code)
	out := myhttp.OnGetJSON(url, "") // 拉取分类
	var resp pegResp
	tools.JSONEncode(out, &resp)
	if resp.Status != 0 {
		if resp.Status == -1 { // 注入
			model.NoPegTblMgr(core.Dao.GetDBw().DB).Save(&model.NoPegTbl{
				Code:      code,
				CreatedAt: time.Now(),
			})
		} else {
			mylog.Error(out)
		}

		return nil, fmt.Errorf(resp.Message)
	}
	return &resp, nil
}
