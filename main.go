package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chutified/resource-finder/protos/commodity"
	"github.com/chutified/resource-finder/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var cfg = struct {
	Host string
	Port int
}{
	Host: "127.0.0.1",
	Port: 10501,
}

func main() {

	// define logger
	l := log.New(os.Stdout, "[COMMODITY SERVICE] ", log.LstdFlags)

	// service server
	cmdSrv := server.New(l)

	// grpc server
	grpcSrv := grpc.NewServer()

	// register
	commodity.RegisterCommodityServer(grpcSrv, cmdSrv)
	reflection.Register(grpcSrv) // support reflection

	// define listen
	lst, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		l.Fatalf("define listening: %v", err)
	}

	// start listening
	err = grpcSrv.Serve(lst)
	if err != nil {
		l.Panicf("listening: %v", err)
	}
}