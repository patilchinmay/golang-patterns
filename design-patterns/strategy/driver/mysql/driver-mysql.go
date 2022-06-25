package driver

import "log"

type MysqlDriver struct {
	Name string
}

func (d MysqlDriver) Connect() {
	log.Println("Connected to " + d.Name)
}

func (d MysqlDriver) Healthcheck() {
	log.Println(d.Name + " is healthy")
}

func (d MysqlDriver) Disconnect() {
	log.Println("Disconnected from " + d.Name)
}
