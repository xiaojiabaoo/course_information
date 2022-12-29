package models

import (
	"course_information/models/do"
	"github.com/pkg/errors"
)

type User interface {
	AddUsers(user *do.Users) error
	AddUserParam(param *do.UserParam) error
	AddUserLoginLogs(logs *do.UserLoginLogs) error
}

type ApiUser struct{}

func (a *ApiUser) AddUsers(user *do.Users) error {
	_, err := db.Insert(user)
	if err != nil {
		return errors.Wrap(err, "添加用户失败")
	}
	return nil
}

func (a *ApiUser) AddUserParam(param *do.UserParam) error {
	_, err := db.Insert(param)
	if err != nil {
		return errors.Wrap(err, "添加用户参数失败")
	}
	return nil
}

func (a *ApiUser) AddUserLoginLogs(logs *do.UserLoginLogs) error {
	_, err := db.Insert(logs)
	if err != nil {
		return errors.Wrap(err, "添加用户登录日志失败")
	}
	return nil
}

func (a *ApiUser) GetUserParam(id int) (*do.UserParam, error) {
	param := &do.UserParam{}
	_, err := db.Where("user_id=?", id).Get(param)
	if err != nil {
		return param, errors.Wrap(err, "获取有用户参数信息失败")
	}
	return param, nil
}