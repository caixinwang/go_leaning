package main

import (
	"fmt"
	"os"
)

/**
注意：
	go中没有分号
	go中的main函数不支持任何返回值---int main
	通过os.Exit来返回状态
*/
func main() {
	fmt.Println("hello world")
	os.Exit(-1)
}
