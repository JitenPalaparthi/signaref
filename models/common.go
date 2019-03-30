package models

//Address model
type Address struct {
	Line1      string `json:"line1" bson:"line1"`
	Line2      string `json:"line2" bson:"line2"`
	Line3      string `json:"line3" bson:"line3"`
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Country    string `json:"country" bson:"country"`
	PostalCode string `json:"postalCode" bson:"postalCode"`
	POBox      string `json:"pobox" bson:"pobox"`
}

//Location model
type Location struct {
	Type      string `json:"type" bson:"type"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}

//Person model
type Person struct {
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	ContactNo string `json:"contactNo" bson:"contactNo"`
}

//BusinessDetails model
type BusinessDetails struct {
	Type        string `json:"type" bson:"type"`
	LicenceType string `json:"licenceType" bson:"licenceType"`
	LicenceCode string `json:"licenceCode" bson:"licenceCode"`
}
