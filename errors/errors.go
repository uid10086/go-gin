package errors


type Error struct {
	code int
	msg  string
}

//系统通用返回
const (
	ErrorCodePanic          = -1
	ErrorCodeUserAuthReq    = 401
	ErrorCodeServerInternal = 500
	ErrorCodeBadRpc         = 601
	ErrorCodeOk             = iota + 1000
	ErrorCodeFailed
	ErrorCodeParmeterFailed
	ErrorCodeTokenError
)

//用户错误
const (
	ErrorCodeUser              = ErrorCodeOk + 1000
	ErrorCodeUser_ErrorUserId  = ErrorCodeOk + 1001 //用户id错误
	ErrorCodeUser_NoINSchool   = ErrorCodeOk + 1002 //用户不属于任何机构
	ErrorCodeUser_NoFindSchool = ErrorCodeOk + 1003 //学校机构不存在
	ErrorCodeUser_PwdError     = ErrorCodeOk + 1004 //用户账号或密码错误
	ErrorCodeUser_NoUser       = ErrorCodeOk + 1005 //用户不存在
	ErrorCodeUser_NoUserExtend = ErrorCodeOk + 1006 //用户扩展不存在
	ErrorCodeUser_HaveSchool   = ErrorCodeOk + 1007   //此用户已属于其他学校机构
)

//课程错误
const (
	ErrorCodeStream = ErrorCodeOk + 2000
)

//数据库错误
const (
	ErrorCodeDb = ErrorCodeOk + 3000
)
const (
	//机构错误
	ErrorCodeOrganization = ErrorCodeOk + 4000
	//团队错误
	ErrorCodeTeam = ErrorCodeOk + 4100
	//卡券错误
	ErrorCodeCard = ErrorCodeOk + 4200
	//mqtt错误
	ErrorCodeMqtt = ErrorCodeOk + 4300
	//图文错误
	ErrorCodeArticle                  = ErrorCodeOk + 4400
	ErrorCodeArticle_ShareWithoutBind = ErrorCodeArticle + 1 //分享尚未绑定
	ErrorCodeArticle_NotFound = ErrorCodeArticle + 2 //该图文已不存在
	ErrorCodeArticle_Comment_NotFound = ErrorCodeArticle + 3 //该评论已不存在
	ErrorCodeArticle_Commit_Fail = ErrorCodeArticle + 4 //审核处理失败


	//提问错误
	ErrorCodeQuestion = ErrorCodeOk + 4500
	ErrorCodeQuestion_NotFound = ErrorCodeQuestion + 1//该提问已不存在
	ErrorCodeQuestion_Forbid = ErrorCodeQuestion_NotFound + 1//您已被禁言

	//微信第三方平台
	ErrorCodeWXThirdPlatform = ErrorCodeOk + 4600
)

const (
	WXThirdPlatform_UnAuthorization = ErrorCodeWXThirdPlatform + 1
	WXThirdPlatform_Default         = ErrorCodeWXThirdPlatform + 2
)

//消息错误
const (
	ErrorCodeMsg     = ErrorCodeOk + 5000
	ErrorCodeHistory = ErrorCodeOk + 5500
)

type Errors []Error

func (this Errors) Getmsg(code int) string {
	for _, v := range defs {
		if v.code == code {
			return v.msg
		}
	}
	return "未知错误"
}

//错误定义中文描述
var defs = Errors{
	{ErrorCodePanic, "系统Panic"},
	{ErrorCodeOk, "执行成功"},
	{ErrorCodeFailed, "系统错误"},
	{ErrorCodeParmeterFailed, "接口参数错误"},
	{ErrorCodeUser, "用户错误"},
	{ErrorCodeUserAuthReq, "需要认证"},
	{ErrorCodeBadRpc, "服务不可达"},
	{ErrorCodeServerInternal, "服务器内部错误"},
	{ErrorCodeStream, "课程错误"},
	{ErrorCodeDb, "数据库错误"},
	{ErrorCodeOrganization, "机构错误"},
	{ErrorCodeTeam, "团队错误"},
	{ErrorCodeMsg, "验证码错误或已失效"},
	{ErrorCodeCard, "邀请卡错误"},
	{ErrorCodeMqtt, "Mqtt错误"},
	{ErrorCodeArticle, "图文类错误"},
	{ErrorCodeArticle_ShareWithoutBind, "分享前须先绑定三方账户"},
	{ErrorCodeTokenError, "token错误"},
	{ErrorCodeQuestion, "提问类错误"},
	{ErrorCodeHistory, "历史记录生成错误"},
	{ErrorCodeUser_PwdError, "用户账号或密码错误"},
	{ErrorCodeQuestion_NotFound,"该提问已不存在"},
	{ErrorCodeHistory,"历史记录生成错误"},
	{ErrorCodeUser_PwdError,"用户账号或密码错误"},
	{ErrorCodeArticle_NotFound,"该图文已不存在"},
	{ErrorCodeArticle_Comment_NotFound,"该评论已不存在"},
	{ErrorCodeUser_NoUser,"用户不存在"},
	{ErrorCodeQuestion_Forbid,"您已被禁言，请稍后重试！"},
	{ErrorCodeUser_NoFindSchool,"您还不是高考宝教师用户"},
	{ErrorCodeUser_HaveSchool,"此用户已属于其他学校机构"},
	{ErrorCodeArticle_Commit_Fail,"审核处理失败"},
}
