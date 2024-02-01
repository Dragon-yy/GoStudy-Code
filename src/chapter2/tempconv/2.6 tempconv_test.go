package tempconv

import (
	"fmt"
	"testing"
)

func TestFToC(t *testing.T) {
	// package tempconv
	fmt.Printf("AbsoluteZeroC: %g°C\n", AbsoluteZeroC) // "AbsoluteZeroC: -273.15°C"
	fmt.Printf("FreezingC: %g°C\n", FreezingC)         // "FreezingC: 0°C"
	fmt.Printf("BoilingC: %g°C\n", BoilingC)           // "BoilingC: 100°C"
	fmt.Printf("AbsoluteZeroF: %g°F\n", CToF(AbsoluteZeroC))
	/*
		go test报错
			./2.6 tempconv_test.go:10:39: undefined: AbsoluteZeroC
			./2.6 tempconv_test.go:11:35: undefined: FreezingC
			./2.6 tempconv_test.go:12:34: undefined: BoilingC
			./2.6 tempconv_test.go:13:39: undefined: CToF
			./2.6 tempconv_test.go:13:44: undefined: AbsoluteZeroC
			原因见https://blog.csdn.net/helen920318/article/details/105017118
			原因是go test会为指定的源码文件生成一个虚拟代码包——“command-line-arguments”，
			而tempconv_test.go引用了其他包中的数据并不属于代码包“command-line-arguments”，
			编译不通过，错误自然发生了。因此，我们可以在go test的时候加上引用的包。
			//执行整个测试文件
			go test -v array_test.go array.go
			//执行测试文件中的指定方法
			go test -v array_test.go array.go -test.run TestNewArray
			//若文件存在多级依赖，可以直接在包目录下执行go test，运行包下所有的测试文件
			go test
	*/
}
