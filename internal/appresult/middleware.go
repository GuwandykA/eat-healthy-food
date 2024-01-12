package appresult

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
