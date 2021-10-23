package actions

import "fmt"

type PowerOff struct {
	Commander
	iServer
}

func (p *PowerOff) ExecuteRoutine() error {
	if p.Server != nil {
		p.Server.Connect()
		p.Server.PowerOff()
		p.Server.Disconnect()

		return nil
	} else {
		return initializeServerError
	}
}

func (p *PowerOff) TearDown() {
	fmt.Println("Cleaning all temporary files...")
}
