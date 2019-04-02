package main

import (
	"flag"
	"net"
	"net/http"

	"signaref/database"
	"signaref/handler"
	gw "signaref/proto"
	pb "signaref/proto"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

var (
	echoEndpoint = flag.String("endpoint", "localhost:50051", "endpoint of YourService")
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

	// create vendor handler instance
	f := new(handler.Vendor)
	f.IVendor = &database.VendorDB{Session: session, DBName: dbName}

	// create user handler instance
	u := new(handler.User)
	u.IUser = &database.UserDB{Session: session, DBName: dbName}

	pr := new(handler.Product)
	pr.IProduct = &database.ProductDB{Session: session, DBName: dbName}
	glog.Info(pr)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterProductServer(server, pr)
	pb.RegisterVendorServer(server, f)
	pb.RegisterUserServer(server, u)

	go run()

	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}

	// This runs the rest client
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

// This is purely rest related stuff
func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterVendorHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	err = gw.RegisterUserHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	err = gw.RegisterProductHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":50052", mux)
}
