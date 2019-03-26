package handler

import (
	"context"
	"signaref/interfaces"
	"signaref/models"
	pb "signaref/proto"
	"time"

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
