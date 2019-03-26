package database

import (
	"errors"
	"signaref/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserDB is to create farmer db entity
type UserDB struct {
	Session interface{} //This should be type asserted based on the database usage
	DBName  string
}

const (

	// ErrUserExists is an error constant
	ErrUserExists = "User already exists or already registered with the given mobile number or email address"
	//ErrWrongInputType is an error constant
	//ErrWrongInputType = "Wrong input type"
)

//IsUserExists is to check vendor exists or not
func (u *UserDB) IsUserExists(mobile string) bool {
	count, err := u.Session.(*mgo.Session).DB(u.DBName).C("user").Find(bson.M{"mobile": mobile}).Count()
	if err != nil {
		if err.Error() == "not found" {
			return false
		}
	}
	if count > 0 {
		return true
	}
	return false
}

//IsUserExistsForLogin is to check vendor exists or not
func (u *UserDB) IsUserExistsForLogin(userLogin models.UserLogin) bool {
	count, err := u.Session.(*mgo.Session).DB(u.DBName).C("user").Find(bson.M{"mobile": userLogin.Mobile, "email": userLogin.Email, "passcode": userLogin.Passcode, "user_type": userLogin.UserType}).Count()
	if err != nil {
		if err.Error() == "not found" {
			return false
		}
	}
	if count > 0 {
		return true
	}
	return false
}

//GetUser is to fetch user based on email and mobile as inputs
func (u *UserDB) GetUser(email, mobile string) (*models.User, error) {
	user := &models.User{}
	if err := u.Session.(*mgo.Session).DB(u.DBName).C("user").Find(bson.M{"mobile": mobile, "email": email}).One(user); err != nil {
		return nil, err
	}
	return user, nil
}

// RegisterUser is to register a farmer to the system
func (u *UserDB) RegisterUser(user models.User) error {
	if !u.IsUserExists(user.Mobile) {
		if err := u.Session.(*mgo.Session).DB(u.DBName).C("user").Insert(user); err != nil {
			return err
		}
		usr, err := u.GetUser(user.Email, user.Mobile)
		if err != nil {
			//Todo the earlier transaction revoke
			return err
		}
		if user.UserType == "vendor" {
			vendor := models.Vendor{}
			vendor.UserID = usr.ID
			if err := u.Session.(*mgo.Session).DB(u.DBName).C("vendor").Insert(vendor); err != nil {
				return err
			}
		}
		if usr.UserType == "agent" {
			//Todo Agent Logic here
		}
		if usr.UserType == "client" {
			// Todo Client Logic here
		}
		if usr.UserType == "user" {
			//Todo user(site user ) Logic here
		}
	} else {
		return errors.New(ErrUserExists)
	}
	return nil
}

//LoginUser is used to logs a user into the system
func (u *UserDB) LoginUser(userLogin models.UserLogin) bool {
	return u.IsUserExistsForLogin(userLogin)
}
