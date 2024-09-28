# 如何使用orange开发项目

### 如何使用orange
1. 下载orange脚手架项目
2. 修改 go.mod 中的项目名称
3. 全局替换 orange项目名称为你的项目名
4. 使用orange的工具生成代码

### 如何生成代码

1. 编译orange 可以使用如下命令
```text
go build -o orange main.go
```

2. 修改配置文件 resources/config.dev.yaml

orange 也支持其他格式的配置文件，只需要你在启动的时候，设定-c参数指定配置文件路径就可以啦

3. 生成代码的例子
shell 命令如下
```text
./orange gen -c resources/config.dev.yaml -m open -d default -t user
```
这个需要你先创建数据库，并且配置在配置文件里。

-m 代表着module，你像生成的模块，这里是open

-d 代表着用哪个数据库，这里指定default默认，也就是配置文件里的open库

-t 代表这哪个表，这里是user表

[user表结构](./use.sql)

运行输出
```shell
2024/09/28 10:56:13 /Users/mac/workspace/chenlian/orange/modules/tools/models/tools/db_tables.go:67
[26.563ms] [rows:1] SELECT * FROM `information_schema`.`tables` WHERE table_schema= 'open'  AND TABLE_NAME = 'user' ORDER BY `tables`.`TABLE_NAME` LIMIT 1
2024/09/28 10:56:13 /Users/mac/workspace/chenlian/orange/modules/tools/models/tools/db_columns.go:78
[8.031ms] [rows:9] SELECT * FROM `information_schema`.`columns` WHERE table_schema= 'open'  AND TABLE_NAME = 'user' ORDER BY ORDINAL_POSITION asc
time=2024-09-28T10:56:13.633+08:00 level=INFO msg=render out=cmd/start/open.go
time=2024-09-28T10:56:13.634+08:00 level=INFO msg=render out=./modules/open/router/router.go
time=2024-09-28T10:56:13.634+08:00 level=INFO msg=render out=./modules/open/models/user.go
time=2024-09-28T10:56:13.635+08:00 level=INFO msg=render out=./modules/open/apis/user.go
time=2024-09-28T10:56:13.635+08:00 level=INFO msg=render out=./modules/open/router/user.go
time=2024-09-28T10:56:13.636+08:00 level=INFO msg=render out=./modules/open/service/dto/user.go
time=2024-09-28T10:56:13.636+08:00 level=INFO msg=render out=./modules/open/service/user.go
time=2024-09-28T10:56:13.636+08:00 level=INFO msg=render out=../orange-admin/src/api/open/user.ts
time=2024-09-28T10:56:13.637+08:00 level=INFO msg=render out=../orange-admin/src/views/open/user/index.vue
time=2024-09-28T10:56:13.637+08:00 level=INFO msg=render out=../orange-admin/src/views/open/user/form.vue
time=2024-09-28T10:56:13.638+08:00 level=INFO msg=render out=../orange-admin/src/views/open/user/utils/hook.tsx

```

前面的sql是数据库里的user 表的信息

后面生成.go文件，以及一些前端文件。

前端文件的路径../orange-admin/src,是可以在配置文件里面修改的。

我们看看生成的go代码
```text
./modules/open/models/user.go 生成的是model
./modules/open/apis/user.go  生成的是api，也可以叫controller 或者有的人也叫handler
./modules/open/service/user.go 生成的是service层代码
./modules/open/router/user.go 生成的是user api的路由配置
./modules/open/router/router.go 当前module的路由配置
./modules/open/service/dto/user.go 生成user相关的dto文件

```

接下来我们看一下各个文件里面的代码吧

先看model

