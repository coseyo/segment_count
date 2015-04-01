package main

import (
	"fmt"
	"os"

	"github.com/coseyo/segment_count"
)

func main() {
	argNum := len(os.Args)
	if argNum != 7 {
		fmt.Printf("\n输入格式: \n\nsegment_count -srctable 源数据表名 -srcfield 源数据表字段名 -disttable 目标数据表名\n\n")
		return
	}

	srcTable := os.Args[2]
	srcField := os.Args[4]
	distTable := os.Args[6]

	err := segment_count.Exec(srcTable, srcField, distTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("完工!")

}
