package main

import (
	"fmt"
	"net/http"
	"simpleBlog/constant"
	"simpleBlog/middlerware"
	"simpleBlog/model"
	"simpleBlog/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color

	r := gin.Default()
	r.Use(middlerware.JwtAuth())
	r.Any("index", func(c *gin.Context) {
		example := c.MustGet("user").(*model.User)
		fmt.Println(example)
		c.String(http.StatusOK, "index")
	})

	r.POST("login", func(c *gin.Context) {
		u := &model.User{}
		c.Bind(u)

		if u.UserName == "" || u.PassWord == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "未输入请重新输入",
			})
			return
		}

		if constant.LoginSuccess != service.CheckLogin(u) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "login error",
			})
			return
		}

		//生成token
		key := "guruiqin"
		myMap := make(map[string]interface{})
		myMap["user_name"] = u.UserName
		myMap["exp"] = time.Now().Format(constant.TimeFormat)
		myMap["pass_word"] = u.PassWord
		token := middlerware.CreateToken(key, myMap)
		fmt.Println(token)
		c.Writer.Header().Set("Authorization", token)
		// c.Header("Authorization", token)
		c.String(http.StatusOK, "success")
		// c.Redirect(http.StatusOK, "http://localhost:8080/index")
	})

	// r.Use(middlerware.Auth())
	blogRouter := r.Group("/blog")
	blogRouter.Use(middlerware.Auth())
	blogRouter.GET("index", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})
	//博客发布编辑页
	blogRouter.POST("/edit", func(c *gin.Context) {
		a := &model.Article{}
		c.Bind(a)
		a.CreateAt = time.Now().Format(constant.TimeFormat)
		a.UpdateAt = time.Now().Format(constant.TimeFormat)
		service.AddBlog(a)
		c.JSON(http.StatusOK, a)
	})

	//博客列表
	blogRouter.Any("/list", func(c *gin.Context) {

		pageNum, _ := strconv.Atoi(c.DefaultPostForm("page_num", "0"))
		res := make(map[string]interface{})
		res["code"] = 0
		res["mes"] = "success"
		res["data"] = service.BlogList(pageNum)
		c.JSON(http.StatusOK, res)
	})

	blogRouter.Any("/search", func(c *gin.Context) {
		c.String(http.StatusOK, "search success")
	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
