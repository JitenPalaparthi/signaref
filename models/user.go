package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email       string        `json:"email" bson:"email"`
	Mobile      string        `json:"mobile" bson:"mobile"`
	Passcode    string        `json:"passcode" bson:"passcode"`
	UserType    string        `json:"user_type" bson:"user_type"`
	Status      string        `json:"status" bson:"status"`
	LastUpdated time.Time     `json:"last_updated" bson:"last_updated"`
}

// ValidateUser validates the User type
func ValidateUser(u User) error {
	if u.Mobile == "" {
		return fmt.Errorf("Mobile " + ErrMandatoryField)
	}
	if u.Email == "" {
		return fmt.Errorf("Email " + ErrMandatoryField)
	}

	if u.Passcode == "" {
		return fmt.Errorf("Passcode " + ErrMandatoryField)
	}

	if u.UserType == "" {
		return fmt.Errorf("UserType " + ErrMandatoryField)
	}
	return nil
}
