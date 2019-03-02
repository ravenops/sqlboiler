package main

import (
	"github.com/ravenops/sqlboiler/drivers"
	"github.com/ravenops/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
