package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/edwbaeza/coverage-api/src/shared"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	// TODO: improve this translation's logic to support more languages (I18n?)
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "oneof":
		return "Select one of " + fe.Param()
	}
	return "Unknown error"
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		err := context.Errors.Last()

		if err != nil {
			log.Println("Error: ", err)
			// For binding errors
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
				}
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
				return
			}

			// For application errors
			var appError *shared.Error
			if errors.As(err, &appError) {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": []ErrorMsg{{appError.Path, appError.Error()}}})
				return
			}

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": []ErrorMsg{{"Unknown", err.Error()}}})
		}
	}
}
