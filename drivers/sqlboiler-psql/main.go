package main

import (
	"github.com/ravenops/sqlboiler/drivers"
	"github.com/ravenops/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
