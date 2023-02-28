package user

import (
	"errors"
	"time"

	"github.com/edwbaeza/coverage-api/src/shared"
	sharedDomain "github.com/edwbaeza/coverage-api/src/shared/domain"
	domain "github.com/edwbaeza/coverage-api/src/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type Creator struct {
	repository domain.UserRepository
}

func NewUserCreator(repository domain.UserRepository) *Creator {
	return &Creator{
		repository: repository,
	}
}

func (creator *Creator) Create(email string, firstName string, lastName string, password string, passwordConfirmation string, userType domain.UserType) (domain.User, error) {

	if !creator.validateEmail(email) {
		return domain.User{}, shared.NewError(errors.New("Email has taken"), "user.email")
	}

	if !creator.validatePasswords(password, passwordConfirmation) {
		return domain.User{}, shared.NewError(errors.New("Passwords do not match"), "user.passwords")
	}

	// TODO: Handle encryption password error
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return creator.repository.Create(
		domain.User{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			Password:  string(hash),
			Role:      userType,
			BaseModel: sharedDomain.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	)
}

func (creator *Creator) validatePasswords(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation
}

func (creator *Creator) validateEmail(email string) bool {
	// Improve validation so as not to retrieve the user from the database
	_, err := creator.repository.FindByEmail(email)

	return err != nil
}
