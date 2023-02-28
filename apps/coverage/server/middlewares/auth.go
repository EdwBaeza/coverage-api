package middlewares

import (
	"net/http"

	domainUser "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/edwbaeza/coverage-api/src/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(repository domainUser.UserRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utils.TokenValid(context)

		if err != nil {
			errors := []ErrorMsg{{"Authorization Token", err.Error()}}
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": errors})
			return
		}

		userId, err := utils.ExtractTokenID(context)

		if err != nil {
			errors := []ErrorMsg{{"Authorization Token ID", err.Error()}}
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": errors})
			return
		}

		user, err := repository.Find(userId)

		if err != nil {
			errors := []ErrorMsg{{"Authorization User", err.Error()}}
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": errors})
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
