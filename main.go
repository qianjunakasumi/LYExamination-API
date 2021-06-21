package main

import (
	"net"

	"git.qianjunakasumi.ren/lyexamination/api/lyexamination"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	lyexamination.RegisterLYExaminationServer(server, lyexamination.LYExamination{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
