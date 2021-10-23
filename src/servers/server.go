package servers

type Interface interface {
	Connect()
	Disconnect()
	PowerOn()
	PowerOff()
}
