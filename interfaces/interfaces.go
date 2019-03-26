package interfaces

import (
	"signaref/models"
)

// Connection interface
type Connection interface {
	GetConnection(...string) (interface{}, error)
}

// VendorInterface interfaces
type VendorInterface interface {
	RegisterVendor(vendor models.Vendor) error
}

// UserInterface interfaces
type UserInterface interface {
	RegisterUser(user models.User) error
	LoginUser(userLogin models.UserLogin) bool
}
