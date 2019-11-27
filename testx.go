package basemongo

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra/base/props"
	"github.com/anypick/infra/base/props/container"
	"path/filepath"
	"runtime"
)

// 初始化测试条件
func InitTest() {
	infra.Register(&container.YamlStarter{})
	Init()
	infra.Register(&infra.BaseInitializerStarter{})
	source := props.NewYamlSource(GetCurrentFilePath("./testx/config.yml", 0))
	app := infra.New(*source)
	app.Start()
}

// 获取文件的绝对路径
func GetCurrentFilePath(fileName string, skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	// 解析出文件路径
	dir := filepath.Dir(file)
	return filepath.Join(dir, fileName)
}



