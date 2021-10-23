package actions

import "fmt"

type Reboot struct {
	Commander
	iServer
}

func (p *Reboot) ExecuteRoutine() error {
	if p.Server != nil {
		p.Server.Connect()
		p.Server.PowerOff()
		p.Server.PowerOn()
		p.Server.Disconnect()

		return nil
	} else {
		return initializeServerError
	}
}

func (p *Reboot) TearDown() {
	fmt.Println("Cleaning all temporary files...")
}
