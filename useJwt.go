package main

import (
	"fmt"
	"time"

	// "reflect"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

func createToken(key string, m map[string]interface{}) string {
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
func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

func main() {

	type UserInfo map[string]interface{}

	t := time.Now()
	key := "serect"

	userInfo := make(UserInfo)
	var expTime int64 = 1000 * 60 * 10
	var tokenState string

	userInfo["exp"] = strconv.FormatInt(t.UTC().UnixNano(), 10)
	userInfo["iat"] = "0"
	userInfo["name"] = "guruiqin"

	tokenString := createToken(key, userInfo)

	fmt.Println(tokenString)
	claims, ok := parseToken(tokenString, key)
	fmt.Println(claims)
	if ok {
		oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
		ct := t.UTC().UnixNano()
		c := ct - oldT
		if c > expTime {
			ok = false
			tokenState = "Token 已过期"
		} else {
			tokenState = "Token 正常"
		}
	} else {
		tokenState = "token无效"
	}

	fmt.Println(tokenState)
	fmt.Println(claims)
}
