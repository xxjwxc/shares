/* 用户信息相关 */
// import { data } from 'uview-ui/libs/mixin/mixin';
import Server from './def'

const shares = {}

shares.GetGroup = async function() {
	let resp = await Server.Post("/shares.get_group", {});
	return Server.GetReturn(resp);
};

shares.GetTop = async function() {
	let resp = await Server.Post("/shares.get_top", {});
	return Server.GetReturn(resp);
};

shares.GetMyGroup = async function(code) {
	let resp = await Server.Post("/shares.get_my_group", {
		code: code
	}); 
	return Server.GetReturn(resp);
};

shares.AddGroup = async function(key) {
	let resp = await Server.Post("/shares.add_group", {
		key: key
	});
	return Server.GetReturn(resp);
};

shares.UpsetGroupCode = async function(code, groupName, userName, isAdd) {
	let resp = await Server.Post("/shares.upset_group_code ", {
		code: code,
		groupName: groupName,
		userName: userName,
		isAdd: isAdd
	});
	return Server.GetReturn(resp);
};

shares.Gets = async function(codes) {
	let resp = await Server.Post("/shares.gets", {
		codes: codes
	});
	return Server.GetReturn(resp);
};

shares.GetAllCodeName = async function(){
	let resp = await Server.Post("/shares.get_all_code_name", {});
	return Server.GetReturn(resp);
};

shares.Search = async function(code, tag) {
	let resp = await Server.Post("/shares.search", {
		code: code,
		tag: tag
	});
	return Server.GetReturn(resp);
};

shares.AddMyCode = async function(req) {
	let resp = await Server.Post("/shares.add_my_code", req);
	return Server.GetReturn(resp);
};

shares.GetMyCode = async function(code,yjzd) {
	let resp = await Server.Post("/shares.get_my_code", {
		code: code,
		yjzd:yjzd,
	});
	return Server.GetReturn(resp);
};

shares.GetMyMinute = async function(codes) {
	let resp = await Server.Post("/shares.get_my_minute", {
		codes: codes
	});
	return Server.GetReturn(resp);
};


shares.GetMsg = async function() {
	let resp = await Server.Post("/shares.get_msg", {});
	return Server.GetReturn(resp);
};

shares.HaveNewMsg = async function() {
	let resp = await Server.Post("/shares.have_new_msg", {});
	return Server.GetReturn(resp);
};

shares.DeleteMyCode = async function(code) {
	let resp = await Server.Post("/shares.delete_my_code", {
		code: code
	});
	return Server.GetReturn(resp);
};



shares.GetVip = async function(days){
	let resp = await Server.Post("/shares.get_vip", {days:days});
	return Server.GetReturn(resp);
};
shares.GetFlkx = async function(days){
	let resp = await Server.Post("/shares.get_flkx", {days:days});
	return Server.GetReturn(resp);
};

shares.GetZdsx = async function(days){
	let resp = await Server.Post("/shares.get_zdsx", {days:days});
	return Server.GetReturn(resp);
};

shares.GetVips = async function(days){
	let resp = await Server.Post("/shares.get_vip_super", {days:days});
	return Server.GetReturn(resp);
};

shares.GetCd = async function(days){
	let resp = await Server.Post("/shares.get_cd", {days:days});
	return Server.GetReturn(resp);
};
shares.GetFl = async function(days){
	let resp = await Server.Post("/shares.get_fl", {days:days});
	return Server.GetReturn(resp);
};

shares.GetLhb = async function(days){
	let resp = await Server.Post("/shares.get_lhb", {days:days});
	return Server.GetReturn(resp);
};

shares.GetUp = async function(days){
	let resp = await Server.Post("/shares.get_up", {days:days});
	return Server.GetReturn(resp);
};

shares.GetHyDjCodeName =  async function(hyCode) {
	let resp = await Server.Post("/shares.get_hy_dj_code_name", {hyCode:hyCode});
	return Server.GetReturn(resp);
};

shares.GetSp = async function(days){
	let resp = await Server.Post("/shares.get_sp", {days:days});
	return Server.GetReturn(resp);
};
shares.GetSpEx = async function(days) {
	let resp = await Server.Post("/shares.get_sp_ex", {
		days:days, 
	});
	return Server.GetReturn(resp);
};
shares.GetZtup = async function() {
	let resp = await Server.Post("/shares.get_ztup", {});
	return Server.GetReturn(resp);
};
shares.GetDbszx = async function(days) {
	let resp = await Server.Post("/shares.get_dbszx", {days:days});
	return Server.GetReturn(resp);
};
shares.GetDwfl = async function(days) {
	let resp = await Server.Post("/shares.get_dwfl", {days:days});
	return Server.GetReturn(resp);
};
shares.GetZljk = async  function(days) {
	let resp = await Server.Post("/shares.get_zljk", {days:days});
	return Server.GetReturn(resp);
};

