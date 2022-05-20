const server = {}

server.Host = "http://localhost:82/shares/api/v1";
server.Appid = "wx6xxxxxxxxxxxa"
//server.Host = "http://localhost:82/shares/api/v1";

server.SessionID = ""; // 用户sessionid
server.OpenID = ""; // 用户openid
server.UserInfo = {};
server.IsLogin = false;


server.GetUrl = (router) => {
	return server.Host + router;
};

/*   
router：相对路由
data：请求参数
返回：成功或者失败
success:成功 调用
fail:自定义失败调用
*/
server.Post = async (router, data) => {
	return new Promise((resolve, reject) => {
		//显示加载动画
		// uni.showLoading({
		// 	title: '请稍后..',
		// 	mask: true,
		// })

		uni.request({
			url: server.GetUrl(router),
			method: "POST",
			data: data,
			dataType: 'json',
			// header: {
			// 	"Content-Type":"application/json",
			// 	"Cookie": "session_token=" + server.SessionID
			// },
			success: (res) => {
				if (res.statusCode == 200) { // 判断http返回状态 
					// console.log(JSON.stringify(res.data))
					if (res.data.state) { // 正常返回 
						resolve({
							"success": res.data.data,
							"fail": null
						})
					} else { // 错误返回
						let code = res.data.code
						switch (code) {
							case 1018: // 用户不存在
								server.SessionID = "";
							default:
								resolve({
									"success": null,
									"fail": res.data
								});
						}
					}
				} else {
					resolve({
						"success": null,
						"fail": {
							error: res.data
						}
					});
				}
			},
			fail: (err) => {
				reject(err);
				console.log(router + " error:")
				console.log(err)
				uni.showToast({
					title: err,
					icon: 'none'
				})
			}
		})
	})
};

server.Get = async (router, data) => {
	return new Promise((resolve, reject) => {
		uni.request({
			url: router,
			method: 'GET',
			data: data,
			success: res => {
				console.log(JSON.stringify(res))
				resolve({
					"success": res,
					"fail": null
				})
			},
			fail: res => {
				resolve({
					"success": null,
					"fail": res
				});
			},
			complete: () => {}
		});
	})
};


// 上传多个文件
server.UploadFileWx = async (files) => {
	let resp = [];
	for (let i = 0; i < files.length; i++) {
		let tmp = await server.UploadFileEx(files[i]);
		if (tmp.success) { // 设置session
			resp.push(tmp.success[0]);
		}
	}
	return resp;
};

// 上传单个文件
server.UploadFileEx = async (file) => {
	return new Promise((resolve, reject) => {
		let fromData = {
			url: server.GetUrl("/upload"),
			dataType: 'json',
			success: res => {
				console.log(JSON.stringify(res.data))
				let resp = JSON.parse(res.data);
				console.log(JSON.stringify(resp))
				if (resp.state) { // 正常返回 
					resolve({
						"success": resp.data,
						"fail": null
					})
				} else { // 错误返回
					console.log("-----");
					resolve({
						"success": null,
						"fail": resp.data
					});
				}
			},
			fail: res => {
				reject(res);
				uni.showToast({
					title: res,
					icon: 'none'
				})
			}
		};

		fromData.filePath = file.uri;
		fromData.name = file.name;
		uni.uploadFile(fromData);
	});
};

// 上传单个文件
server.UploadFile = async (files) => {
	return new Promise((resolve, reject) => {
		let fromData = {
			url: server.GetUrl("/upload"),
			dataType: 'json',
			success: res => {
				console.log(JSON.stringify(res.data))
				let resp = JSON.parse(res.data);
				if (resp.state) { // 正常返回 
					resolve({
						"success": resp.data,
						"fail": null
					})
				} else { // 错误返回
					console.log("-----");
					resolve({
						"success": null,
						"fail": resp.data
					});
				}
			},
			fail: res => {
				reject(res);
				uni.showToast({
					title: res,
					icon: 'none'
				})
			}
		};

		fromData.files = files;
		uni.uploadFile(fromData);
	});
};

server.GetReturn = (resp) => {
	if (resp.success) { // 设置session
		return resp.success
	}
	if (resp.fail) { // 失败打印输出
		console.log(resp)
		switch (resp.fail.code) {
			case 2:
				break
			default:
			uni.showModal({
				content: resp.fail.error,
				showCancel: false
			});
		}
	}
	return null;
};



export default server;
