package main

import (
	"crypto/tls"
	"net"

	"git.qianjunakasumi.ren/lyexamination/api/lyexamination"
	pb "git.qianjunakasumi.ren/lyexamination/api/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	lis, err := net.Listen("tcp", ":59392")
	if err != nil {
		panic(err)
	}

	CDNCert, err := tls.LoadX509KeyPair("cert/cdn.pem", "cert/cdn.key")
	if err != nil {
		panic(err)
	}

	SrvCert, err := tls.LoadX509KeyPair("cert/srv.pem", "cert/srv.key")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&CDNCert)),
		grpc.Creds(credentials.NewServerTLSFromCert(&SrvCert)),
	)
	pb.RegisterLYExaminationServer(server, lyexamination.LYExamination{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