shares.GetHyTime = async function() {
	let resp = await Server.Post("/shares.get_hy_time", {
		"num": 5
	});
	return Server.GetReturn(resp);
};
shares.GetBs = async function(days) {
	let resp = await Server.Post("/shares.get_bs", {days:days});
	return Server.GetReturn(resp);
};
shares.GetBxDaily = async function(data) { 
	let resp = await Server.Post("/shares.get_bs_daily", data);
	return Server.GetReturn(resp);
};

shares.GetBsHyDaily = async function() {
	let resp = await Server.Post("/shares.get_bs_hy_daily", {
		"num": 5
	});
	return Server.GetReturn(resp);
};

shares.GetLytlm = async function(days) {
	let resp = await Server.Post("/shares.get_lytlm", {days:days});
	return Server.GetReturn(resp);
};

shares.GetZt = async function(days) {
	let resp = await Server.Post("/shares.get_zt", {days:days});
	return Server.GetReturn(resp);
};

shares.GetCx = async  function(days) {
	let resp = await Server.Post("/shares.get_cx", {days:days});
	return Server.GetReturn(resp);
};

shares.GetMa = async function(days) {
	let resp = await Server.Post("/shares.get_ma", {days:days});
	return Server.GetReturn(resp);
};

shares.GetZc = async function(days) {
	let resp = await Server.Post("/shares.get_zc", {days:days});
	return Server.GetReturn(resp);
};
shares.GetK30 = async function(days){
	let resp = await Server.Post("/shares.get_k30", {days:days});
	return Server.GetReturn(resp);
};
shares.GetYxr = async function(days){
	let resp = await Server.Post("/shares.get_yxr", {days:days});
	return Server.GetReturn(resp);
};
shares.GetYxrLine = async function(days){
	let resp = await Server.Post("/shares.get_yxr_line", {days:days});
	return Server.GetReturn(resp);
};
shares.GetAixg = async function(code){
	let resp = await Server.Post("/shares.get_aixg", {code:code});
	return Server.GetReturn(resp);
};
shares.GetDailyCheck = async function() {
	let resp = await Server.Post("/shares.get_daily_check", {});
	return Server.GetReturn(resp);
};

shares.GetHyRm = async function(days) {
	let resp = await Server.Post("/shares.get_hy_rm", {
		Days: days
	});
	return Server.GetReturn(resp);
};



shares.GetMyYd = async function(day,tag) {
	let resp = await Server.Post("/shares.get_my_yd", {
		day: day,
		tag:tag,
	});
	return Server.GetReturn(resp);
};

shares.GetMyTeam = async function() {
	let resp = await Server.Post("/shares.get_my_team", {});
	return Server.GetReturn(resp);
};

shares.GetMyTeamDetail = async function(name) {
	let resp = await Server.Post("/shares.get_my_team_detail", {
		name: name
	});
	return Server.GetReturn(resp);
};

shares.GetHyMmadd = async function(data) { // 密码加仓 tag = 1;// 类型 0:机构,1:北上
	let resp = await Server.Post("/shares.get_hy_mmadd", data);
	return Server.GetReturn(resp);
};
shares.GetBsjgMmadd = async function(data) { // 密码加仓 tag = 1;// 类型 0:机构,1:北上
	let resp = await Server.Post("/shares.get_bsjg_mmadd", data);
	return Server.GetReturn(resp);
};


shares.GetHyDj = async function(data) { // 密码加仓 tag = 1;// 类型 0:机构,1:北上
	let resp = await Server.Post("/shares.get_hy_dj", data);
	return Server.GetReturn(resp);
};

shares.GetHyDjCode = async function(data) { // 密码加仓 tag = 1;// 类型 0:机构,1:北上
	let resp = await Server.Post("/shares.get_hy_dj_code", data);
	return Server.GetReturn(resp);
};


shares.GetHyZyb = async function(data) { // 中意榜 tag = 1;// 类型 1:北上,2:公募
	let resp = await Server.Post("/shares.get_hy_zyb", data);
	return Server.GetReturn(resp);
};

