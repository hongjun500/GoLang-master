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

// var 语句可以出现在包或函数级别
// 声明两个string变量
var java, goLang string = "java", "go"

// 也可以使用此声明变量列表
var (
	php bool = true
	// 初始化了值可以省略类型，(自动推导类型)
	python = 1024

	price float32 = 9.9
	// 浮点类型默认float64
	price1 float64 = 9.9
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
	// 换个行
	fmt.Print("\n")
	fmt.Printf("打印事先声明的变量%v,%v,%v,%v", java, goLang, python, php)
	// 换个行
	fmt.Print("\n")

	// 数值类型显示转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("\nx,y,f,z的值%v,%v,%v,%v", x, y, f, z)

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
返回值的命名应当具有一定含义，可以作为文档使用
返回值命名时，返回最后可以直接return
*/
func moreReturn(x, y, z int) (a, b int, c string) {
	a = x + y + z
	b = x * y * z
	// return a, b, "string"
	// 效果同上return,如果方法比较长这种方式可读性比较差
	return
}

// 对比上类方法
func moreReturnV1(x, y, z int) (int, int, string) {
	// := 语句可以省略var声明
	// 但是 := 只能在函数内声明，方法外只能以关键字开头
	a := x + y + z
	b := x * y * z
	return a, b, "string"
}
