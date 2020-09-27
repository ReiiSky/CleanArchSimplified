package usecase

import "github.com/Satssuki/CleanArchSimplified/internal/interfaces"

// User struct to storing user info
type User struct {
	Username  string `json:"username"`
	LikeCount uint   `json:"likeCount"`
}

// Serialize ..
func (u *User) Serialize(s *interfaces.Serializer) error {
	return s.Transform(u)
}
