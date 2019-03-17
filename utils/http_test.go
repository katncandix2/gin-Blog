package utils_test

import (
	"simpleBlog/constant"
	"simpleBlog/model"
	"simpleBlog/utils"
	"strconv"
	"testing"
	"time"
)

var a *model.Article
var url string

func init() {
	a = &model.Article{}
	a.ArticleId = 333
	a.ArticleTitle = "测试文章标题"
	a.Context = "测试文章内容"
	a.Tags = []string{"java", "hadoop"}
	a.UserId = 2013214104
	a.UserName = "guruiqin"
	a.CreateAt = time.Now().Format(constant.TimeFormat)
	a.UpdateAt = time.Now().Format(constant.TimeFormat)
	url = constant.BaseEsUrl + constant.ArticleIndex + "/" + strconv.Itoa(int(a.ArticleId))
}

func TestPost(t *testing.T) {
	utils.Post(url, a)
}

func BenchmarkPost(b *testing.B) {

	for i := 0; i < b.N; i++ {
		utils.Post(url, a)
	}
}
