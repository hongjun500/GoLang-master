// 从main包开始
package main

// 单个导入依赖包
import "fmt"

// 分组导入多个(推荐)
import (
	"math"
	"time"
)

/**
主程序
*/
func main() {
	// 打印不换行
	fmt.Print("hello, GoLang!")

	// 打印换行
	fmt.Println(time.Now())

	fmt.Println(time.Hour)

	fmt.Println(math.Pi)
	// 报错，小写字母开头未导出(即为无法引用)
	// 不同包内小写开头无法导入
	// fmt.Println(math.pi)

	fmt.Println(add(1, 1))

	// 有参无返回值
	logNoReturn("hello, go")

	// 无参且无返回值
	logNoReturnAndParam()
}

/**
两个int类型入参，返回一个int类型参数
*/
func add(x, y int) int {
	return x + y
}

/**
一个string类型入参并打印，无法返回值
*/
func logNoReturn(str string) {
	fmt.Println(str)
}

/**
无参数无返回值的打印
*/
func logNoReturnAndParam() {
	fmt.Println("无参数无返回值的打印")
}
