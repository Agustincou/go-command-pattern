package main

import (
	"fmt"
	"go-command-pattern/src/servers"
	"go-command-pattern/src/servers/actions"
	"go-command-pattern/src/servers/invoker"
)

func main() {
	fmt.Println("Init a new server")

	desiredCommander := new(actions.PowerOn)
	desiredCommander.InitializeServer(new(servers.Argentinian))

	commanderInvoker := new(invoker.Invoker)
	commanderInvoker.Initialize(desiredCommander)

	invokerError := commanderInvoker.Run()

	fmt.Println("Finish program result:", invokerError)
}
