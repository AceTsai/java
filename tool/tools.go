package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PageResult 分页结果返回
type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}

// SuccessPage 请求成功返回分页结果
func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
	page := &PageResult{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{200, message, page})
}

var SigningKey = []byte("global.Config.Jwt.SigningKey")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToke 生成Token
func GenerateToke(username string) (string, error) {
	claims := Claims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60,
		Issuer:    username,
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}
