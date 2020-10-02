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

// Compose member model in userlogic
func (l *UserLogic) Compose(s *interfaces.Serializer) error {
	return l.user.Serialize(s)
}

// Register pure business logic for registering user
// first return value is response to the user whats happend in the server
// check cmd/application/handler/helpers.go
func (l *UserLogic) Register() (string, error) {
	return "conflict->user already exist", nil
}
