package driver

import "log"

type PostgresDriver struct {
	Name string
}

func (d PostgresDriver) Connect() {
	log.Println("Connected to " + d.Name)
}

func (d PostgresDriver) Healthcheck() {
	log.Println(d.Name + " is healthy")
}

func (d PostgresDriver) Disconnect() {
	log.Println("Disconnected from " + d.Name)
}
