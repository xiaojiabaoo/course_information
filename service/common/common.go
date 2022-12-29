package common

import (
	dto "course_information/dto/common"
	"course_information/models"
	"course_information/models/do"
	"course_information/models/model"
	"course_information/pkg"
	myErr "course_information/pkg/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

type Common interface {
	Login(param dto.Login) (model.Login, error)
	Verification(param dto.Verification) error
	GetUserInfo(c *gin.Context) (int, error)
}

type ApiCommon struct{}

func (a *ApiCommon) Login(param dto.Login) (model.Login, error) {
	var (
		err        error
		loginParam = model.Login{}
		uuid       = pkg.CreateUUID()
		now        = time.Now()
		code       string
		userData   = &do.Users{}
		apiCommon  = models.ApiCommon{}
		apiUser    = models.ApiUser{}
	)

	switch param.LoginType {
	case pkg.LOGIN_TYPE_PASS:
		//查询账号密码是否输入正确
		userData, err = apiCommon.GetUserDataByAccPass(param.Account, param.Password)
		if err != nil {
			return loginParam, err
		}
		//查询不到账号信息，账号不存在或密码错误
		if userData.Id == pkg.EMPTY_INT {
			return loginParam, myErr.CustomError(myErr.AccountPassError)
		}
		//登录成功，把token信息存入redis
		err = a.login(userData.Id, uuid)
		if err != nil {
			return loginParam, err
		}
	case pkg.LOGIN_TYPE_VERIFY:
		//检测验证码是否正确
		code, err = models.GetStr(fmt.Sprintf(pkg.SECURITY_CODE_KEY, param.Account))
		if err != nil {
			return loginParam, errors.Wrap(err, "获取验证码失败")
		}
		if param.SecurityCode != pkg.Str2Int(code) {
			//验证码错误
			return loginParam, myErr.CustomError(myErr.SecurityCodeError)
		}
		//验证码通过后查询账号信息
		userData, err = apiCommon.GetUserDataByAccount(param.Account)
		if err != nil {
			return loginParam, err
		}
		if userData.Id == pkg.EMPTY_INT {
			//没有账号-注册
			users := &do.Users{
				Email:         param.Account,
				Password:      fmt.Sprintf(pkg.USER_DEFAULT_PASS, param.Account),
				AddTime:       int(now.Unix()),
				UpdateTime:    int(now.Unix()),
				DateTime:      now,
				AccountStatus: 0,
			}
			err = apiUser.AddUsers(users)
			if err != nil {
				return loginParam, err
			}
			//添加用户参数
			userParam := &do.UserParam{
				UserId:   users.Id,
				TryCount: pkg.USER_DEFAULT_TRY_COUNT,
				TryTime:  pkg.USER_DEFAULT_TRY_TIME,
				AddTime:  int(now.Unix()),
				UpTime:   int(now.Unix()),
				DateTime: now,
			}
			err = apiUser.AddUserParam(userParam)
			if err != nil {
				return loginParam, err
			}
			userData.Id = users.Id
		}
		//登录
		err = a.login(userData.Id, uuid)
		if err != nil {
			return loginParam, err
		}
	default:
		return loginParam, myErr.CustomError(myErr.AccountLoginTypeError)
	}
	//添加登录日志
	logs := &do.UserLoginLogs{
		UserId:      userData.Id,
		LoginTime:   int(now.Unix()),
		LoginIp:     "127.0.0.1",
		LoginArea:   "中国大陆",
		LoginMethod: param.LoginWay,
		DateTime:    now,
	}
	err = apiUser.AddUserLoginLogs(logs)
	if err != nil {
		return loginParam, err
	}
	loginParam.UserId = userData.Id
	loginParam.Account = param.Account
	loginParam.NickName = userData.NickName
	loginParam.Logo = userData.Logo
	loginParam.Token = uuid
	return loginParam, nil
}

func (a *ApiCommon) login(userId int, uuid string) error {
	err := models.SetStr(fmt.Sprintf(pkg.USER_LOGIN_KEY, userId), uuid, pkg.USER_ONLINE_EXPIRE)
	err = models.SetInt(fmt.Sprintf(pkg.USER_LOGIN_TOKEN_KEY, uuid), userId, pkg.USER_ONLINE_EXPIRE)
	if err != nil {
		return errors.Wrap(err, "用户登录redis存储token失败")
	}
	return nil
}

func (a *ApiCommon) Verification(param dto.Verification) error {
	var (
		body    string
		err     error
		codeNum = pkg.RandNum(100000, 999999) //生成验证码
	)
	if param.Account == "" {
		return myErr.CustomError(myErr.AccountNotEmpty)
	}
	if !pkg.VerifyEmail(param.Account) {
		return myErr.CustomError(myErr.AccountFormatError)
	}

	switch param.SendType {
	case pkg.VERIFY_CODE_REGISTER:
		body = fmt.Sprintf(pkg.REGISTER_VERIFY, codeNum)
	case pkg.VERIFY_CODE_LOGIN:
		body = fmt.Sprintf(pkg.LOGIN_VERIFY, codeNum)
	default:
		return myErr.CustomError(myErr.SendTypeError)
	}
	err = pkg.SendMail([]string{param.Account}, body, codeNum)
	if err != nil {
		return err
	}
	err = models.SetStr(fmt.Sprintf(pkg.SECURITY_CODE_KEY, param.Account), codeNum, pkg.SECURITY_CODE_EXPIRE)
	if err != nil {
		return errors.Wrap(err, "发送验证码redis存储验证码失败")
	}
	return nil
}

func (a *ApiCommon) GetUserInfo(c *gin.Context) (int, error) {
	token := c.GetHeader("token")
	if token == "" {
		return 0, myErr.CustomError(myErr.TokenError)
	}
	str, err := models.GetStr(fmt.Sprintf(pkg.USER_LOGIN_TOKEN_KEY, token))
	if err != nil {
		return 0, errors.Wrap(err, "根据token获取用户信息失败")
	}
	userId := pkg.Str2Int(str)
	if userId == 0 {
		return 0, myErr.CustomError(myErr.AccountLoginError)
	}
	return userId, nil
}
