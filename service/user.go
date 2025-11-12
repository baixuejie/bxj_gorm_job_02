package service

import (
	"bxj_gorm_job_02/config"
	"bxj_gorm_job_02/model"
	utils "bxj_gorm_job_02/util"
	"errors"
)

type UserService struct {
}

func (s *UserService) GetUserInfoByUserName(username string) (*model.User, error) {
	var user model.User
	config.DB.Where("username = ?", username).First(&user)
	return &user, nil
}

func (s *UserService) Register(u *model.User) (*model.User, error) {
	u.Password = utils.BcryptHash(u.Password)
	config.DB.Create(u)
	return u, nil
}

func (s *UserService) GetUserInfo(id string) (model.User, error) {
	var u model.User
	config.DB.Where("id = ?", id).First(&u)
	return u, nil
}
func (s *UserService) UpdateUserInfo(u *model.User) (*model.User, error) {
	config.DB.Where("id = ?", u.Id).Updates(u)
	return u, nil
}
func (s *UserService) DeleteUser(u *model.User) (*model.User, error) {
	config.DB.Where("id = ?", u.Id).Delete(u)
	return u, nil
}
func (u *UserService) Login(username string, password string) (model.LoginResponse, error) {
	var user model.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return model.LoginResponse{}, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.BcryptCheck(password, user.Password) {
		return model.LoginResponse{}, errors.New("用户名或密码错误")
	}

	// 生成token
	token, expireAt, err := utils.GenerateToken(&user)
	if err != nil {
		return model.LoginResponse{}, errors.New("生成token失败")
	}

	loginResponse := model.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: expireAt,
	}
	return loginResponse, nil
}
