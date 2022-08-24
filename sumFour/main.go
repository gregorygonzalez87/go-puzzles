package main

import (
	"fmt"
	"sort"
)

func sumFour(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
	min := sum(arr[0:4])
	max := sum(arr[1:5])
	fmt.Println(max)
	fmt.Printf("%v %v", min, max)
}

func sum(arr []int32) int64 {
	sum := int64(0)
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	return sum
}

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	sumFour(arr)
}
