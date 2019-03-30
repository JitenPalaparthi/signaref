package handler

import (
	"context"
	"signaref/interfaces"
	"signaref/models"
	pb "signaref/proto"
	"time"

	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc/metadata"

	"flag"

	"github.com/golang/glog"
)

// Vendor type
type Vendor struct {
	IVendor interfaces.VendorInterface
}

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
}

// RegisterVendor is to register a vendor and store it in the database
func (s *Vendor) RegisterVendor(ctx context.Context, in *pb.VendorDetails) (*pb.GeneralResponse, error) {
	v := models.Vendor{}
	v.Name = in.Name
	v.Status = "active"
	v.LastUpdated = time.Now().UTC()
	out := &pb.GeneralResponse{}
	if err := models.ValidateVendor(v); err != nil {
		glog.Error(err)
		return nil, err
	}

	if err := s.IVendor.RegisterVendor(v); err != nil {
		glog.Info(err)
		return nil, err
	}
	out.Code = 201
	out.Message = "Success"
	return out, nil
}

// FillVendorDetails is to register a vendor and store it in the database
func (s *Vendor) FillVendorDetails(ctx context.Context, in *pb.VendorDetails) (*pb.GeneralResponse, error) {
	v := &models.Vendor{}
	var userid string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userid = md["userid"][0]
	}
	v.UserID = bson.ObjectIdHex(userid)
	v.Name = in.Name
	v.Website = in.Website
	if in.Address != nil {
		v.SetAddress(in.Address.Line1, in.Address.Line2, in.Address.Line3, in.Address.Street, in.Address.City, in.Address.State, in.Address.Country, in.Address.Postalcode, in.Address.Pobox)
	}
	if in.Location != nil {
		v.SetLocation(in.Location.Type, in.Location.Longitude, in.Location.Latitude)
	}
	if in.ContactPerson != nil {
		v.SetContactPerson(in.ContactPerson.Name, in.ContactPerson.Email, in.ContactPerson.ContactNo)
	}
	if in.BusinessDetails != nil {
		v.SetBusinessDetails(in.BusinessDetails.Type, in.BusinessDetails.LicenceType, in.BusinessDetails.LicenceCode)
	}
	v.Status = "filled"
	v.LastUpdated = time.Now().UTC()
	out := &pb.GeneralResponse{}
	/*if err := models.ValidateVendor(v); err != nil {
		glog.Error(err)
		return nil, err
	}*/

	if err := s.IVendor.FillVendorDetails(userid, v); err != nil {
		glog.Info(err)
		return nil, err
	}
	out.Code = 201
	out.Message = "Success"
	return out, nil
}
