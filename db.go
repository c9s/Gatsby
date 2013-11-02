package gatsby

import "database/sql"

const (
	DriverPg = iota
	DriverMysql
	DriverSqlite
)

var conn *sql.DB
var driverType int

func SetupConnection(c *sql.DB, t int) {
	conn = c
	driverType = t
}

func CloseConnection() {
	if conn != nil {
		conn.Close()
	}
}

func GetConnection() *sql.DB {
	return conn
}

type ConnectionHandle struct {
	conn       *sql.DB
	driverType int
}

func (self *ConnectionHandle) Load(val PtrRecord, pkId int64) *Result {
	return Load(self.conn, val, pkId)
}

func (self *ConnectionHandle) LoadByCols(val PtrRecord, cols WhereMap) *Result {
	return LoadByCols(self.conn, val, cols)
}

func (self *ConnectionHandle) Create(val PtrRecord, driver int) *Result {
	return Create(self.conn, val, driver)
}

func (self *ConnectionHandle) Update(val PtrRecord, driver int) *Result {
	return Update(self.conn, val, driver)
}

func (self *ConnectionHandle) Delete(val PtrRecord) *Result {
	return Delete(self.conn, val)
}
