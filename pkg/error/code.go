package error

const (
	//系统相关错误码
	Success     = 0
	ServerError = 99999
	//参数验证相关错误码
	InvalidParams     = 20000
	ParamError        = 20001
	TokenError        = 20002
	SubjectEmptyError = 20003
	PieceEmptyError   = 20004
	//账号问题相关错误码
	AccountExists         = 30000
	AccountNotExists      = 30001
	AccountException      = 30002
	AccountFormatError    = 30003
	AccountPassError      = 30004
	AccountBlockedError   = 30005
	AccountNotEmpty       = 30006
	AccountLoginTypeError = 30007
	AccountLoginError     = 30008
	//验证码问题相关错误码
	SecurityCodeExists = 40000
	SmsSendError       = 40001
	SecurityCodeError  = 40002
	SendTypeError      = 40003
	//用户相关错误码
	UserTryMaxError  = 50000
	UserTrySameError = 50001
)