shares.GetBsjgZyb = async function(data) { // 中意榜 tag = 1;// 类型 1:北上,2:公募
	let resp = await Server.Post("/shares.get_bsjg_zyb", data);
	return Server.GetReturn(resp);
};
shares.GetBsjgCode = async function(data) { // 中意榜 tag = 1;// 类型 1:北上,2:公募
	let resp = await Server.Post("/shares.get_bsjg_code", data);
	return Server.GetReturn(resp);
};

shares.GetSharesKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_shares_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};
shares.GetSharesKlineMore = async function(code,month) { 
	let resp = await Server.Post("/shares.get_shares_kline_more", {code:code,month:month});
	return Server.GetReturn(resp);
};
shares.GetHyKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_hy_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};
shares.GetGzKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_gz_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};
shares.GetHyGzKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_hy_gz_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};
shares.GetAllHyCodeName = async function() {
	let resp = await Server.Post("/shares.get_all_hy_code_name", {});
	return Server.GetReturn(resp);
};

shares.GetAllYzCodeName = async function() {
	let resp = await Server.Post("/shares.get_all_yz_code_name", {});
	return Server.GetReturn(resp);
};

shares.GetAllBsjgCodeName = async function() {
	let resp = await Server.Post("/shares.get_all_bsjg_code_name", {});
	return Server.GetReturn(resp);
};

shares.GetHotHyName = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_hy_name", {names:names,days:days});
	return Server.GetReturn(resp);
};
shares.GetHotYzName = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_yz_name", {names:names,days:days});
	return Server.GetReturn(resp);
};


shares.GetHotHyCodes = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_hy_codes", {names:names,days:days});
	return Server.GetReturn(resp);
};
shares.GetHotYzCodes = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_yz_codes", {names:names,days:days});
	return Server.GetReturn(resp);
};


shares.GetHotHyNameEx = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_hy_name_ex", {names:names,days:days});
	return Server.GetReturn(resp);
};

shares.GetHotHyCodesEx = async function(names,days)  {
	let resp = await Server.Post("/shares.get_hot_hy_codes_ex", {names:names,days:days});
	return Server.GetReturn(resp);
};

shares.GetFundKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_fund_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};

shares.GetJgKline = async function(code,month) { 
	let resp = await Server.Post("/shares.get_jg_kline", {code:code,month:month});
	return Server.GetReturn(resp);
};

shares.GetYyq = async function(days)  {
	let resp = await Server.Post("/shares.get_yyq", {days:days});
	return Server.GetReturn(resp);
};

shares.GetHySharesTop = async function(data)  {
	let resp = await Server.Post("/shares.get_hy_shares_top", data);
	return Server.GetReturn(resp);
};

shares.GetMrtList = async function(isMy){
	let resp = await Server.Post("/shares.get_mrt_list", {isMy:isMy});
	return Server.GetReturn(resp);
}
shares.GetMrtCode = async function(mrt,code){
	let resp = await Server.Post("/shares.get_mrt_code", {mrt:mrt,code:code});
	return Server.GetReturn(resp);
}

shares.UpsetMrtCode = async function(data){
	let resp = await Server.Post("/shares.upset_mrt_code", data);
	return Server.GetReturn(resp);
}

shares.GetMrtDetail = async function(mrt){
	let resp = await Server.Post("/shares.get_mrt_detail", {mrt:mrt});
	return Server.GetReturn(resp);
}

shares.GetZybHyKline = async function(days,showCd){
	let resp = await Server.Post("/shares.get_zyb_hy_kline", {days:days,showCd:showCd});
	return Server.GetReturn(resp);
}

shares.GetZybZsKline = async function(days,showCd){
	let resp = await Server.Post("/shares.get_zyb_zs_kline", {days:days,showCd:showCd});
	return Server.GetReturn(resp);
}

shares.GetHejjw = async function(dayStr){
	let resp = await Server.Post("/shares.get_hejjw", {
		dayStr:dayStr
	});
	return Server.GetReturn(resp);
};

shares.GetClmx = async function() {
	let resp = await Server.Post("/shares.get_clmx", {});
	return Server.GetReturn(resp);
};

shares.GetYhxg = async function(hexinv) {
	let resp = await Server.Post("/shares.get_yhxg", {hexinv:hexinv});
	return Server.GetReturn(resp);
};


shares.GetYearLine = async function(hexinv) {
	let resp = await Server.Post("/shares.get_year_line", {hexinv:hexinv});
	return Server.GetReturn(resp);
};