```text
// User 用户表
type User struct {
	Id        uint64     `json:"id"  form:"id"  gorm:"type:bigint unsigned;primaryKey;autoIncrement;comment:主键"`         //主键
	UserId    string     `json:"userId" form:"userId" gorm:"type:varchar(64);comment:用户id"`                              //用户id
	Username  string     `json:"username" form:"username" gorm:"type:varchar(50);comment:用户名"`                           //用户名
	Mobile    string     `json:"mobile" form:"mobile" gorm:"type:varchar(18);comment:手机号"`                               //手机号
	Nickname  string     `json:"nickname" form:"nickname" gorm:"type:varchar(64);comment:昵称"`                            //昵称
	Avatar    string     `json:"avatar" form:"avatar" gorm:"type:varchar(255);comment:头像"`                               //头像
	Email     string     `json:"email" form:"email" gorm:"type:varchar(128);comment:邮箱"`                                 //邮箱
	CreatedAt time.Time  `json:"createdAt" form:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"` //创建时间
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"type:datetime;comment:更新时间"`                           //更新时间
}

const TBUser = "user"

func (User) TableName() string {
	return TBUser
}
```

细心的同学肯定注意到了 json:"id"  form:"id" 等tag,这么做是为了方便你复制到请求参数里面，可以当请求体的字段来用。
第二点是字段后面的舒适，如 "//用户名"  有了这些注释，就可以用swagger一键生成文档啦。
如果你喜欢Yapi这种文档管理工具，是可以直接通过swagger json 文件导入的，用起来相当方便。如果有同学不懂这块的，可以私聊我。我可以写一个这样的文章。

在看看service 

```text
package service

import (
	"github.com/mooncake9527/orange-core/core/base"
	"github.com/mooncake9527/orange/modules/open/models"
	"github.com/mooncake9527/x/xerrors/xerror"
)

type UserService struct {
	*base.BaseService
}

var SerUser = UserService{
	base.NewService("default"),
}

func (x *UserService) Create(data *models.User) error {
	err := x.DB().Create(data).Error
	if err != nil {
		return xerror.New(err.Error())
	}
	return nil
}

func (x *UserService) GetAll(list *[]*models.User) (err error) {
	err = x.DB().Model(&models.User{}).Find(list).Error
	if err != nil {
		return xerror.Wrap(err, "query User all error")
	}
	return nil
}

func (x *UserService) Get(req, u *models.User) (err error) {
	q := x.DB()
	if req.Id != 0 {
		q = q.Where("id=?", req.Id)
	}
	err = q.First(&u).Error
	if err != nil {
		return xerror.Wrap(err, "query User error")
	}
	return nil
}

```

这里只生成了常用的3个函数，另外更常用的增删改查呢？注意看
```text
type UserService struct {
	*base.BaseService
}
```
猜的没错，都封装在BaseService里面了。


再来看看router文件
```text
package router

import (
	"github.com/mooncake9527/orange/modules/open/apis"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserRouter)
}

// 默认需登录认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("user") //.Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiUser.Get)
		r.POST("/create", apis.ApiUser.Create)
		r.POST("/update", apis.ApiUser.Update)
		r.POST("/page", apis.ApiUser.QueryPage)
		r.POST("/del", apis.ApiUser.Del)
	}
}
```
比较简单，就是增删改查的路由配置。

再来看看api
```text
package apis

import (
	...
)

type UserApi struct {
	apis.NApi
}

var ApiUser = UserApi{}

// QueryPage 获取用户表列表
// @Summary 获取用户表列表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Param data body dto.UserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.User}} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/page [post]
// @Security Bearer
func (e *UserApi) QueryPage(c *gin.Context) {
	var req dto.UserGetPageReq
    if err := e.Bind(c, &req); err != nil {
        e.ParamError(c, err)
        return
    }
    
    if err := req.Valid(); err != nil {
        e.ParamError(c, err)
        return
    }
    
	list := make([]models.User, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerUser.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.PageOK(c, list, total, req.GetPage(), req.GetSize())
}

```

api文件，也很简单，就是调用service的函数，然后返回给前端。工具已经帮你生成好了swagger文档，这很方便，再也不用自己写了。

```text
type UserApi struct {
	apis.NApi
}
````

这个apis.NApi封装了api层常用的一些方法。好奇的同学可以看看里面的源码。



