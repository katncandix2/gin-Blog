package model

type User struct {
	UserId   int    `form:"user_id" json:"user_id"`
	UserName string `form:"user_name" json:"user_name"`
	PassWord string `form:"pass_word" json:"pass_word"`
}
