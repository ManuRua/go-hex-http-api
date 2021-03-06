package bootstrap

import "go-hex-http-api/internal/infra/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}