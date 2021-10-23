package actions

import (
	"errors"
	"go-command-pattern/src/servers"
)

var initializeServerError = errors.New("server not defined, initialize the commander first")

type Commander interface {
	ExecuteRoutine() error
	TearDown()
}

type iServer struct {
	Server servers.Interface
}

func (p *iServer) InitializeServer(server servers.Interface) {
	p.Server = server
}
