package config

const (
	MongoPrefix = "mongo"
)

type MongoConfig struct {
	Prefix      string
	Addr        string `yaml:"addr"` // 127.0.0.1:27017,127.0.0.1:27018
	DataBase    string `yaml:"database"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxPoolSize int    `yaml:"maxPoolSize"`
	MinPoolSize int    `yaml:"minPoolSize"`
	Params      string `yaml:"params"` // 例子： ?safe=true;w=2;wtimeoutMS=2000
}

func (m *MongoConfig) ConfigAdd(config map[interface{}]interface{}) {
	m.Addr = config["addr"].(string)
	m.DataBase = config["database"].(string)
	m.Username = config["username"].(string)
	m.Password = config["password"].(string)
	m.MaxPoolSize = config["maxPoolSize"].(int)
	m.MinPoolSize = config["minPoolSize"].(int)
	m.Params = config["params"].(string)
}
