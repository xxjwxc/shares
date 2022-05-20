/* 管理相关信息 */
import Server from './def';

const admin = {}

// 账号相关
// 登录
admin.Login =async function(username,password){
	let result = await Server.Post("/login/login",{username:username,password:password});
	let resp = Server.GetReturn(result);
	
	if (typeof(resp) == "undefined" || resp == null){ // 失败
		return resp;
	}
	var nowTime = ((new Date()).getTime())
	resp.expire_time = nowTime + (resp.expire_time * 1000);
	uni.setStorage({
		key:"access_token",
		data:JSON.stringify(resp)
	})
	Server.SessionID = resp.access_token;
	
	return resp;
};

// 刷新token
admin.RefreshToken =async function(refresh_token){
	let result = await Server.Post("/login.refresh_token",{refresh_token:refresh_token});
	let resp = Server.GetReturn(result);
	
	if (typeof(resp) == "undefined" || resp == null){ // 失败
		return resp;
	}
	var nowTime = ((new Date()).getTime())
	resp.expire_time = nowTime + (resp.expire_time * 1000);
	uni.setStorage({
		key:"access_token",
		data:JSON.stringify(resp)
	})
	Server.SessionID = resp.access_token;

	return resp;
};

// 设置token
admin.SetAccessToken = async function(access_token){
	console.log(access_token);
	Server.SessionID = access_token;
};

// 订单相关

// 获取订单
admin.GetBillInfoByAdmin =  async function(status,pageNumber,search){
	console.log(Server.SessionID)
	let resp = await Server.Post("/order.get_bill_info_by_admin",{status:status,pageNumber:pageNumber,search:search});
	return Server.GetReturn(resp);
}

// 获取快递公司
admin.GetShipmentPost =  async function(){
	console.log(Server.SessionID)
	let resp = await Server.Post("/order.get_shipment_post",{});
	return Server.GetReturn(resp);
}

// 订单添加运单号
admin.AddBillShipment =  async function(billId,shipmentId,postKey){
	console.log(Server.SessionID)
	let resp = await Server.Post("/order.add_bill_shipment",{billId:billId,shipmentId:shipmentId,postKey:postKey});
	return Server.GetReturn(resp);
}


// 获取售后回复
admin.GetAfterMsg =  async function(billId){
	console.log(Server.SessionID)
	let resp = await Server.Post("/order.get_after_msg",{billId:billId});
	return Server.GetReturn(resp);
}



export default admin