package main

import "fmt"

func mapSlice() {
	fmt.Println("无锁map扩展部分--------------------------------------")
	// 无锁map的扩展部门
	// 元素类型为map的切片(长度为2)
	// make(t Type, size ...IntegerType) Type
	// 创建了一个值为切片，并且切片的类型是一个map,而map的键值对都是string
	var mapSlice = make([]map[string]string, 2)
	// fmt.Println("初始化创建的mapSlice", mapSlice)
	for key, value := range mapSlice {
		fmt.Printf("切片mapSlice未赋值时的key=%v   ,value = %v \n", key, value)
	}

	// 对切片mapSlice进行初始化赋值
	mapSlice[0] = make(map[string]string, 100)
	mapSlice[0]["name"] = "hongjun500"
	mapSlice[0]["age"] = "23"
	// mapSlice[1]["gener"] = "2"  这一行编译通过，但是运行报错因为这里代表取切片下标为的元素map,且map的键为gender的值
	for key, value := range mapSlice {
		fmt.Printf("切片mapSlice通过赋值后的key=%v   ,value = %v \n", key, value)

		// 如果此轮循环体中切片的元素不为空则对元素即map再次进行一个遍历
		if value != nil {

			fmt.Printf("切片位置mapSlice[%v],此轮循环体中切片不为空 \n", key)
			/*for key, value := range value {
				fmt.Printf("切片mapSlice[%v] = %v   , \n", key, value)
			}*/
		} else {
			fmt.Printf("切片位置mapSlice[%v],此轮循环体中切片是空的 \n", key)
		}

		if len(value) != 0 {

			for key, value := range value {
				fmt.Printf("切片mapSlice[%v] = %v  \n", key, value)
			}
		}
	}

	// 值为切片的map示例
	// make(t Type, size ...IntegerType) Type 定义
	// 这一行的code示例代笔创建了一个map,并且map的键是int ,map的值为string类型的切片，长度指定为2(make创建必须指定)
	var sliceMap = make(map[int][]string, 2)

	value1, key1 := sliceMap[1]
	value2, key2 := sliceMap[2]
	if key1 && key2 {
		fmt.Printf("sliceMap[\"%v\"]=%v, sliceMap[\"%v\"]=%v \n", key1, value1, key2, value2)
	} else {
		value1 = []string{"见到你"}
		value2 = []string{"很高兴"}
		msg := append(value1, value2...)
		sliceMap[0] = msg
	}
	fmt.Println(sliceMap[0])
}
