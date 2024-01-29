// 打印 os.Args[0] ，即被执行命令本身的名字
package main

import "os"

func echo1() {
	println(os.Args[0])
}

func echo2() {
	for i := range os.Args[1:] {
		println(i, os.Args[i+1])
	}
}

func main() {
	//echo1()
	echo2()
}
