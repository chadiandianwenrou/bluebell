package controller

import (
	"bluebell/global"
	"strconv"

	"bluebell/models/req"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//CreatePostHandler 创建帖子 处理函数
func CreatePostHandler(c *gin.Context) {
	//1. 获取参数和参数的校验
	var p req.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("c.ShouldBindJSON error", zap.Any("err", err))
		zap.L().Error("create post invalid param")
		global.ResponseError(c, global.CodeInvalidParam)
		return
	}

	//2.  从c 获取当前发送请求的userID
	userID, err := global.GetCurrenUserID(c)
	if err != nil {
		global.ResponseError(c, global.CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	//3.创建帖子
	if err := service.CreatePost(&p); err != nil {
		zap.L().Error(" service.CreatePost failed", zap.Error(err))
		//任何服务端的错误 就返回serverBusy，不把错误暴露给外部，自己记录个日志就好.
		global.ResponseError(c, global.CodeServerBusy)
		return
	}

	//4.返回响应
	global.ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情 处理函数
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（从URL中获取帖子ID）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post data with invalid param", zap.Error(err))
		global.ResponseError(c, global.CodeInvalidParam)
	}
	//2.根据ID取出帖子数据
	data, err := service.GetPostByID(pid)
	if err != nil {
		zap.L().Error("service.GetPostByID(pid) failed", zap.Error(err))
		global.ResponseError(c, global.CodeServerBusy)
	}
	//3.返回响应
	global.ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	//1.获取分页参数
	pageIndexStr := c.Query("pageIndex") //查多少页
	pageSizeStr := c.Query("pageSize")   //一页几条数据

	pageIndex, err := strconv.ParseInt(pageIndexStr, 10, 64)
	if err != nil {
		pageIndex = 1
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}

	//2.获取数据
	data, err := service.GetPostList(pageIndex, pageSize)
	if err != nil {
		zap.L().Error("service.GetPostList() failed", zap.Error(err))
		global.ResponseError(c, global.CodeServerBusy)
		return
	}
	//3.返回
	global.ResponseSuccess(c, data)
}
