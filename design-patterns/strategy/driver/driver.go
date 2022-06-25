package driver

type DatabaseDriver interface {
	// getname() string
	Connect()
	Healthcheck()
	Disconnect()
}
