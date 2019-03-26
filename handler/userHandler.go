package handler

import (
	"context"
	"errors"
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

// RegisterUser is to register a user and store it in the database
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

// LoginUser is to register a vendor and store it in the database
func (u *User) LoginUser(ctx context.Context, in *pb.LoginDetails) (*pb.GeneralResponse, error) {
	userDetails := models.UserLogin{}
	userDetails.Email = in.Email
	userDetails.Mobile = in.Mobile
	userDetails.Passcode = in.Passcode
	userDetails.UserType = in.UserType

	out := &pb.GeneralResponse{}
	if err := models.ValidateUserLogin(userDetails); err != nil {
		glog.Error(err)
		return nil, err
	}

	if !u.IUser.LoginUser(userDetails) {
		glog.Info("Your account or password is incorrect")
		return nil, errors.New("Your account or password is incorrect")
	}

	out.Code = 201
	out.Message = "Success"
	return out, nil
}
