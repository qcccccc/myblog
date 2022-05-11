package service

import (
	"errors"
	"gocode/dao"
	"gocode/models"
	"gocode/utils"
)

func Register(userName, passwd string) error {
	user := dao.UserIsexit(userName, passwd)
	if user == nil {
		return errors.New("账户已存在")
	}
	return nil
}

func Login(userName, passwd string) (*models.LoginRes, error) {
	user := dao.GetUser(userName, passwd)
	//fmt.Println(passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}

	uid := user.Uid
	var userInfo models.UserInfo
	userInfo.Uid = uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未完成")
	}
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
