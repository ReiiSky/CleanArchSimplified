package domain

import (
	"fmt"

	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"github.com/Satssuki/CleanArchSimplified/internal/usecase"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	exist, err := l.user.IsExist()
	var userResponse = "internal-error->server does not serve properly"

	if err == nil {
		if !exist {
			var oID *primitive.ObjectID

			oID, err = l.user.Insert()
			if err == nil {
				log.
					Info().
					Msgf("user registered with username: %v", l.user.Username)
				userResponse = fmt.Sprintf("created->user has been created with %v", oID)
			} else {
				log.
					Error().
					Msgf("insert user document error with message: %v", err)
			}
		} else {
			userResponse = "conflict->user already exist"
		}
	}
	return userResponse, err
}
