package error

var MsgFlags = map[int]string{
	//系统相关错误码
	Success:     "操作成功",
	ServerError: "服务器异常，请稍后再试！",
	//参数验证相关错误码
	InvalidParams:     "参数非法或缺失!",
	ParamError:        "参数输入有误，请重新输入!",
	TokenError:        "token错误或不存在",
	SubjectEmptyError: "请输入课程ID",
	PieceEmptyError:   "请输入块ID",
	//账号问题相关错误码
	AccountExists:         "账号已存在",
	AccountNotExists:      "账号不存在",
	AccountNotEmpty:       "账号不能为空",
	AccountException:      "账号异常",
	AccountFormatError:    "电话或电子邮件格式不正确",
	AccountPassError:      "账号或密码错误",
	AccountBlockedError:   "账号已被拉黑，请联系管理员处理！",
	AccountLoginTypeError: "用户登录类型错误",
	AccountLoginError:     "请登录账号",
	//验证码问题相关错误码
	SecurityCodeExists: "验证码已存在，请勿重复操作",
	SecurityCodeError:  "验证码输入错误",
	SmsSendError:       "短息发送失败",
	SendTypeError:      "短息发送类型错误",
	//用户相关错误码
	UserTryMaxError:  "试用机会已用完，通过积分可兑换试用机会或购买本课程学习",
	UserTrySameError: "课程正在试用中，不能重复试用",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ServerError]
}

type MyError struct {
	err  string
	code int
}

func (m *MyError) Error() string {
	return m.err
}

func (m *MyError) Code() int {
	return m.code
}

func CustomError(code int) error {
	mine := &MyError{
		code: code,
		err:  MsgFlags[code],
	}
	return mine
}
