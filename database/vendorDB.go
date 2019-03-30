package database

import (
	"errors"
	"signaref/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//VendorDB is to create farmer db entity
type VendorDB struct {
	Session interface{} //This should be type asserted based on the database usage
	DBName  string
}

const (

	// ErrVendorExists is an error constant
	ErrVendorExists = "Vendor already exists or already registered with the given mobile number"
	//ErrWrongInputType is an error constant
	ErrWrongInputType = "Wrong input type"
)

//IsVendorExists is to check vendor exists or not
func (r *VendorDB) IsVendorExists(userid string) bool {
	count, err := r.Session.(*mgo.Session).DB(r.DBName).C("vendor").Find(bson.M{"_userid": userid}).Count()
	if err != nil {
		if err.Error() == "not found" {
			return false
		} else {
			return false
		}
	}
	if count > 0 {
		return true
	}
	return false
}

// RegisterVendor is to register a farmer to the system
func (r *VendorDB) RegisterVendor(vendor models.Vendor) error {
	if !r.IsVendorExists(vendor.Name) {

		if err := r.Session.(*mgo.Session).DB(r.DBName).C("vendor").Insert(vendor); err != nil {
			return err
		}
	} else {
		return errors.New(ErrVendorExists)
	}
	return nil
}

// FillVendorDetails is to fill vendor details in the system
func (r *VendorDB) FillVendorDetails(userid string, vendor *models.Vendor) error {
	if !r.IsVendorExists(vendor.Name) {
		if err := r.Session.(*mgo.Session).DB(r.DBName).C("vendor").Update(bson.M{"_userid": bson.ObjectIdHex(userid)}, &vendor); err != nil {
			return err
		}
	} else {
		return errors.New(ErrVendorExists)
	}
	return nil

	return nil
}

/*
// UpdateRawFood is to update raw food
func (r *RawFoodDB) UpdateRawFood(selector, data interface{}) error {
	if _, ok := selector.(map[string]interface{}); !ok {
		return errors.New(ErrWrongInputType)
	}
	if _, ok := data.(map[string]interface{}); !ok {
		return errors.New(ErrWrongInputType)
	}

	if err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Update(bson.M(selector.(map[string]interface{})), bson.M(data.(map[string]interface{}))); err != nil {
		return err
	}
	return nil
}


//GetRawFood is to fetch rawfood
func (r *RawFoodDB) GetRawFood(selector interface{}) (*models.RawFood, error) {
	if _, ok := selector.(map[string]interface{}); !ok {
		return nil, errors.New(ErrWrongInputType)
	}
	//var result map[string]interface{}
	var result *models.RawFood

	err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Find(bson.M(selector.(map[string]interface{}))).One(&result)
	if err != nil {
		return result, err
	}
	fmt.Println(result)
	return result, nil
}

//GetRawFoods is to fetch rawfood
func (r *RawFoodDB) GetRawFoods(skip, limit int32, selector interface{}) ([]models.RawFood, error) {
	if _, ok := selector.(map[string]interface{}); !ok {
		return nil, errors.New(ErrWrongInputType)
	}
	var result []models.RawFood
	err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Find(bson.M(selector.(map[string]interface{}))).Skip(int(skip)).Limit(int(limit)).Sort("-_id").All(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

// DeleteRawFood is to delete raw food
func (r *RawFoodDB) DeleteRawFood(selector interface{}) error {
	if _, ok := selector.(map[string]interface{}); !ok {
		return errors.New(ErrWrongInputType)
	}
	if err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Remove(selector); err != nil {
		return err
	}
	return nil
}

// AddNutrition add nutrition to the existing food
func (r *RawFoodDB) AddNutrition(selector interface{}, nutrition models.Nutrition) error {
	//query := bson.M{"code": q.(string)}
	update, err := ToMap(nutrition, "json")
	if err != nil {
		return err
	}
	update1 := bson.M{"$push": bson.M{"nutritions": update}}
	if err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Update(selector, update1); err != nil {
		return err
	}
	return nil
}

// DeleteNutrition delete nutrition to the existing food
func (r *RawFoodDB) DeleteNutrition(rootSelector, childSelector interface{}) error {
	update := bson.M{"$pull": bson.M{"nutritions": childSelector}}
	if err := r.Session.(*mgo.Session).DB(r.DBName).C("rawfood").Update(rootSelector, update); err != nil {
		return err
	}
	return nil
}*/
