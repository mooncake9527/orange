package main

import (
	"github.com/mooncake9527/orange/cmd"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.io,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init --parseDependency --parseDepth=6

// @title orange API
// @version V0.0.1
// @description 致力于做一个开发快速，运行稳定的框架
// @contact.name   victor
// @contact.url    https://github.com/mooncake9527/orange
// @contact.email  tusihao@gmail.com
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
