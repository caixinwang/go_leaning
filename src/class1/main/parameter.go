package main

import (
	"fmt"
	"os"
)

/**
go语言中main函数不支持传入任何参数例如args，命令行参数args可以用os.Args来获得
*/
func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello !!", os.Args)
	}
}
