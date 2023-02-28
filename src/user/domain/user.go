package domain

import (
	shared "github.com/edwbaeza/coverage-api/src/shared/domain"
)

type UserType int

const (
	Admin    UserType = 1
	Customer          = 2
)

type User struct {
	Password string `json:"-" bson:"password,omitempty"`
	// TODO: Permit to use multiple tokens
	Token            string          `json:"token,omitempty" bson:"token,omitempty"`
	FirstName        string          `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName         string          `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email            string          `json:"email,omitempty" bson:"email,omitempty"`
	Role             UserType        `json:"user_role,omitempty" bson:"user_role,omitempty"` // TODO: Improve user's roles and permissions management with more complex structur
	shared.BaseModel `bson:"inline"` // TODO: Improve token management
}

func (user User) IsSuperAdmin() bool {
	return user.Role == Admin
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Find(id string) (User, error)
	Create(user User) (User, error)
	Update(id string, user User) (User, error)
}
