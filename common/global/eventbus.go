package global

const (
	TopicApiCountIncrease string = "event:api:counter:increase"       // 接口调用计数+1
	TopicAccountApiDeduct string = "event:api:counter:account:deduct" // 接口扣款
	TopicAccountRecharge  string = "event:account:recharge"           // 用户充值
	TopicAccountChange    string = "event:account:change"             // 账户余额变动,会生成balance
	TopicCoreInitFinish   string = "event:init:core:finish"           // 核心服务初始化完成
	TopicApplicationClose string = "event:application:close"          // 应用关闭
	TopicTrafficLog       string = "event:traffic:log"                // 流量日志,记录与其他服务的http
)
