package models

import (
	"bluebell/global"
	"bluebell/models/req"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chadiandianwenrou/go/types"
)

const secret = "baidu.com"

//CheckUserExist 检查指定用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := fmt.Sprintf(`select count(user_id) from bluebell.user where username='%s'`, username)
	var count int32

	if err := global.Mysql.Get(context.Background(), sqlStr, func(row map[string]interface{}) error {
		for _, v := range row {
			c, _ := types.ToInt(v)
			count = int32(c)
		}
		return nil
	}); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

//InserUser 向数据库插入一条新的用户记录
func InserUser(user *req.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := fmt.Sprintf(`insert into bluebell.user(user_id,username,password) values(%d,'%s','%s')`, user.UserID, user.Username, user.Password)
	_, err = global.Mysql.Exec(context.Background(), sqlStr)
	if err != nil {
		return
	}
	return

}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *req.User) (err error) {
	oPassword := user.Password
	sqlStr := fmt.Sprintf(`select user_id,username,password from bluebell.user where username='%s'`, user.Username)
	//获取到userID username，password
	if err = global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		json.Unmarshal(jsonString, user)
		return nil
	}); err != nil {
		//查询数据库错误
		return err
	}

	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}

	return
}

func GetUsetByID(uid int64) (user *req.User, err error) {
	sqlStr := fmt.Sprintf(`select user_id,username from bluebell.user where user_id=%d`, uid)
	if err = global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		json.Unmarshal(jsonString, &user)
		return nil
	}); err != nil {
		return
	}
	return
}
