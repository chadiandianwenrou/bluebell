package controller

import (
	"bluebell/global"
	"bluebell/models/req"
	lib_validator "bluebell/pkg/validator"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//SingUpHandler 处理注册请求的函数
func SingUpHandler(c *gin.Context) {
	//1. 获取参数和参数校验
	var data req.SingUp
	if err := c.ShouldBindJSON(&data); err != nil {
		zap.L().Error("SingUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.ResponseError(c, global.CodeInvalidParam)
			return
		}

		global.ResponseErrorWithMsg(c, global.CodeInvalidParam, lib_validator.RemoveTopStruct(errs.Translate(lib_validator.Trans)))
		return

	}

	//2.业务处理
	if err := service.SingUp(&data); err != nil {
		global.ResponseErrorWithMsg(c, global.CodeInvalidParam, "注册失败")
		return
	}
	//3.返回响应
	global.ResponseSuccess(c, nil)
}

//LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	//1.获取请求参数和参数校验
	var data req.Login
	if err := c.ShouldBindJSON(&data); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.ResponseError(c, global.CodeInvalidParam)
			return
		}
		global.ResponseErrorWithMsg(c, global.CodeInvalidParam, lib_validator.RemoveTopStruct(errs.Translate(lib_validator.Trans)))
		return

	}
	//2. 业务逻辑处理
	token, err := service.Login(&data)
	if err != nil {
		zap.L().Error("service.Login failed", zap.String("username", data.Username), zap.Error(err))
		global.ResponseError(c, global.CodeInvalidPassword)
		return
	}
	//3.返回响应
	global.ResponseSuccess(c, token)
}
