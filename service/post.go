package service

import (
	"bluebell/models"
	"bluebell/models/req"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

// 创建帖子
func CreatePost(p *req.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GetID()
	//2.保存到数据库
	return models.CreatePost(p)
	//3.返回

}

// 根据帖子ID 查询帖子详情数据
func GetPostByID(pid int64) (data *req.ApiPostDetail, err error) {
	//查询并组合接口想要的数据

	post, err := models.GetPostByID(pid)
	if err != nil {
		zap.L().Error("models.GetPostByID(pid) failed", zap.Error(err))
		return
	}

	//根据作者ID 查询作者ID
	user, err := models.GetUsetByID(post.AuthorID)
	if err != nil {
		zap.L().Error("models.GetUsetByID(post.AuthorID) failed", zap.Error(err))
		return
	}

	//根据社区ID查询社区详细信息
	communityDetail, err := models.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error(" models.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
		return
	}

	//返回数据拼接
	data = &req.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}

	return

}

//获取帖子列表
func GetPostList(pageIndex, pageSize int64) (data []*req.ApiPostDetail, err error) {
	posts, err := models.GetPostList(pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		//根据作者ID 查询作者ID
		user, err := models.GetUsetByID(post.AuthorID)
		if err != nil {
			zap.L().Error("models.GetUsetByID(post.AuthorID) failed", zap.Error(err))
			continue
		}

		//根据社区ID查询社区详细信息
		communityDetail, err := models.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error(" models.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
			continue
		}
		postDetail := &req.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postDetail)

	}
	return
}
