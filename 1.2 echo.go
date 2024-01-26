// 打印每个参数的索引和值，每个一行。
package main

import "os"

func main() {
	for i := range os.Args[1:] {
		println(i, os.Args[i+1])
	}
}
