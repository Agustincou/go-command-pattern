package main

import (
	"fmt"
	"go-command-pattern/src/servers"
	"go-command-pattern/src/servers/invoker"
)

func main() {
	fmt.Println("Init a new server")

	serverInvoker := new(invoker.Server)
	serverInvoker.Initialize(new(servers.Argentinian))

	invokerError := serverInvoker.Reboot()

	fmt.Println("Finish program result:", invokerError)
}
