package models

import (
	"bluebell/global"
	"bluebell/models/rsp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetCommunityList() (communityList []*rsp.Community, err error) {
	sqlStr := `select community_id,community_name from bluebell.community`
	if err := global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		s := rsp.Community{}
		json.Unmarshal(jsonString, &s)
		communityList = append(communityList, &s)
		return nil
	}); err != nil {
		err = errors.New("查询失败")
		return nil, err
	}
	return
}

func GetCommunityDetailByID(communityId int64) (communityDetail *rsp.CommunityDetail, err error) {
	sqlStr := fmt.Sprintf(`select community_id,community_name,introduction,create_time
				from bluebell.community 
				where community_id=%d`,
		communityId)
	if err := global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		json.Unmarshal(jsonString, &communityDetail)
		// doc = append(doc,s)
		return nil
	}); err != nil {
		err = errors.New("查询失败")
		return communityDetail, err
	}
	return
}
