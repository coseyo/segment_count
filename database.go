package segment_count

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func (this *Database) Init() error {
	db, err := sql.Open("mysql", MYSQL_USER+":"+MYSQL_PWD+"@tcp("+MYSQL_HOST+":"+MYSQL_PORT+")/"+MYSQL_DB+"?charset=utf8")
	if err != nil {
		log.Println("database initialize error : ", err.Error())
		return err
	}
	this.db = db
	return nil
}
