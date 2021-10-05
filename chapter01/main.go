// 从main包开始
package main

// 单个导入依赖包
import (
	"fmt"
)

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

	fmt.Println(sum(1, 1))

	// 有参无返回值
	logNoReturn("hello, go")

	// 无参且无返回值
	logNoReturnAndParam()

	// 方法可以有多个参数并且有多个返回值
	arr1, arr2, arr3 := moreReturn(1, 2, 3)
	/*intArrays := make([]int, 3)
	stringArrays := make([]string, 1)
	intArrays[0] = arr1
	intArrays[1] = arr2
	stringArrays[0] = arr3
	fmt.Println(intArrays, stringArrays)*/

	// 打印时可接受参数%v
	fmt.Printf("返回的三个参数分别是%v,%v,%v", arr1, arr2, arr3)

}

/**
一个string类型入参并打印，无法返回值
*/
func logNoReturn(str string) {
	fmt.Println(str)
}

/**
两个int类型入参，返回一个int类型参数
*/
func sum(x, y int) int {
	return x + y
}

/**
无参数无返回值的打印
*/
func logNoReturnAndParam() {
	fmt.Println("无参数无返回值的打印")
}

/**
接收三个参数，并且返回两个int类型一个string类型的值
注意这里的返回值a,b,c被命名了
*/
func moreReturn(x, y, z int) (a, b int, c string) {
	a = x + y + z
	b = x * y * z
	return a, b, "string"
}

// 对比上类方法
func moreReturnV1(x, y, z int) (int, int, string) {
	a := x + y + z
	b := x * y * z
	return a, b, "string"
}
