/**
初始化infra-mongo项目
 */
package basemongo

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-mongo/config"
	"github.com/anypick/infra/base/props/container"
)

func Init() {
	container.Add(&config.MongoConfig{Prefix: config.MongoPrefix})
	infra.Register(&MongoStarter{})
}
