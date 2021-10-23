package actions

import "fmt"

type PowerOn struct {
	Commander
	iServer
}

func (p *PowerOn) ExecuteRoutine() error {
	if p.Server != nil {
		p.Server.Connect()
		p.Server.PowerOn()
		p.Server.Disconnect()

		return nil
	} else {
		return initializeServerError
	}
}

func (p *PowerOn) TearDown() {
	fmt.Println("Cleaning all temporary files...")
}
