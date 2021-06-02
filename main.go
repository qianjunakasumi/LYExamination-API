package main

import (
	"net"

	"git.qianjunakasumi.ren/lyexamination/api/lyexamination"
	pb "git.qianjunakasumi.ren/lyexamination/api/protobuf"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":59392")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterLYExaminationServer(server, lyexamination.LYExamination{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
