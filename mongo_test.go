package basemongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestGetClient(t *testing.T) {
	InitTest()

	client := GetClient()
	collection := client.Database("ziidu").Collection("ziidu")
	one := collection.FindOne(context.Background(), bson.D{{"name", "系统管理"}})
	var data map[string]interface{}
	one.Decode(&data)
	fmt.Println(data)
	time.Sleep(time.Millisecond * 1)
}
