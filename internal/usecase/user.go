package usecase

import (
	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct to storing user info
type User struct {
	mgm.DefaultModel
	Username  string `json:"username" bson:"username"`
	LikeCount uint   `json:"likeCount" bson:"likeCount"`
}

// Serialize as you can see user serialize self by serializer method
func (u *User) Serialize(s *interfaces.Serializer) error {
	return s.Transform(u)
}

// IsExist check username if already exits
func (u *User) IsExist() (bool, error) {
	// future safety
	// i must check the u.Username is nil or not
	// because it will caused panic
	count, err := mgm.Coll(u).CountDocuments(mgm.Ctx(), bson.M{
		"username": u.Username,
	})
	return count > 0, err
}

// Insert user document to the database
func (u *User) Insert() (*primitive.ObjectID, error) {
	result, err := mgm.Coll(u).InsertOne(mgm.Ctx(), u)
	var objectInsertedID *primitive.ObjectID = nil

	if err == nil {
		v, x := result.InsertedID.(primitive.ObjectID)
		if x {
			objectInsertedID = &v
		}
	}
	return objectInsertedID, err
}
