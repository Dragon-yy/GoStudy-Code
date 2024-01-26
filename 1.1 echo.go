// 打印 os.Args[0] ，即被执行命令本身的名字
package main

import "os"

func main() {
	println(os.Args[0])
}
