package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

type Database interface {
	Deletedata_db(id string) error
}

type Databaseconn struct {
	conn *sql.DB
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "psblyu136250"
	db_name  = "postgres"
)

func NewDatabase() (*Databaseconn, error) {
	connStr := "host=" + host + " port=" + port + " dbname=" + db_name + " user=" + user + " password=" + password + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("\n DATABASE CONNECTION SUCESSFULL=1")
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()

	return &Databaseconn{conn: db}, err
}

func (d Databaseconn) Deletedata_db(id string) error {
	//TODO implement me
	panic("implement me")
}
