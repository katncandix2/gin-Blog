package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"simpleBlog/constant"
	"strings"
)

func Get(url string) []byte {

	if len(url) < 0 {
		log.Fatalln("url len <0")
		return nil
	}

	res, err := http.Get("http://localhost:9200/website/blog/1")

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("err2")
		return nil
	}

	return body
}

func Post(url string, val interface{}) error {

	if len(url) < 0 {
		log.Fatalln("url is empty")
		return errors.New("url length")
	}

	jRes, err := json.Marshal(val)

	if err != nil {
		log.Fatalln("json err:", err.Error())
		return err
	}

	body := strings.NewReader(string(jRes))

	res, err := http.Post(url, constant.HttpContentTypeJson, body)

	if err != nil {
		log.Fatalln("post err:", err.Error())
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalln("err here:", res.StatusCode)
		return errors.New("some err happen")
	}

	return nil
}
