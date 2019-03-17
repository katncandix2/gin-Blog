package service

import (
	"context"
	"log"
	"reflect"
	"simpleBlog/constant"
	"simpleBlog/model"
	"simpleBlog/utils"
	"strconv"

	"github.com/olivere/elastic"
)

var es = utils.EsClient()

//存储编辑文章
func AddBlog(article *model.Article) *elastic.IndexResponse {

	if es == nil {
		log.Println("es nil")
		return nil
	}

	putRes, err := es.Index().
		Index(constant.EsIndexBlog).
		Type(constant.EsTypeArticle).
		Id(strconv.Itoa(int(utils.GetUUID()))).
		BodyJson(article).
		Do(context.Background())

	if err != nil {
		log.Println()
		return nil
	}

	log.Println("index->", putRes.Index, "type->", putRes.Type, "id->", putRes.Id)
	return putRes
}

//获取博客列表
func BlogList(pageNum int) []model.Article {

	if es == nil {
		log.Println("BlogList es nil")
		return nil
	}

	searchResult, err := es.Search().
		Index(constant.EsIndexBlog). // search in index "twitter"
		// Query(termQuery).        // specify the query
		Sort("article_id", true).    // sort by "user" field, ascending
		From(pageNum * 10).Size(10). // take documents 0-9
		Pretty(true).                // pretty print request and response JSON
		Do(context.Background())     // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	blogList := make([]model.Article, 10)
	var article model.Article
	for index, item := range searchResult.Each(reflect.TypeOf(article)) {
		blogList[index] = item.(model.Article)
	}

	return blogList
}

// 通过文章标题与标签搜索文章
func SearchBlogByTitleAndTags() {

	if es == nil {
		return
	}

}
