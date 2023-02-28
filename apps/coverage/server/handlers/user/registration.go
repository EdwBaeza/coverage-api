package user

import (
	"net/http"

	userApplication "github.com/edwbaeza/coverage-api/src/user/application/user"
	"github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/gin-gonic/gin"
)

// Improve authentication and authorization workflow maybe using JWT and OAuth2
type RegisterInput struct {
	Email                string `key:"email" json:"email" binding:"required"`
	Password             string `key:"password" json:"password" binding:"required"`
	PasswordConfirmation string `key:"password_confirmation" json:"password_confirmation" binding:"required"`
	FirstName            string `key:"first_name" json:"first_name" binding:"required"`
	LastName             string `key:"last_name" json:"last_name" binding:"required"`
}

func RegistrationHandler(repository domain.UserRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		var input RegisterInput
		err := context.ShouldBindJSON(&input)

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypeBind).SetMeta(http.StatusBadRequest)
			return
		}

		creator := userApplication.NewUserCreator(repository)
		user, err := creator.Create(input.Email, input.FirstName, input.LastName, input.Password, input.PasswordConfirmation, 2)

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusCreated, gin.H{"data": user})
	}
}
