/* 订单相关信息 */
import Server from './def' 

const order = {}


//  费用计算
order.ReckonFee =async function(buyType,ids,addrId,couponId){
	let resp = await Server.Post("/order.reckon_fee",{buyType:buyType,ids:ids,addrId:addrId,couponId:couponId});
	return Server.GetReturn(resp);
};

// 下单
order.PlaceOrder = async function(buyType,ids,addrId,couponId,desc){
	let resp = await Server.Post("/order.place_order",{buyType:buyType,ids:ids,addrId:addrId,couponId:couponId,desc:desc});
	return Server.GetReturn(resp);
};

order.Pay = async function(billId){
	let resp = await Server.Post("/order.pay",{billId:billId});
	return Server.GetReturn(resp);
}

order.GetMyOrders = async function(status,pageNumber){
	let resp = await Server.Post("/order.get_my_orders",{status:status,pageNumber:pageNumber});
	return Server.GetReturn(resp);
}
order.GetMyOrdersConfig = async function(isAdmin){
	let resp = await Server.Post("/order.get_my_orders_config",{isAdmin:isAdmin});
	return Server.GetReturn(resp);
}

order.DealOrder = async function(billId,type,contact,remak,imgs){
	let resp = await Server.Post("/order.deal_order",{billId:billId,type:type,contact:contact,remak:remak,imgs:imgs});
	return Server.GetReturn(resp);
}

order.DealSystem = async function(billId,type,contact,remak,imgs){
	let resp = await Server.Post("/order.deal_system",{billId:billId,type:type,contact:contact,remak:remak,imgs:imgs});
	return Server.GetReturn(resp);
}


// 优惠券相关
// 获取促销优惠券
order.GetPromotionCoupon = async function(){
	let resp = await Server.Post("/coupon.get_promotion_coupon",{});
	return Server.GetReturn(resp);
}

// 一键获取优惠券
order.GoGetCoupon = async function(ids){
	let resp = await Server.Post("/coupon.go_get_coupon",{ids:ids});
	return Server.GetReturn(resp);
}

// 获取我的优惠券
order.GetMyCoupon = async function(){
	let resp = await Server.Post("/coupon.get_my_coupon",{});
	return Server.GetReturn(resp);
}

order.CheckPay = async function(billId){
	let resp = await Server.Post("/order.check_pay",{billId:billId});
	return Server.GetReturn(resp);
}

order.GetOrdertrackInfo = async function(billId){
	let resp = await Server.Post("/order.get_ordertrack_info",{billId:billId});
	return Server.GetReturn(resp);
}

// --------end

// 调起支付
order.WxPay =async function(info){
		uni.requestPayment({
			provider: 'wxpay',
			timeStamp: info.timeStamp,//时间戳
			nonceStr: info.nonceStr,//随机字符串
			package: info.package,//统一下单接口返回的 prepay_id 参数值
			signType: info.signType,
        	paySign: info.paySign,//签名内容
           success: async function(res) {
			   	let result = await order.CheckPay(info.orderId);// 检测一次
              	uni.redirectTo({
					url:'/pages/money/paySuccess?orderId=' + info.orderId,
				});
           },
           fail: function (err) {
			uni.showToast({
				title: "支付失败，可在我的订单重新发起支付。",
				icon: 'none'
			})	
              console.log('fail:' + JSON.stringify(err));
           },
       });
}


export default order