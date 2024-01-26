package main

import (
	"bufio"
	"os"
	"strings"
)

// 从标准输入中获取数据，然后输出重复行的个数和文本
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	// 注意：忽略 input.Err() 中可能的错误
	for line, n := range counts {
		if n > 1 {
			println(line, n)
		}
	}
}

// map本身是引用类型，所以不需要返回值
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
}

type FileCounts struct {
	fileName string
	counts   map[string]int
}

// 从文件或用户输入中获取数据，然后输出重复行的个数和文本
func dup2() {
	fileCounts := []FileCounts{{"", map[string]int{}}}
	files := os.Args[1:]
	if len(files) == 0 {
		// 如果没有文件参数，从标准输入中读取数据
		countLines(os.Stdin, fileCounts[0].counts)
	} else {
		// 如果有文件参数，从文件中读取数据
		i := 0
		for _, arg := range files {
			//fmt.Println(arg)
			fileCounts[i].fileName = arg
			f, err := os.Open(arg)
			if err != nil {
				// 如果文件打开失败，输出错误信息并继续处理下一个文件
				println("dup2:", err)
				continue
			}
			countLines(f, fileCounts[i].counts)
			defer f.Close()
			i++
			fileCounts = append(fileCounts, FileCounts{"", map[string]int{}})
		}
	}
	// 注意：忽略 input.Err() 中可能的错误
	for _, f := range fileCounts {
		for line, n := range f.counts {
			if n > 1 {
				println(f.fileName, line, n)
			}
		}
	}

}

// 第二种实现dup2功能的方式
func dup3() {
	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := os.ReadFile(file)
		//fmt.Println(string(data))

		if err != nil {
			println("dup3:", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[string(line)]++
		}
		for line, n := range counts {
			if n > 1 {
				println(line, n)
			}
		}
	}
}

func main() {
	//dup1()
	dup2()
	//dup3()
}
