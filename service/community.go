package service

import (
	"bluebell/models"
	"bluebell/models/rsp"
)

func GetCommunityList() ([]*rsp.Community, error) {
	//查询数据库，查找到所有community并返回
	return models.GetCommunityList()

}

func GetCommunityDetail(communityId int64) (communityDetail *rsp.CommunityDetail, err error) {
	return models.GetCommunityDetailByID(communityId)
}
