package invoker

import (
	"errors"
	"go-command-pattern/src/servers"
	"go-command-pattern/src/servers/actions"
)

var initializeError = errors.New("server not defined, initialize the invoker first")

type Server struct {
	server servers.Interface
}

func (i *Server) Initialize(server servers.Interface) {
	i.server = server
}

func (i *Server) Up() error {
	var errorToReturn error

	if i.server != nil {
		desiredCommander := new(actions.PowerOn)
		//ToDo: adding a command manager to avoid repetition code
		desiredCommander.InitializeServer(i.server)
		errorToReturn = desiredCommander.ExecuteRoutine()
		desiredCommander.TearDown()
	} else {
		errorToReturn = initializeError
	}

	return errorToReturn
}

func (i *Server) Down() error {
	var errorToReturn error

	if i.server != nil {
		desiredCommander := new(actions.PowerOff)
		//ToDo: adding a command manager to avoid repetition code
		desiredCommander.InitializeServer(i.server)
		errorToReturn = desiredCommander.ExecuteRoutine()
		desiredCommander.TearDown()
	} else {
		errorToReturn = initializeError
	}

	return errorToReturn
}

func (i *Server) Reboot() error {
	var errorToReturn error

	if i.server != nil {
		desiredCommander := new(actions.Reboot)
		//ToDo: adding a command manager to avoid repetition code
		desiredCommander.InitializeServer(i.server)
		errorToReturn = desiredCommander.ExecuteRoutine()
		desiredCommander.TearDown()
	} else {
		errorToReturn = initializeError
	}

	return errorToReturn
}
