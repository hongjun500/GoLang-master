package main

import "fmt"

func main() {
	// 数组声明形式：var 数组名 [数组长度]数组类型，数组类型也可以是数组，因此可以构成数组类型的数组，即多维数组
	// 声明一个二维数组
	var arrArray [3][2]int
	fmt.Printf("arrArray零值%v \n", arrArray[0][0])

	// 使用三个...来定义数组长度时代表其长度根据数据初始化时的个数来决定
	// var arrLen = [...]string{"hello"}
	// fmt.Printf(len(arrLen))  这一行会报错，因为数组长度固定无法求取，除非是切片

	// 声明一个名为array类型为int,长度为10的数组
	var array [10]int
	for i := 0; i < len(array); i++ {
		// 从0开始给数组赋值
		array[i] = i
	}
	fmt.Printf("数组array%v \n", array)
	fmt.Printf("数组array最后一个值=%v \n", array[len(array)-1])

	// 由于数组是定长的(数组的长度是其类型的一部分，因此数组不能改变大小)
	// 因此在Go中更多使用的是切片，(切片是长度可以变化的数组)切片和数组又息息相关
	var arrSlice []string
	if arrSlice == nil {
		fmt.Println("切片arrSlice是空的")
	}
	// arrSlice[0] = "hello"  这一行编译通过，但是运行报错，切片只是声明了但没有初始化
	// 本质上切片的底层实现是一个指向数组的指针，再没有存入一个具体的数组之前，就是nil
	// arrSlice = []string{""}  直接是以这种方式简单方式就可以使用
	// 或者通过内置函数make初始化 make([]T,len,cap)
	// 长度必须定义，容量可以省略并且容量必须大于等于长度
	arrSlice = make([]string, 10, 10)
	arrSlice[0] = "hello"
	fmt.Printf("切片arrSlice%v \n", arrSlice)
	fmt.Printf("切片arrSlice长度%v \n", len(arrSlice))
	fmt.Printf("切片arrSlice容量%v \n", cap(arrSlice))

	// 切片可以截取
	// 切片通过两个下标来截取，一个上界一个下界两者之间以冒号分割
	// splice[low:high],半开区间，包含low,排除high
}
