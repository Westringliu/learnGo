package main

import "fmt"

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0]
		//make只能用于slice，map，chan，如mySlice2 := make([]int, 5, 10)
		//创建一个初始元素长度为5的数组切片，元素初始值为0，并预留10个元素的存储空间：
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, splitdata)
		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = quickSort(low), quickSort(high)
		//将切片打散一个个append
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(quickSort(arr))
}
