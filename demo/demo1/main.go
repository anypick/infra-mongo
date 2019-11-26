package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Company struct {
	CompanyId    string
	Name         string
	Addr         string
	NumberPeople int64
}

func main() {
	var (
		clientOptions *options.ClientOptions
		cancelFunc    context.CancelFunc
		timeoutCtx    context.Context
		client        *mongo.Client
		collection    *mongo.Collection
		err           error
	)

	// 多个host使用 , 分开
	clientOptions = options.Client().ApplyURI("mongodb://192.168.56.136:27017")

	// 设置连接超时
	timeoutCtx, cancelFunc = context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	if client, err = mongo.Connect(timeoutCtx, clientOptions); err != nil {
		fmt.Println(err)
		return
	}

	collection = client.Database("tutorial").Collection("users")
	//fFlag := false
	//tFlag := true
	//CreateIndies(collection, mongo.IndexModel{Keys: bsonx.Doc{{"name", bsonx.Int32(1)}}})
	// Delete(collection)
	//DeleteIndex(collection, "name_1")
	GetIndies(collection)

	//client.Database("").RunCommand()
}

// 创建索引
// keys表示是一个键值对，value是索引字段，value是索引的类型
// 1表示升序建立索引， -1表示降序建立索引
func CreateIndies(collection *mongo.Collection, model mongo.IndexModel) {
	var (
		name string
		err  error
	)
	if name, err = collection.Indexes().CreateOne(context.Background(), model); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
}

func GetIndies(collection *mongo.Collection) {
	if cursor, err := collection.Indexes().List(context.Background()); err != nil {
		fmt.Println(err)
		return
	} else {
		for cursor.Next(context.Background()) {
			fmt.Println(cursor.Current.String())
		}
	}
}

// 删除索引，name为索引名称，而不是字段名称
func DeleteIndex(collection *mongo.Collection, name string) {
	if raws, err := collection.Indexes().DropOne(context.Background(), name); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(raws.String())
	}
}

// 插入数据
func InsertOne(collection *mongo.Collection) {
	var (
		insertResult *mongo.InsertOneResult
		err          error
	)



	if insertResult, err = collection.InsertOne(context.Background(), &Company{Name: "七牛云", Addr: "上海张江",
		NumberPeople: 1000}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(insertResult.InsertedID)
}

func UpdateOne(collection *mongo.Collection) {

}

// 删除数据
func Delete(collection *mongo.Collection) {
	var (
		deleteResult *mongo.DeleteResult
		err          error
	)
	if deleteResult, err = collection.DeleteMany(context.Background(), bson.D{{"name", "七牛云"}}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResult.DeletedCount)
}
