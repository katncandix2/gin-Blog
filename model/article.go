package model

// 文章内容
type Article struct {
	ArticleId    int32    `form:"article_id" json:"article_id"`
	ArticleTitle string   `form:"article_name" json:"article_title"`
	Context      string   `form:"context" json:"context"`
	Tags         []string `form:"tags" json:"tags"`
	UserId       int32    `form:"user_id" json:"user_id"`
	UserName     string   `form:"user_name" json:"user_name"`
	CreateAt     string   `json:"create_at"`
	UpdateAt     string   `json:"update_at"`
}
