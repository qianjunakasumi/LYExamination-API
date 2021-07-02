package main

import (
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		With().Caller().
		Logger()

	log.Info().Msg("build a friendly Longyan examination App for students | welcome to join the QQ group: 422954727")

	lis, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Panic().Err(err)
	}
	if lis == nil {
		log.Panic().Msg("监听 :443 失败")
	}

	server := grpc.NewServer()
	RegisterLYExaminationServer(server, &LYExamination{})
	log.Info().Msg("监听 :443 正常")
	if err := server.Serve(lis); err != nil {
		log.Panic().Err(err)
	}
}
