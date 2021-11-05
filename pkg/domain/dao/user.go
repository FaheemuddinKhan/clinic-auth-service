package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	FirstName     string              `bson:"first_name"`
	LastName      string              `bson:"last_name"`
	Username      string              `bson:"username"`
	Email         string              `bson:"email"`
	Phone         string              `bson:"phone"`
	Password      string              `bson:"password"`
	DOB           time.Time           `bson:"dob"`
	Age           int8                `bson:"age"`
	PhoneVerified bool                `bson:"phone_verified"`
	EmailVerified bool                `bson:"email_verified"`
	Address       Address             `bson:"address,omitempty"`
	CreatedAt     time.Time           `bson:"created_at"`
	ModifiedAt    primitive.Timestamp `bson:"modified_at,omitempty"`
	IsDeleted 	  bool `bson:"is_deleted"`
	DeletedAt primitive.Timestamp `bson:"deleted_at,omitempty"`
}
