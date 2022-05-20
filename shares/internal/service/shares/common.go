package shares

import (
	"shares/internal/core"
	"shares/internal/model"
	"time"
)

func UpDataGroup(groups []string, code, userName string) {
	orm := core.Dao.GetDBw()
	for _, v := range groups {
		if len(v) > 0 {
			out, _ := model.GroupTblMgr(orm.DB).FetchUniqueIndexByGroupName(v, code)
			out.Code = code
			out.GroupName = v
			out.UserName = userName
			if out.ID == 0 {
				out.CreatedAt = time.Now()
				out.UpdateAt = time.Now()
			} else {
				out.UpdateAt = time.Now()
			}
			out.Wi += 1
			model.GroupTblMgr(orm.DB).Save(&out)
		}
	}
}

// func GetShare(code string, rg bool) (*proto.SharesInfo, error) {
// 	tmp := &proto.SharesInfo{}
// 	err := mycache.NewCache(event.CacheCode).Value(code, tmp)
// 	if err == nil {
// 		return tmp, nil
// 	}

// 	// 开盘期间实时获取
// 	day0 := tools.GetUtcDay0(time.Now())
// 	offset := time.Now().Unix() - day0
// 	week := tools.GetTimeWeek(day0)
// 	if (week > 0 && week < 6) && (offset >= (9*60*60+15*60) && offset <= (15*60*60+1*60)) { // 实时获取
// 		outs := event.Searchs([]string{code})
// 		if len(outs) > 0 {
// 			v := outs[0]
// 			return &proto.SharesInfo{
// 				Code:       v.Code,
// 				SimpleCode: v.SimpleCode,
// 				Ext:        v.Ext,
// 				Name:       v.Name,
// 				Price:      v.Price,
// 				Percent:    v.Percent,
// 				Color:      GetColor(rg, v.Percent),
// 			}, nil
// 		}
// 	}
// 	// ----------------end

// 	orm := core.Dao.GetDBr()
// 	v, _ := model.SharesInfoTblMgr(orm.DB).GetFromCode(code)
// 	return &proto.SharesInfo{
// 		Code:       v.Code,
// 		SimpleCode: v.SimpleCode,
// 		Ext:        v.Ext,
// 		Name:       v.Name,
// 		Price:      v.Price,
// 		Percent:    v.Percent,
// 		Color:      GetColor(rg, v.Percent),
// 	}, nil
// }

// func GetShares(codes []string, rg bool) (out []*proto.SharesInfo, err error) {
// 	var remain []string
// 	for _, code := range codes { // 先走缓存拿
// 		tmp := &proto.SharesInfo{}
// 		err := mycache.NewCache(event.CacheCode).Value(code, tmp)
// 		if err != nil {
// 			remain = append(remain, code)
// 		} else {
// 			tmp.Color = GetColor(rg, tmp.Percent)
// 			out = append(out, tmp)
// 		}
// 	}

// 	// 开盘期间实时获取
// 	day0 := tools.GetUtcDay0(time.Now())
// 	offset := time.Now().Unix() - day0
// 	week := tools.GetTimeWeek(day0)
// 	if (week > 0 && week < 6) && (offset >= (9*60*60+15*60) && offset <= (15*60*60+1*60)) { // 实时获取
// 		outs := event.Searchs(remain)
// 		for _, v := range outs {
// 			out = append(out, &proto.SharesInfo{
// 				Code:       v.Code,
// 				SimpleCode: v.SimpleCode,
// 				Ext:        v.Ext,
// 				Name:       v.Name,
// 				Price:      v.Price,
// 				Percent:    v.Percent,
// 				Color:      GetColor(rg, v.Percent),
// 			})
// 		}

// 		return out, nil
// 	}
// 	// ----------------end

// 	orm := core.Dao.GetDBr()
// 	list, _ := model.SharesInfoTblMgr(orm.DB).GetBatchFromCode(remain)
// 	for _, v := range list {
// 		out = append(out, &proto.SharesInfo{
// 			Code:       v.Code,
// 			SimpleCode: v.SimpleCode,
// 			Ext:        v.Ext,
// 			Name:       v.Name,
// 			Price:      v.Price,
// 			Percent:    v.Percent,
// 			Color:      GetColor(rg, v.Percent),
// 		})
// 	}

// 	return out, nil
// }

// func GetColor(rg bool, percent float64) string {
// 	if !rg && percent >= 0 { // 默认红色
// 		return "#ff0000"
// 	}

// 	if rg && percent <= 0 { // 默认红色
// 		return "#ff0000"
// 	}

// 	return "#1AAD19" // 绿色
// }
