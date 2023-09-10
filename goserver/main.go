package main

import (
    "context"
	"log"
    "github.com/gofrs/uuid"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "protobuf/csProto"
	"database/mgo"
)

const (
	address = "127.0.0.1:50051"
)

type server struct {
	roleMap map[int]int
	userMap map[int]bool
}



func _get_user_map (svr *server) {
	if svr.userMap == nil {
		svr.userMap = make(map[int]bool) 
	}
	return svr.userMap
}

func _set_user_map (svr *server, uid) {
	if svr.userMap == nil {
		svr.userMap = make(map[int]bool)
	}
	svr.userMap[uid] = true
}

func _get_role_map (svr *server) {
	if svr.roleMap == nil {
		svr.roleMap = make(map[int]int) 
	}
	return svr.roleMap
}

func _set_role_map (svr *server, uid, rid) {
	if svr.roleMap == nil {
		svr.roleMap = make(map[int]int)
	}
	svr.roleMap[uid] = rid
}


//用户注册
/*
func() {
	out, err := uuid.NewV4()
	if err != nil {
		return rsp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}
	//user_name不允许重名 role_id全服唯一
	//user_id、role_id
}
*/


//用户登录
func (svr *server) userLogin(ctx context.Context, req *csProto.CSUserLoginReq) (rsp *csProto.SCRsp, err error) {
	userMap := _get_user_map(svr)

	uname = req.userName
	userInfo := mgo.QueryUserByName(uname)
	uid := userInfo.uid
	if (uid == nil || userMap[uid] != nil) {
		return &csProto.SCRsp{
			errcode : csProto.errno.USER_ALREADY_LOGIN
		}
	}
	
	user_roles := mgo.QueryUser(uid)
	roles := []mgo.RoleInfo{}
	for _, uinfo := range user_roles {
		roles[len(roles)] = {
			role_id : uinfo.rid,
			user_id : uinfo,
			name : uinfo.Name,
			zone_id = uinfo.zid,
		}
	}

	_set_user_map(svr, uid)
	return &csProto.SCRsp{
		errcode : csProto.errno.SUCCESS
		scRsp : {
			userLoginRsp : {
				role_list : roles,
			}
		}
	}
}


//用户退出
func (svr *server) userLogout(ctx context.Context, req *csProto.CSUserLogoutReq) (rsp *csProto.SCRsp, err error) {
	uid := req.user_id
	userMap := _get_user_map(svr)
	if userMap[uid] == nil {
		return &csProto.SCRsp{
			errcode : csProto.errno.USER_UNLOGIN
		}
	}

	rsp = &csProto.SCRsp{}

	delete(svr.userMap, uid)
	return rsp
}

//角色登录
func (svr *server) roleLogin(ctx context.Context, req *csProto.CSRoleLoginReq) (rsp *csProto.SCRsp, err error) {

	uid := req.user_id
	userMap := _get_user_map(svr)
	if userMap[uid] == nil {
		return &csProto.SCRsp{
			errcode : csProto.errno.USER_UNLOGIN
		}
	}

	rsp = &csProto.SCRsp{}

	rid := req.role_id
	roleMap := _get_role_map(svr)
	if roleMap[uid] != nil && roleMap[uid] == rid {
		return rsp   //role already login
	}

	//mgo.QueryRole(uid, rid)

	_set_role_map(svr, uid, rid)
	return rsp
}

//角色退出
func (svr *server) roleLogout(ctx context.Context, req *csProto.CSRoleLogoutReq) (rsp *csProto.SCRsp, err error) {
	uid := req.user_id
	userMap := _get_user_map(svr)
	if userMap[uid] == nil {
		return &csProto.SCRsp{
			errcode : csProto.errno.USER_UNLOGIN
		}
	}

	rsp = &csProto.SCRsp{}

	roleMap := _get_role_map(svr)
	if roleMap[uid] == nil {
		return rsp   //role unlogin
	}

	delete(svr.roleMap, uid)
	return rsp
}

//查询用户下的所有角色信息
func (svr *server) queryUserAllRole(ctx context.Context, req *csProto.CSQueryUserAllRoleReq) (rsp *csProto.SCRsp, err error) {
	rsp = &csProto.SCRsp{}

	uid := req.user_id
	//mgo.QueryUser(uid)

	return rsp
}

//查询用户下的单个角色信息
func (svr *server) queryUserRole(ctx context.Context, req *csProto.CSQueryUserRoleReq) (rsp *csProto.SCRsp, err error) {
	rsp = &csProto.SCRsp{}

	uid := req.user_id
	rid := req.role_id
	//mgo.QueryRole(uid, rid)

	return rsp
}

//批量查询多个用户的所有角色信息
func (svr *server) batchQueryUserAllRole(ctx context.Context, req *csProto.CSBatchQueryUserAllRoleReq) (rsp *csProto.SCRsp, err error) {
	rsp = &csProto.SCRsp{}

	uids := req.user_ids
	//mgo.BatchQueryUser(uids)

	return rsp
}

func main() {
	err_message := mgo.ConnectDB("gamedata")
	if err_message != nil {
		log.Println("net listen err ", err) //log.Fatalf("failed to listen: %v", err)
		return
	}
	mgo.GenTestData()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("net listen err ", err) //log.Fatalf("failed to listen: %v", err)
		return
	}
	
	svr := grpc.NewServer()
	csProto.RegisterCSProtoServer(svr, &server{})

    //往grpc服务端注册反射服务
	reflection.Register(s)

	//启动grpc服务
	if err := svr.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}

 }

