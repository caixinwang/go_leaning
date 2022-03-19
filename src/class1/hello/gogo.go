/**
注意到这里我们的目录名--包的名字是可以不是main，但是里面的package一定要是main，不然不能运行。
也就是说有main函数，上面就一定要是main包
有关于应用程序入口：
	必须是main包---package main
	必须有main方法---func main()
	文件名不一定是main.go---比如我们这边用gogo.go
*/
package main //这里一定要是main
import "fmt" // 引入代码依赖

// 功能实现
func main() {
	fmt.Println("Hello World")
}
