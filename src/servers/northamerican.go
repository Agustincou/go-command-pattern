package servers

import "fmt"

type NorthAmerican struct {
	Interface
}

func (n *NorthAmerican) Connect() {
	fmt.Println("Connecting to North American server")
}

func (n *NorthAmerican) Disconnect() {
	fmt.Println("Disconnecting from North American server")
}

func (n *NorthAmerican) PowerOn() {
	fmt.Println("PowerOn North American server")
}

func (n *NorthAmerican) PowerOff() {
	fmt.Println("PowerOff North American server")
}
