package servers

import "fmt"

type Argentinian struct {
	Interface
}

func (n *Argentinian) Connect() {
	fmt.Println("Connecting to Argentinian server")
}

func (n *Argentinian) Disconnect() {
	fmt.Println("Disconnecting from Argentinian server")
}

func (n *Argentinian) PowerOn() {
	fmt.Println("PowerOn Argentinian server")
}

func (n *Argentinian) PowerOff() {
	fmt.Println("PowerOff Argentinian server")
}
