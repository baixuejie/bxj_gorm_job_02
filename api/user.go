package api

import (
	"bxj_gorm_job_02/model"
	"bxj_gorm_job_02/model/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (api *UserApi) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		response.FailWithMessage("用户名或密码不能为空", c)
		return
	}
	loginResponse, err := userService.Login(username, password)
	if err != nil {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	response.OkWithDetailed(loginResponse, "登录成功", c)
}

func (api *UserApi) Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	_, err = userService.Register(&user)
	if err != nil {
		response.FailWithMessage("注册失败", c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

func (api *UserApi) GetUserInfo(c *gin.Context) {
	user, err := userService.GetUserInfo(c.Query("name"))
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	response.OkWithDetailed(user, "获取用户信息成功", c)
}
