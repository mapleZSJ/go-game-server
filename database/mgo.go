package database


import (
    "context"
	"log"
    "github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	address = "127.0.0.1:27017"
)

type UserInfo struct {
	Name   string `bson:"name"`
	uid    uint32 `bson:"uid"`
	//...
}

type RoleInfo struct {
	Name   string `bson:"name"`
	Level  uint32 `bson:"level"`
	rid    uint32 `bson:"rid"`  //role_id
	uid    uint32 `bson:"uid"`  //user_id
	zid    uint32 `bson:"zid"`  //区服id
	//...
}

var db dbtype  //error
var ctx context.Context


func ConnectDB(dbname string)  {
	ctx = context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://" + address})
	//cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://" + address, Database: "gamedata", Coll: "user"})

	if err != nil {
		log.Println("mongoDB connect fail", err)
		return err
	}

	db = client.Database(dbname)
	defer func() {
		if err = client.Close(ctx); err != nil {
			panic(err)
		}
	}()
	return err
}

func GenTestData()  {
	coll := db.Collection("user")   //有问题，暂时先这样写
	var batchUserInfo = []interface{}{
		UserInfo{Name: "u1", uid: 101},
		UserInfo{Name: "u2", uid: 102},
		UserInfo{Name: "u3", uid: 103},
		UserInfo{Name: "u4", uid: 104},
	}
	result, err := coll.InsertMany(ctx, batchUserInfo)
	if err != nil {
		log.Println("gen user test data fail", err)
	}

	coll := db.Collection("role")
	var batchRoleInfo = []interface{}{
		RoleInfo{Name: "s1", Level: 6, rid: 10001, uid: 101, zid: 10001},
		RoleInfo{Name: "s2", Level: 6, rid: 10002, uid: 101, zid: 10002},
		RoleInfo{Name: "s3", Level: 6, rid: 10003, uid: 102, zid: 10001},
		RoleInfo{Name: "s4", Level: 6, rid: 10004, uid: 102, zid: 10001},
		RoleInfo{Name: "s5", Level: 7, rid: 10005, uid: 103, zid: 10001},
		RoleInfo{Name: "s6", Level: 8, rid: 10006, uid: 104, zid: 10002},
	}
	result, err = coll.InsertMany(ctx, batchRoleInfo)
	if err != nil {
		log.Println("gen role test data fail", err)
	}
}

func QueryUserByName(user_name string)  {
	info := UserInfo{}
	client.Find(ctx, bson.M{"name": user_name).All(&info)
	return &info
}

func BatchQueryUser(uids []int32)  {
	batch := []RoleInfo{}
	for _, uid := range uids {
		infos := QueryUser(uid)
		for _, roleInfo := range infos {
			batch[len(batch)] = roleInfo
		}
	}
	return &batch
}

func QueryUser(user_id int32)  {
	batch := []RoleInfo{}
	client.Find(ctx, bson.M{"uid": user_id).All(&batch)
	return &batch
}

func QueryRole(user_id int32, role_id int32)  {
	info := RoleInfo{}
	client.Find(ctx, bson.M{"uid": user_id, "rid": role_id}).One(&info)
	return &info
}

