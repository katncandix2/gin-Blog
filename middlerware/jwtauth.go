package middlerware

import (
	"fmt"
	"net/http"
	"simpleBlog/constant"
	"simpleBlog/model"
	"simpleBlog/service"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(
				http.StatusOK, gin.H{
					"code": -1,
					"msg":  "未登录请去登录",
				},
			)
			return
		}

		res, err := parseToken(token, constant.AuthSerect)

		if err != constant.LoginSuccess {
			c.JSON(
				http.StatusOK, gin.H{
					"code": -1,
					"msg":  "内部错误请稍后再试",
				},
			)
			return
		}

		//todo 判断token 过期时间
		u := &model.User{}

		//todo 此处若token 不合法则会pannic
		u.UserName = res.(jwt.MapClaims)["user_name"].(string)
		u.PassWord = res.(jwt.MapClaims)["pass_word"].(string)
		if constant.LoginSuccess != service.CheckLogin(u) {
			c.JSON(
				http.StatusOK, gin.H{
					"code": -1,
					"msg":  "密码错误",
				},
			)
			return
		}

		c.Set("user", u)
		c.Next()

	}
}

//解析token
func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(key), nil
		})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func CreateToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	// fmt.Println(_map)
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}
