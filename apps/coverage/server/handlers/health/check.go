package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
