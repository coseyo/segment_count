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

func Test_dataInit(t *testing.T) {
	dataObj := data{}
	err := dataObj.init()
	expect(t, err, nil)
	err = dataObj.db.Ping()
	expect(t, err, nil)
}

func Test_dataCount(t *testing.T) {
	dataObj := data{}
	dataObj.init()
	total, err := dataObj.count()
	expect(t, err, nil)
	fmt.Printf("source table count : %d\n", total)
}

func Test_dataRead(t *testing.T) {
	dataObj := data{}
	dataObj.init()
	titleArray, err := dataObj.read(0, 1)
	expect(t, err, nil)
	fmt.Println(titleArray)
}

func Test_dataWrite(t *testing.T) {
	dataObj := data{}
	dataObj.init()
	err := dataObj.write([]string{"datatest", "datatest2", "datatest"})
	expect(t, err, nil)
}
