package consts

import (
	"fmt"
)

type ReMsg struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func NewReMsg(code int, msg string) ReMsg {
	if _, ok := reMsgMap[code]; ok {
		panic(fmt.Sprintf("ReMsg code:{%d} already exist", code))
	} else {
		reMsgMap[code] = struct{}{}
	}
	return ReMsg{
		code: code,
		msg:  msg,
	}
}

func (r ReMsg) Error() string {
	return r.msg
}

func (r ReMsg) Code() int {
	return r.code
}

var reMsgMap = make(map[int]struct{})

var (
	Ok                  = NewReMsg(200, "Ok")
	Err400              = NewReMsg(400, "参数错误")
	Err500              = NewReMsg(500, "服务器错误")
	ErrToken            = NewReMsg(1001, "Token失效")
	ErrNoToken          = NewReMsg(1023, "登录用户才能访问")
	ErrNoAccessRights   = NewReMsg(1024, "没有访问权限")
	ErrUseLimitUser     = NewReMsg(1026, "您不在访问授权时间内")
	ErrUserNotExist     = NewReMsg(1027, "当前账号不存在，请先注册")
	ErrUserLock         = NewReMsg(1028, "当前账号被冻结")
	ErrPwdNotExist      = NewReMsg(1029, "尚未设置密码，请选择验证码登录")
	ErrMealLimit        = NewReMsg(1030, "成员数已达到套餐的限制，请升级套餐重试")
	ErrAccessApplyExits = NewReMsg(1031, "您已经提交了登录申请，请耐心等待上级授权，通过授权即可登录")
	ErrMpExpire         = NewReMsg(1056, "已过期")

	ErrMissingAppId        = NewReMsg(1060, "appId缺失")
	ErrMissingTimeStamp    = NewReMsg(1061, "timestamp缺失")
	ErrMissingSignature    = NewReMsg(1062, "signature缺失")
	ErrMissingEnvelop      = NewReMsg(1063, "envelop缺失")
	ErrOverMaxTimestamp    = NewReMsg(1064, "timestamp误差过大")
	ErrAppIdInvalid        = NewReMsg(1065, "appId非法")
	ErrEnvelopInvalid      = NewReMsg(1066, "envelop无效")
	ErrSignatureInvalid    = NewReMsg(1067, "signature无效")
	ErrAppSecret           = NewReMsg(1068, "appSecret非法")
	ErrAppIdNotExist       = NewReMsg(1069, "appId不存在")
	ErrCallOverMaxTimes    = NewReMsg(1070, "调用次数超限,请充值")
	ErrInsufficientBalance = NewReMsg(1071, "余额不足,请充值")
	ErrAccountNotExist     = NewReMsg(1072, "账号不存在，请先建账号")
	ErrApiNotOpen          = NewReMsg(1073, "api未开放")

	ErrRePassword                       = NewReMsg(10001, "重复密码不一致")
	ErrPasswordFMT                      = NewReMsg(10002, "密码长度必须在6-24位")
	ErrMobileOrEmail                    = NewReMsg(10003, "必须手机号或者邮箱注册")
	ErrParams                           = NewReMsg(10004, "参数错误")
	ErrVerifyCode                       = NewReMsg(10005, "验证码错误")
	ErrUnLogin                          = NewReMsg(10006, "未登录")
	ErrNotSelectCompany                 = NewReMsg(10007, "未选中企业")
	ErrBind                             = NewReMsg(10008, "绑定失败")
	ErrCompanyNotExist                  = NewReMsg(10009, "企业不存在")
	ErrCompanyUserNotExist              = NewReMsg(10010, "当前企业用户被锁定或者离开")
	ErrUserExist                        = NewReMsg(10011, "账号已经注册，请直接登录")
	ErrRoleOnePerm                      = NewReMsg(10012, "至少选择一项权限")
	ErrMobileLogin                      = NewReMsg(10013, "登录超时，请重试")
	ErrSendFrequent                     = NewReMsg(10014, "发送频率过快，请稍后再试")
	DepartmentCannotDeleted             = NewReMsg(10015, "部门内有员工，不能被删除")
	WrongDepartmentID                   = NewReMsg(10016, "错误的部门id")
	UnauthorizedDeletion                = NewReMsg(10017, "无权删除")
	ApprovalInProgress                  = NewReMsg(10018, "审批中，请耐心等待")
	OnTheEmployeeList                   = NewReMsg(10019, "您已在员工列表内")
	InvalidCode                         = NewReMsg(10020, "无效的邀请码")
	InvalidAccount                      = NewReMsg(10021, "请输入正确的账号")
	UnauthorizedModification            = NewReMsg(10022, "无权修改")
	MemberDoesNotExist                  = NewReMsg(10023, "成员不存在")
	AlreadyApproved                     = NewReMsg(10024, "已经审批过了")
	AccountAlreadyExists                = NewReMsg(10025, "账户已存在")
	NameAlreadyExists                   = NewReMsg(10026, "名称已存在")
	RoleDoesNotExist                    = NewReMsg(10027, "请求角色不存在")
	IdentityAlreadyExists               = NewReMsg(10028, "当前标识已存在")
	NeedsToDeletedFirst                 = NewReMsg(10029, "角色已分配给成员，需要先删除")
	AccountOrPasswordError              = NewReMsg(10030, "账号或密码错误，请重新输入")
	TooMoreTimes                        = NewReMsg(10031, "调用次数超限")
	CompanyNotMatchLedgerPerson         = NewReMsg(10035, "法人与企业信息核验不一致")
	CompanyFourElementsVerificationFail = NewReMsg(10036, "企业四要素不通过")
	CompanyNotFound                     = NewReMsg(10037, "查无此企业")
	LedgerPersonNotFound                = NewReMsg(10038, "查无此法人")
	AvatarTypeEmpty                     = NewReMsg(10043, "头像类型不能为空")
	AvatarTypeUnknown                   = NewReMsg(10044, "未知的头像类型")
	AvatarImgEmpty                      = NewReMsg(10045, "图片不能为空")
	ErrURLFormat                        = NewReMsg(10047, "请填写正确的url格式")
	ErrFEEpCertNo                       = NewReMsg(10050, "企业证件号错误")
	ErrFEEpCertName                     = NewReMsg(10051, "企业名称错误")
	ErrFELegalPersonCertName            = NewReMsg(10052, "企业法人姓名错误")
	ErrFELegalPersonCertNo              = NewReMsg(10053, "企业法人身份证号码错误")
	ErrServiceAlreadyExist              = NewReMsg(10054, "service已存在")
	ErrServiceApiAlreadyExist           = NewReMsg(10055, "接口已存在")
	ErrFingerNoCannotBeEmpty            = NewReMsg(10057, "envId不能为空")
	ErrFingerNotExist                   = NewReMsg(10058, "指纹信息不存在")
	ErrAppVersionExist                  = NewReMsg(10059, "该版本app已经存在")
)
