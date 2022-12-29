package models

import (
	"course_information/models/do"
	"github.com/pkg/errors"
)

type Common interface {
	GetUserDataByAccPass(account, password string) (*do.Users, error)
	GetUserDataByAccount(account string) (*do.Users, error)
}

type ApiCommon struct{}

func (a *ApiCommon) GetUserDataByAccPass(account, password string) (*do.Users, error) {
	users := &do.Users{}
	_, err := db.Where("password=?", password).
		And("phone=? or email=?", account, account).Get(users)
	if err != nil {
		return nil, errors.Wrap(err, "根据账号和密码获取用户信息失败")
	}
	return users, err
}

func (a *ApiCommon) GetUserDataByAccount(account string) (*do.Users, error) {
	users := &do.Users{}
	_, err := db.Where("phone=? or email=?", account, account).Get(users)
	if err != nil {
		return nil, errors.Wrap(err, "根据账号获取用户信息失败")
	}
	return users, err
}
