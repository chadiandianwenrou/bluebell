package req

import (
	"bluebell/models/rsp"
	"time"
)

type Post struct {
	ID          int64     `json:"post_id" db:"post_id"`
	Title       string    `json:"title" db:"title"  binding:"required"`
	Content     string    `json:"content" db:"content"  binding:"required"`
	AuthorID    int64     `json:"author_id" db:"author_id" `
	CommunityID int64     `json:"community_id" db:"community_id"  binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName           string                    `json:"author_name"`
	*Post                                          //嵌入帖子详情信息
	*rsp.CommunityDetail `json:"community_detail"` // 嵌入社区信息
}
