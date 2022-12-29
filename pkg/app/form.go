package app

import (
	errors "course_information/pkg/error"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, params interface{}) error {
	err := c.ShouldBindJSON(params)
	if err != nil {
		return errors.CustomError(errors.InvalidParams)
	}

	valid := validation.Validation{}
	check, err := valid.Valid(params)
	if err != nil {
		return err
	}
	if !check {
		MarkErrors(valid.Errors)
		return errors.CustomError(errors.InvalidParams)
	}

	return nil
}