shares.GetProduct = async function(pageNumber) {
	let resp = await Server.Post("/order.get_product", {pageNumber:pageNumber});
	return Server.GetReturn(resp);
};

shares.GetOneHyTime = async function(code) {
	let resp = await Server.Post("/shares.get_one_hy_time", {code:code});
	return Server.GetReturn(resp);
};

shares.GetZybYb = async function(days) {
	let resp = await Server.Post("/shares.get_zyb_yb", {
		Days: days
	});
	return Server.GetReturn(resp);
};


//////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////analy////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
shares.AnalyCode = async function(code) {
	let resp = await Server.Post("/analy.analy_code", {
		code: code
	});
	return Server.GetReturn(resp);
};
shares.GetArtic = async function(code,len) {
	let resp = await Server.Post("/analy.get_artic", {
		code: code,
		len:len,
	});
	return Server.GetReturn(resp);
};

shares.GetHy = async function(days){
	let resp = await Server.Post("/analy.get_hy", {days:days});
	return Server.GetReturn(resp);
};

shares.HySearch = async function(code, tag) {
	let resp = await Server.Post("/analy.hy_search", {
		code: code,
		tag: tag
	});
	return Server.GetReturn(resp);
};
shares.GetHyAllCodeName = async function() {
	let resp = await Server.Post("/analy.get_all_code_name", {});
	return Server.GetReturn(resp);
};

shares.GetHyCodes = async function(code,days) {
	let resp = await Server.Post("/analy.get_hy_codes", {
		code: code,
		days:days,
	});
	return Server.GetReturn(resp);
};

shares.AnalyCode = async function(code) {
	let resp = await Server.Post("/analy.analy_code", {
		code: code
	});
	return Server.GetReturn(resp);
};
shares.GetAllSp = async function(code) {
	let resp = await Server.Post("/analy.get_all_sp", {
		code: code
	});
	return Server.GetReturn(resp);
};
shares.GetWdjDetail = async function() {
	let resp = await Server.Post("/analy.get_wdj_detail", {});
	return Server.GetReturn(resp);
};
shares.GetBxNxKlineInfo = async function() {
	let resp = await Server.Post("/analy.get_bx_nx_kline_info", {});
	return Server.GetReturn(resp);
};
shares.GetSampleWdj = async function(hexinv) {
	let resp = await Server.Post("/analy.get_sample_wdj", {
		hexinv: hexinv,
	});
	return Server.GetReturn(resp);
};
shares.GetGjzb = async function() {
	let resp = await Server.Post("/analy.get_gjzb", {});
	return Server.GetReturn(resp);
};

shares.GetTouTiao = async function(){
	let resp = await Server.Post("/analy.get_tou_tiao", {});
	return Server.GetReturn(resp);
};



shares.GetHyRmEx = async function(days) {
	let resp = await Server.Post("/analy.get_hy_rm", {
		Days: days
	});
	return Server.GetReturn(resp);
};

shares.GetCjrl = async function(day) {
	let resp = await Server.Post("/analy.get_cjrl", {day:day});
	return Server.GetReturn(resp);
};

shares.GetHyMmaddEx = async function(data) { // 密码加仓 tag = 1;// 类型 0:机构,1:北上
	let resp = await Server.Post("/analy.get_hy_mmadd", data);
	return Server.GetReturn(resp);
};

shares.GetShareMain = async function(code) { 
	let resp = await Server.Post("/analy.get_share_main", {code:code});
	return Server.GetReturn(resp);
};

shares.GetShareCjl = async function(code,days) { 
	let resp = await Server.Post("/analy.get_share_cjl", {code:code,days:days});
	return Server.GetReturn(resp);
};

shares.GetZcyl = async function(code) { 
	let resp = await Server.Post("/analy.get_zcyl", {code:code});
	return Server.GetReturn(resp);
};

shares.GetGdData = async function(code,tag) {
	let resp = await Server.Post("/analy.get_gd_data", {code:code,tag:tag});
	return Server.GetReturn(resp);
};

shares.GetShareCjl1 = async function(codes,days) { 
	let resp = await Server.Post("/analy.get_share_cjl1", {codes:codes,days:days});
	return Server.GetReturn(resp);
};

shares.GetUsdcny = async function(days) { 
	let resp = await Server.Post("/analy.get_usdcny", {days:days});
	return Server.GetReturn(resp);
};
//////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////analy////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
export default shares
