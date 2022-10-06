package postgresql

import (
	"Grpc/services/createmem/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"fmt"
)

type Database interface {
	Postdata_db(m models.CreateReq) models.CreateRes
}

type database struct {
	conn *sql.DB
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "psblyu136250"
	db_name  = "postgres"
)

func NewDatabase() (*database, error) {
	connStr := "host=" + host + " port=" + port + " dbname=" + db_name + " user=" + user + " password=" + password + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("\n DATABASE CONNECTION SUCESSFULL=1")
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()

	return &database{conn: db}, err
}

func (db *database) Postdata_db(m models.CreateReq) models.CreateRes {
	q := "INSERT INTO public.mem (parentId,userId,postId,content) VALUES ((SELECT MAX(postId) FROM public.mem)+1,'" + m.ParentId + "','" + m.UserId + "','" + m.Content + "')"
	row, err := db.conn.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return models.CreateRes{}
}
