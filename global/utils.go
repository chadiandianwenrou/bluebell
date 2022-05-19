package global

import (
	"errors"
	"github.com/gin-gonic/gin"
)

//获取当前登录用户的ID
func GetCurrenUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get("userID")
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	return userID, nil
}
