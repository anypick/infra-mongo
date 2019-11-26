package basemongo

import (
	"context"
	"fmt"
	"github.com/anypick/infra"
	"github.com/anypick/infra-mongo/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

// 设置默认值
const (
	defaultMaxPoolSize uint64 = 100
	defaultMinPoolSize uint64 = 10
)

var client *mongo.Client

func GetClient() *mongo.Client {
	return client
}

type MongoStarter struct {
	infra.BaseStarter
}

// 装配连接, 装配必须使用用户名和密码, TODO 适用其他情况的装配
func (MongoStarter) Setup(ctx infra.StarterContext) {
	var (
		clientOptions *options.ClientOptions
		conf          = ctx.Yaml()[config.MongoPrefix].(*config.MongoConfig)
		body          string // mongodb连接主体， 即：username:password@127.0.0.1:27017
		maxPoolSize   = defaultMaxPoolSize
		minPoolSize   = defaultMinPoolSize
		addrs         []string
		err           error
	)
	// mongodb://username:password@127.0.0.1:27017,username:password@127.0.0.1:27017,username:password@127.0.0.1:27017/?safe=true;w=2;wtimeoutMS=2000
	addrs = strings.Split(conf.Addr, ",")
	for _, k := range addrs {
		body = fmt.Sprintf("%s:%s@%s", conf.Username, conf.Password, k)
	}
	clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb://%s%s", body, conf.Params))
	if conf.MaxPoolSize != 0 {
		maxPoolSize = uint64(conf.MaxPoolSize)
	}
	if conf.MinPoolSize != 0 {
		minPoolSize = uint64(conf.MinPoolSize)
	}
	clientOptions.MaxPoolSize = &maxPoolSize
	clientOptions.MinPoolSize = &minPoolSize
	if client, err = mongo.NewClient(clientOptions); err != nil {
		panic(err)
	}
}

func (MongoStarter) Stop(ctx infra.StarterContext) {
	if err := client.Disconnect(context.Background()); err != nil {
		logrus.Error("mongo connect close fail, ", err)
		return
	}
	logrus.Info("mongo disconnect success")
}
