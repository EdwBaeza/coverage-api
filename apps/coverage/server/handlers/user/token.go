package user

import (
	"errors"
	"net/http"

	"github.com/edwbaeza/coverage-api/src/shared"
	"github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/edwbaeza/coverage-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `key:"username" json:"username" binding:"required"`
	Password string `key:"password" json:"password" binding:"required"`
}

func TokenHandler(repository domain.UserRepository) func(context *gin.Context) {
	return func(context *gin.Context) {
		var input LoginInput
		err := context.ShouldBindJSON(&input)

		// TODO: Add Finder
		user, err := repository.FindByEmail(input.Username)

		if err != nil || !ComparePasswords(user.Password, input.Password) {
			context.Error(shared.NewError(errors.New("Invalid Password or Username"), "username")).
				SetType(gin.ErrorTypeBind).
				SetMeta(http.StatusBadRequest)
			return
		}

		token, err := utils.GenerateToken(user.ID)
		// TODO: Add Updater
		_, err = repository.Update(user.ID, domain.User{Token: token})

		if err != nil {
			context.Error(err).SetType(gin.ErrorTypeBind).SetMeta(http.StatusBadRequest)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
