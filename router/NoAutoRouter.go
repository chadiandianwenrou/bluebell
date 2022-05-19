package router

import (
	"github.com/gin-gonic/gin"
)

func NoAuthRouter(r *gin.Engine) *gin.Engine {

	r1 := r.Group("/api/v2/lion/sms/voice")
	r1.Use()
	{
		//r1.POST("/callback", oncall.CallBack)
	}

	return r
}
