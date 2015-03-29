package segment_count

import (
	"fmt"
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_Init(t *testing.T) {
	dbt := Database{}
	err := dbt.Init()
	expect(t, err, nil)
}

func Test_Read(t *testing.T) {
	dbt := Database{}
	dbt.Init()
	sql := fmt.Sprintf("select %s from %s limit 1", MYSQL_SOURCE_FIELD, MYSQL_SOURCE_TABLE)
	rows, err := dbt.db.Query(sql)
	expect(t, err, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var title string
	for rows.Next() {
		if err := rows.Scan(&title); err == nil {
			fmt.Println(title)
		}
	}
	expect(t, err, nil)
}
