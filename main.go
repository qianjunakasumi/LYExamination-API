package main

import (
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	RegisterLYExaminationServer(server, LYExamination{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
