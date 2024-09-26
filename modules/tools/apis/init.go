package apis

import (
	"github.com/mooncake9527/orange-core/common/apis"
	"github.com/mooncake9527/orange/modules/tools/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Init struct {
	apis.NApi
}

// // DoInit 初始化
// // @Summary 初始化
// // @Tags 工具 / 初始化
// // @Accept application/json
// // @Product application/json
// // @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// // @Router /api/v1/tools/doInit [post]
// // @Security Bearer
// func (e *Init) DoInit(c *gin.Context) {
// 	fmt.Println("开始运行初始化")

// 	// service.ImportSql("resources/dbs/orange-db.sql", consts.DbDefault)
// 	// service.ImportSql("resources/dbs/dental-db.sql", "dental")

// 	result := "执行成功"
// 	if err := core.DB().AutoMigrate(); err != nil {
// 		result = "sys执行失败"
// 	}

// 	// t1, err := template.ParseFiles("modules/tools/apis/tmpls/result.html")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// t1.Execute(c.Writer, result)
// 	e.Ok(c, result)

// }

var last time.Time

var server utils.Server

// Monitor 监控
// @Summary 监控
// @Tags 工具 / 监控
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=utils.Server} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/monitor [post]
// @Security Bearer
func (e *Init) Monitor(c *gin.Context) {
	cur := time.Now()
	if cur.Sub(last) < time.Second {
		e.OK(c, server)
		return
	}
	last = cur
	server.Os = utils.InitOS()
	cpu, err := utils.InitCPU()
	if err == nil {
		server.Cpu = cpu
	}
	d, err := utils.InitDisk()
	if err == nil {
		server.Disk = d
	}
	ram, err := utils.InitRAM()
	if err == nil {
		server.Ram = ram
	}
	e.OK(c, server)
}
