package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ErrMndatoryField is an error message
var (
	ErrMandatoryField = "field is mandatory"
)

// Farmer model
type Vendor struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code        int32         `json:"code" bson:"code"`
	Name        string        `json:"name" bson:"name"`
	ContactNo   string        `json:"contactno" bson:"contactno"`
	Email       string        `json:"email" bson:"email"`
	Passcode    string        `json:"passcode" bson:"passcode"`
	Status      string        `json:"status" bson:"status"`
	LastUpdated time.Time     `json:"last_updated" bson:"last_updated"`
}

// ValidateVendor validates the Vendor type
func ValidateVendor(v Vendor) error {
	if v.Code == 0 {
		return fmt.Errorf("Code " + ErrMandatoryField)
	}

	if v.Name == "" {
		return fmt.Errorf("Name " + ErrMandatoryField)
	}

	if v.ContactNo == "" {
		return fmt.Errorf("Contact Number " + ErrMandatoryField)
	}

	if v.Email == "" {
		return fmt.Errorf("Email " + ErrMandatoryField)
	}
	return nil
}
