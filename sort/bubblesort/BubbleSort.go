package bubblesort

func BubbleSort(value []int) {
	flag := true
	for i := 0; i < len(value)-1; i++ {
		for j := 0; j < len(value)-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
