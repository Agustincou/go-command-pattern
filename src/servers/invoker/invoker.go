package invoker

import (
	"errors"
	"go-command-pattern/src/servers/actions"
)

var initializeError = errors.New("commander not defined, initialize the invoker first")

type Invoker struct {
	commander actions.Commander
}

func (i *Invoker) Initialize(commander actions.Commander) {
	i.commander = commander
}

func (i *Invoker) Run() error {
	var errorToReturn error

	if i.commander != nil {
		errorToReturn = i.commander.ExecuteRoutine()
		i.commander.TearDown()
	} else {
		errorToReturn = initializeError
	}

	return errorToReturn
}
