package segment_count

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	db        *sql.DB
	srcDB     string
	SrcTable  string
	SrcField  string
	DistTable string
}

func (this *data) init() error {
	if this.srcDB == "" {
		this.srcDB = SRC_DB
	}
	db, err := sql.Open("mysql", MYSQL_USER+":"+MYSQL_PWD+"@tcp("+MYSQL_HOST+":"+MYSQL_PORT+")/"+this.srcDB+"?charset=utf8")
	if err != nil {
		fmt.Println("database initialize error : ", err.Error())
		return err
	}
	this.db = db

	if this.SrcTable == "" {
		this.SrcTable = SRC_TABLE
	}
	if this.SrcField == "" {
		this.SrcField = SRC_FIELD
	}
	if this.DistTable == "" {
		this.DistTable = DIST_TABLE
	}

	if err := this.createDistTable(); err != nil {
		return err
	}

	return nil
}

func (this *data) createDistTable() error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (`word` varchar(100) NOT NULL, `total` int(11) unsigned NOT NULL, PRIMARY KEY (`word`)) ENGINE=InnoDB DEFAULT CHARSET=utf8", this.DistTable)
	_, err := this.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (this *data) read(offset int, limit int) ([]string, error) {
	sql := fmt.Sprintf("SELECT `%s` FROM `%s` LIMIT %d, %d", this.SrcField, this.SrcTable, offset, limit)
	stmt, err := this.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var title string
	titleArray := []string{}
	for stmt.Next() {
		if err := stmt.Scan(&title); err != nil {
			return nil, err
		}
		titleArray = append(titleArray, title)
	}
	return titleArray, nil
}

func (this *data) write(word []string) error {
	var values string
	for _, value := range word {
		values += fmt.Sprintf("(\"%s\", 1),", value)
	}
	values = strings.TrimRight(values, ",")
	sql := fmt.Sprintf("INSERT INTO `%s` (`word`, `total`) VALUES %s ON DUPLICATE KEY UPDATE `total`=`total`+1", this.DistTable, values)
	_, err := this.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (this *data) count() (count int, err error) {
	sql := fmt.Sprintf("SELECT COUNT(1) FROM `%s`", this.SrcTable)
	err = this.db.QueryRow(sql).Scan(&count)
	if err != nil {
		return
	}
	return
}
