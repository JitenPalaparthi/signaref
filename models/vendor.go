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
	ID              bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	UserID          bson.ObjectId   `json:"user_id" bson:"_userid,omitempty"`
	Code            int32           `json:"code" bson:"code"`
	Name            string          `json:"name" bson:"name"`
	Website         string          `json:"website" bson:"website"`
	Address         Address         `json:"address" bson:"address"`
	Location        Location        `json:"location" bson:"location"`
	ContactPerson   Person          `json:"contactPerson" bson:"contactPerson"`
	BusinessDetails BusinessDetails `json:"businessDetails" bson:"businessDetails"`
	Status          string          `json:"status" bson:"status"`
	LastUpdated     time.Time       `json:"lastUpdated" bson:"lastUpdated"`
}

// SetAddress sets the address
func (v *Vendor) SetAddress(line1, line2, line3, street, city, state, country, postalcode, pobox string) {
	v.Address.Line1 = line1
	v.Address.Line2 = line2
	v.Address.Line3 = line3
	v.Address.Street = street
	v.Address.City = city
	v.Address.State = state
	v.Address.Country = country
	v.Address.PostalCode = postalcode
	v.Address.POBox = pobox
}

// SetContactPerson sets the contactperson for the vendor
func (v *Vendor) SetContactPerson(name, email, contactno string) {
	v.ContactPerson.Name = name
	v.ContactPerson.Email = email
	v.ContactPerson.ContactNo = contactno
}

// SetLocation sets the location for the vendor
func (v *Vendor) SetLocation(_type, longitude, latitude string) {
	v.Location.Type = _type
	v.Location.Longitude = longitude
	v.Location.Latitude = latitude
}

// SetBusinessDetails sets the businessdetails for the vendor
func (v *Vendor) SetBusinessDetails(_type, licencetype, licencecode string) {
	v.BusinessDetails.Type = _type
	v.BusinessDetails.LicenceCode = licencecode
	v.BusinessDetails.LicenceType = licencetype
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
