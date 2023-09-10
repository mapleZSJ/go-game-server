package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
    "golang.org/x/net/context"
	pb "goGame/proto"
)
​
const (
    address = "localhost:50051"
)

func userLogin(ctx context.Context, client pb.CSProtoClient) (scinfo pb.SCInfo) {
	req := &pb.CSUserLoginReq{userName: "u1"}
	SCRsp, err := client.userLogin(ctx, req)
	if err != nil {
	   log.Println("user login fail.", err)
	   return
	}
	log.Println("user login success")
	return SCRsp.scRsp
}


func main() {
	//连接grpc服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}

	//延迟关闭连接
	defer conn.Close()
	
	//初始化客户端
	client := pb.NewCSProtoClient(conn)
	ctx := context.Background()
	
	rsp := userLogin(ctx, client)
	//SCRsp, err := client.roleLogin(ctx, &pb.CSRoleLoginReq{user_id: 101, role_id:10001})
	//SCRsp, err := client.roleLogout(ctx, &pb.CSRoleLogoutReq{user_id: 101})
	//SCRsp, err := client.userLogout(ctx, &pb.CSUserLogoutReq{user_id: 101})

	//SCRsp, err := client.queryUserAllRole(ctx, &pb.CSQueryUserAllRoleReq{user_id: 102})
	//SCRsp, err := client.queryUserRole(ctx, &pb.CSQueryUserRoleReq{user_id: 102})  //{user_id: 102, role_id:xx}
	//SCRsp, err := client.batchQueryUserAllRole(ctx, &pb.CSBatchQueryUserAllRoleReq{user_ids : {101, 102}})
}

