package controller

import (
	"bluebell/global"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//-------跟社区相关----------
// CommunityHandler查询所有社区分类
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id,community_name）,以切片的形式返回
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList() failed", zap.Error(err))
		global.ResponseError(c, global.CodeServerBusy)
		return
	}
	global.ResponseSuccess(c, data)
}

// CommunityDetailHandler查询社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseInt(communityIdStr, 10, 64)
	if err != nil {
		global.ResponseError(c, global.CodeInvalidParam)
		return
	}
	data, err := service.GetCommunityDetail(communityId)
	if err != nil {
		zap.L().Error("service.GetCommunityDetail() failed", zap.Error(err))
		global.ResponseError(c, global.CodeServerBusy)
		return
	}
	global.ResponseSuccess(c, data)

}
