package postgresql

import (
	"Grpc/services/getmem/models"
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

type Database interface {
	Getdata_db(m string) (models.GetRes, error)
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

func (db *Databaseconn) Getdata_db(m string) (models.GetRes, error) {
	q := `
			SELECT parentid, userid, postid, content, status FROM public.mem WHERE slug = $1
		`
	var mem models.GetRes
	err := db.conn.QueryRow(q, m).Scan(&mem.ParentId, &mem.UserId, &mem.PostId, &mem.Content, &mem.Status)
	if err != nil {
		return models.GetRes{}, err
	}
	return mem, nil
}
