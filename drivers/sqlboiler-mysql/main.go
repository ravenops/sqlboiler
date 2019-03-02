package main

import (
	"github.com/ravenops/sqlboiler/drivers"
	"github.com/ravenops/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
