// 从main包开始
package main

// 单个导入依赖包
import "fmt"

import (
	"time"
	// 导入但不引用改包的函数,即忽略
	_ "math"
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
}
