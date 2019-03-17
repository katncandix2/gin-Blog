package service

import (
	"log"
	"simpleBlog/constant"
	"simpleBlog/model"
	"simpleBlog/utils"
)

func CheckLogin(u *model.User) bool {

	r := utils.GetRedis()

	if r == nil {
		log.Println("redis err")
		return constant.LoginFail
	}

	curKey := u.UserName
	cacheKey := r.Get(curKey).Val()

	if u.PassWord == cacheKey {
		return constant.LoginSuccess
	}

	return constant.LoginFail
}
