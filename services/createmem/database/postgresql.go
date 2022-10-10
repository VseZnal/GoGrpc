package postgresql

import (
	"Grpc/services/createmem/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database interface {
	Postdata_db(m *models.CreateReq) error
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

func (db *Databaseconn) Postdata_db(m *models.CreateReq) error {
	/*
		q := `
				INSERT INTO public.mem
				    (parentId,userId,postId,content)
				VALUES
				    ($1, $2, $3, $4)
				RETURNING *;
				`

		if err := db.conn.QueryRow(q, m.ParentId, m.UserId, m.PostId, m.Content).Scan(&m.ParentId, &m.UserId, &m.PostId, m.Content); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				pgErr = err.(*pgconn.PgError)
				newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
				return newErr
			}
			return err
		}

		return nil
	*/
	q := "INSERT INTO public.mem (parentId,userId,postId,content) VALUES ('" + m.ParentId + "','" + m.UserId + "','" + m.PostId + "','" + m.Content + "')"
	row, err := db.conn.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return err
}
