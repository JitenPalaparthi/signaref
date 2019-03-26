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

// Vendor model
type Vendor struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID      bson.ObjectId `json:"user_id" bson:"_userid,omitempty"`
	Code        int32         `json:"code" bson:"code"`
	Name        string        `json:"name" bson:"name"`
	WebSite     string        `json:"website" bson:"website"`
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
	return nil
}
