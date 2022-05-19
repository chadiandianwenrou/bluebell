package service

import (
	"bluebell/models"
	"bluebell/models/req"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SingUp(s *req.SingUp) (err error) {
	//1.判断用户是否存在
	if err = models.CheckUserExist(s.Username); err != nil {
		return err
	}
	//2.生成uid
	userID := snowflake.GetID()

	//3.保存进数据库
	user := &req.User{
		UserID:   userID,
		Username: s.Username,
		Password: s.Password,
	}
	return models.InserUser(user)
}

func Login(l *req.Login) (token string, err error) {
	//用户生成传到jwt token 里所用到
	user := &req.User{
		Username: l.Username,
		Password: l.Password,
	}
	//传递过去的user是指针， 就能拿到user.UserID
	if err = models.Login(user); err != nil {
		return "", err
	}
	//生成token， 把userID和username传过去，加到jwt token里
	return jwt.GenToken(user.UserID, user.Username)

}
