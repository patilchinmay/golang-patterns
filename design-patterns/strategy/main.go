package main

import (
	driver "github.com/patilchinmay/golang-patterns/design-patterns/strategy/driver"
	mysql "github.com/patilchinmay/golang-patterns/design-patterns/strategy/driver/mysql"
	postgres "github.com/patilchinmay/golang-patterns/design-patterns/strategy/driver/postgres"
)

func main() {
	mysqlDriver := &mysql.MysqlDriver{Name: "mysql"}
	postgresDriver := &postgres.PostgresDriver{Name: "postgres"}

	databaseOperations(mysqlDriver)
	databaseOperations(postgresDriver)
}

func databaseOperations(d driver.DatabaseDriver) {
	d.Connect()
	d.Healthcheck()
	d.Disconnect()
}
