package router

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	//跨域配置
	//r.Use(cors.Cors())
	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/singup", controller.SingUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(jwt.JWTAuthMiddleware()) // 启用JWT认证中间件
	{
		// 查询所有社区
		v1.GET("/community", controller.CommunityHandler)
		//查询社区详情
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		//创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		//获取帖子详情
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		//获取帖子列表
		v1.GET("/posts", controller.GetPostListHandler)
	}
	//全局404配置
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})

	//r   = FsRouter(r,authMiddleware)
	//r	= MailRouter(r,authMiddleware)
	//r	= SmsTextRouter(r,authMiddleware)
	//r	= SmsVoiceRouter(r,authMiddleware)

	return r
}
