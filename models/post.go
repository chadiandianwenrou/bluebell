package models

import (
	"bluebell/global"
	"bluebell/models/req"
	"context"
	"encoding/json"
	"fmt"
)

func CreatePost(p *req.Post) (err error) {
	sqlStr := fmt.Sprintf(`insert into bluebell.post(post_id,title,content,author_id,community_id) values(%d,'%s','%s',%d,%d)`,
		p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	_, err = global.Mysql.Exec(context.Background(), sqlStr)
	return

}

func GetPostByID(pid int64) (post *req.Post, err error) {
	sqlStr := fmt.Sprintf(`select post_id,title,content,author_id,community_id,create_time from bluebell.post where post_id=%d`, pid)
	if err = global.Mysql.Get(context.Background(), global.SQLFormat(sqlStr), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		json.Unmarshal(jsonString, &post)
		return nil
	}); err != nil {
		return
	}
	return
}

func GetPostList(pageIndex, pageSize int64) (posts []*req.Post, err error) {
	//select post_id,title,content,author_id,community_id,create_time from bluebell.post limit (page-1)*size 从哪开始 ,size 取多少条
	sqlStr := fmt.Sprintf(`select post_id,title,content,author_id,community_id,create_time from bluebell.post ORDER BY create_time DESC`)
	sqlPage := fmt.Sprintf("%s LIMIT %d  OFFSET %d", sqlStr, pageSize, (pageIndex-1)*pageSize)
	if err := global.Mysql.Get(context.Background(), global.SQLFormat(sqlPage), func(row map[string]interface{}) error {
		jsonString, _ := json.Marshal(row)
		s := req.Post{}
		json.Unmarshal(jsonString, &s)
		posts = append(posts, &s)
		return nil
	}); err != nil {
		return nil, err
	}
	return
}
