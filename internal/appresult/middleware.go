package appresult

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"saglyk-backend/internal/config"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

func HeaderContentTypeJson() (string, string) {
	return "Content-Type", "application/json"
}
func AccessControlAllow() (string, string) {
	return "Access-Control-Allow-Origin", "*"
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func JwtTokenCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		cfg := config.GetConfig()
		claims, err := TokenClaims(token, cfg.JwtKey)
		if err != nil || fmt.Sprint(claims["id"]) == "" {
			respondWithError(c, 400, err)
			return
		}

		c.Set("id", fmt.Sprint(claims["id"]))
		c.Next()
	}
}

func TokenClaims(token, secretKey string) (jwt.MapClaims, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("err", err)
		return nil, ErrMissingParam
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !ok {
		// TODO tokenin omrini test etmeli
		return nil, ErrInternalServer
	}

	return claims, nil
}
