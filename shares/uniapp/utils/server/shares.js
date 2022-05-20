/* 用户信息相关 */
import Server from './def' 

const shares = {}

shares.GetGroup =async getGroup=>{
	let resp = await Server.Post("/shares.get_group",{});
	return Server.GetReturn(resp);
};

shares.GetMyGroup =async function(code){
	let resp = await Server.Post("/shares.get_my_group",{code:code});
	return Server.GetReturn(resp);
};

shares.AddGroup =async function(key){
	let resp = await Server.Post("/shares.add_group",{key:key});
	return Server.GetReturn(resp);
};

shares.UpsetGroupCode =async function(code,groupName,userName,isAdd){
	let resp = await Server.Post("/shares.upset_group_code ",{code:code,groupName:groupName,userName:userName,isAdd:isAdd});
	return Server.GetReturn(resp);
};

shares.Gets =async function(codes){
	let resp = await Server.Post("/shares.gets",{codes:codes});
	return Server.GetReturn(resp);
};

shares.GetAllCodeName =async getAllCodeName =>{
	let resp = await Server.Post("/shares.get_all_code_name",{});
	return Server.GetReturn(resp);
};

shares.Search =async function(code,tag) {
	let resp = await Server.Post("/shares.search",{code:code,tag:tag});
	return Server.GetReturn(resp);
};

shares.AddMyCode =async function(req){
	let resp = await Server.Post("/shares.add_my_code",req);
	return Server.GetReturn(resp);
};

shares.GetMyCode = async function(code){
	let resp = await Server.Post("/shares.get_my_code",{code:code});
	return Server.GetReturn(resp);
};

shares.GetMsg = async function(){
	let resp = await Server.Post("/shares.get_msg",{});
	return Server.GetReturn(resp);
};

shares.HaveNewMsg = async function(){
	let resp = await Server.Post("/shares.have_new_msg",{});
	return Server.GetReturn(resp);
};

shares.DeleteMyCode = async function(code){
	let resp = await Server.Post("/shares.delete_my_code",{code:code});
	return Server.GetReturn(resp);
};


export default shares