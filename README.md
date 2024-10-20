# orange
#### 简介


- orange 是一个集成了 自动生成代码、Gin的开发框架,主打简洁、易上手，为初学golang服务端研发的同学设计。
- orange 拥有简洁的生成代码命令，可以从mysql、pg等数据库的表接口，一键生成代码。
- orange 大幅提高了开发效率和降低了开发难度，是很适合新手学习的。


##### 主要功能

orange 包含常用的组件(按需使用)：

```text
Web 框架 gin
RPC 框架 grpc （开发中）
配置解析 viper  
配置中心 nacos 
日志 slog 
数据库组件 gorm
缓存组件 go-redis
自动化api文档 swagger
鉴权 jwt
校验 validator （建议手写）
消息队列组件 rabbitmq, kafka
分布式锁 redis lock
自适应限流 ratelimit
服务注册与发现 etcd, consul, nacos
持续集成部署 CICD jenkins, docker, kubernetes
```

#### 快速开始

1. 安装xt 工具
```
go install github.com/mooncake9527/xt
```

2. 创建项目，例如helloworld 
```
xt new helloworld
```

3.修改配置文件
修改resources/config.dev.yaml 配置文件，注意修改redis连接、数据库连接、日志路径等

4.启动项目
```
cd helloworld
go mod tidy
go run main.go start -c resources/config.dev.yaml

或者 编译之后

go build -o helloworld
./helloworld start -c resources/config.dev.yaml
```

5. 服务启动
你会看到一个大大的皮卡丘
像这样
![](./doc/image.png)
下面还打印了服务监听端口和swagger的地址

#### 直接clone orange 进行开发

[使用orange脚手架快速开发项目](./doc/how%20to%20use.md)





