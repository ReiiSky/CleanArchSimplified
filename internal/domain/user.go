package domain

import (
	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"github.com/Satssuki/CleanArchSimplified/internal/usecase"
)

// UserLogic ..
type UserLogic struct {
	user usecase.User
}

// CreateUserLogic ..
func CreateUserLogic() *UserLogic {
	return new(UserLogic)
}

// Compose ..
func (l *UserLogic) Compose(s *interfaces.Serializer) error {
	return l.user.Serialize(s)
}

// RegisterIfNotExist ..
func (l *UserLogic) RegisterIfNotExist() (string, error) {
	return "conflict->user already exist", nil
}
