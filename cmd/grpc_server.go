package main

import (
	"net"

	"signaref/database"
	"signaref/handler"
	pb "signaref/proto"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	mgo "gopkg.in/mgo.v2"
)

const (
	port = ":50051"
)

var (
	dbConnectionStr string
	dbName          string

	err error
)

func init() {
	viper.SetConfigName("app")        // no need to include file extension
	viper.AddConfigPath("../config/") // set the path of your config file

	err := viper.ReadInConfig()

	if err != nil {
		glog.Fatal(err)
	} else {
		dbConnectionStr = viper.GetString("connections.DBConnection")
		dbName = viper.GetString("connections.DBName")
	}
}

func main() {

	session, err := database.GetConnection(dbConnectionStr, dbName)
	defer glog.Flush()
	if err != nil {
		if session != nil {
			session.(*mgo.Session).Close()
		}
		glog.Fatal("mongodb database is not connected")
	}
	defer session.(*mgo.Session).Close()

	// create hvendor andler instance
	f := new(handler.Vendor)
	f.IVendor = &database.VendorDB{Session: session, DBName: dbName}

	// create user handler instance
	u := new(handler.User)
	u.IUser = &database.UserDB{Session: session, DBName: dbName}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterVendorServer(server, f)
	pb.RegisterUserServer(server, u)

	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}

}
