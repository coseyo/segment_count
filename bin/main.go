package main

import (
	"fmt"
	"os"

	"github.com/coseyo/segment_count"
)

func main() {
	argNum := len(os.Args)
	if argNum != 9 {
		fmt.Printf("\n输入格式: \n\nsegment_count -srcDB 源数据库 -srctable 源数据表名 -srcfield 源数据表字段名 -disttable 目标数据表名\n\n")
		return
	}

	srcDB := os.Args[2]
	srcTable := os.Args[4]
	srcField := os.Args[6]
	distTable := os.Args[8]

	err := segment_count.Exec(srcDB, srcTable, srcField, distTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("完工!")

}
