package postgresql

import (
	protocreatemem "Grpc/services/createmem/protocreatemem/proto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
	Postdata_db(m *protocreatemem.Memcreate) error
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (db *Databaseconn) Postdata_db(m *protocreatemem.Memcreate) error {

	q := `
				INSERT INTO public.mem
				    (slug,parentid,userid,postid,content,status)
				VALUES
				    ($1, $2, $3, $4, $5, $6)
				RETURNING *;
				`
	/*
		if err := db.conn.QueryRow(q, m.Slug, m.ParentId, m.UserId, m.PostId, m.Content, m.Status).Scan(&m.Slug, &m.ParentId, &m.UserId, &m.PostId, &m.Content, &m.Status); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				pgErr = err.(*pgconn.PgError)
				newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
				return newErr
			}
			return err
		}
	*/
	_, e := db.conn.Exec(q, m.Slug, m.ParentId, m.UserId, m.PostId, m.Content, m.Status)
	CheckError(e)

	return nil

	/*
		q := "INSERT INTO public.mem (parentId,userId,postId,content) VALUES ('" + m.ParentId + "','" + m.UserId + "','" + m.PostId + "','" + m.Content + "')"
		row, err := db.conn.Query(q)
		if err != nil {
			log.Fatal(err)
		}
		defer func(row *sql.Rows) {
			err := row.Close()
			if err != nil {

			}
		}(row)
		return err
	*/

}
