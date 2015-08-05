package main

import (
	"net"

	"github.com/lawnchairstudios/RSVP/proto"
	"github.com/lawnchairstudios/RSVP/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// create a list of invitees

func main() {

	// bootstrap the application
	var application = &system.Application{}

	application.Init()
	defer application.DeInit()

	// listen on port 8000
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// setup the options
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterRSVPServer(grpcServer, application.ServerServices.Invitees)
	grpcServer.Serve(lis)

}
