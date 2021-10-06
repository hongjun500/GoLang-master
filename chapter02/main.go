package main

import (
	"fmt"
	"math"
)

func main() {
	// 循环只有for
	// 不需要小括号，但大括号是必须
	// 分号;隔开的三个表达式
	// 初始化语句，条件表达式，后置语句
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	fmt.Println("sum值", sum)

	// 初始化语句和后置语句不是必须 例
	i := 100 // 本质上这里的i赋值100就是初始化语句
	for i > 0 {
		i-- // 而这里的i--则是后置语句，作用其实就是改变初始语句赋的值
		// 使其条件表达式的值为false结束循环，否则将陷入死循环
	}
	fmt.Println("i值", i)

	sum = 0
	for i < 100 {
		sum = sum + i
		i++
	}
	fmt.Println("更改后i值", sum)

	a := 0
	for a < 10 {
		// if语句不需要小括号，必需大括号，类似for
		if a != 0 && a%2 == 0 {
			fmt.Println("偶数", a)
		}
		a++
	}

	value := pow(2, 3, 100)
	fmt.Println(value)

	/*fmt.Println(Sqrt(2))
	fmt.Println(SqrtFloat64(2))*/

	// 这两个结果不一样
	fmt.Println(1 / 2)
	fmt.Println(1 / 2.0)
}

/**
// if语句的表达式
*/
func pow(x, y, z float64) float64 {
	// math.Pow(x, y) 求x的y次方
	if value := math.Pow(x, y); value < z {
		return value
	}
	// return value, 报错，if语句中的变量声明只能在if语句之内使用
	return z
}

/**
if语句中的变量声明在连着if后面的else中使用
*/
func powElse(x, y, z float64) float64 {
	// math.Pow(x, y) 求x的y次方
	if value := math.Pow(x, y); value < z {
		return value
	} else if value == 0 {
		return value
	} else if value == 100 {
		return value
	} else {
		value = 9999
	}
	// return value, 报错，if语句中的变量声明只能在if语句之内使用
	return z
}

func Sqrt(x int) int {
	z := 1
	for x < 10 {
		// 这里的相除结果因为是int取值时会取商,而不是同float64类型时取的精确的结果值
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/**
这里的结果会和上面不一样本质上是因为入参类型不同
*/
func SqrtFloat64(x float64) float64 {
	z := 1.0
	for x < 10 {
		// 这里的相除因为是float会取精确的结果值，而不是同int类型时取的商
		z -= (z*z - x) / (2 * z)
	}
	return z
}
