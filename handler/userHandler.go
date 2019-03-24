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

// User type
type User struct {
	IUser interfaces.UserInterface
}

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
}

// RegisterVendor is to register a vendor and store it in the database
func (u *User) RegisterUser(ctx context.Context, in *pb.UserDetails) (*pb.GeneralResponse, error) {
	user := models.User{}
	user.Email = in.Email
	user.Mobile = in.Mobile
	user.Passcode = in.Passcode
	user.UserType = in.UserType
	user.Status = "active"
	user.LastUpdated = time.Now().UTC()
	out := &pb.GeneralResponse{}
	if err := models.ValidateUser(user); err != nil {
		glog.Error(err)
		return nil, err
	}

	if err := u.IUser.RegisterUser(user); err != nil {
		glog.Info(err)
		return nil, err
	}
	out.Code = 201
	out.Message = "Success"
	return out, nil
}
