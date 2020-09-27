package domain

import (
	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"github.com/Satssuki/CleanArchSimplified/internal/usecase"
)

// UserLogic wrapper for logic user
type UserLogic struct {
	user usecase.User
}

// CreateUserLogic initializer userlogic
func CreateUserLogic() *UserLogic {
	return new(UserLogic)
}

// Compose composing member model in userlogic
func (l *UserLogic) Compose(s *interfaces.Serializer) error {
	return l.user.Serialize(s)
}

// RegisterIfNotExist pure business logic for registering user
func (l *UserLogic) RegisterIfNotExist() (string, error) {
	return "conflict->user already exist", nil
}
