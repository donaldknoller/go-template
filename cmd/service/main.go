package main

import (
	"go-template/internal/config"
	"go-template/internal/server"

	"github.com/rs/zerolog/log"
)

// Bound when built with go build -ldflags="-X main.Commit=$(git rev-parse HEAD)"
var Commit string

func main() {
	config.InitDefault()
	s := server.NewServer()
	log.Info().Msg(Commit)
	s.Start()
}
