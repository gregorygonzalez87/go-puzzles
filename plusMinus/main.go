package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	n := len(arr)
	sort.Slice(arr, func(i, j int) bool { return arr[j] < arr[i] })
	firstPos := sort.Search(len(arr), func(i int) bool { return arr[i] <= 0 })
	fmt.Println(arr)
	fmt.Println(firstPos)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	lastNeg := sort.Search(len(arr), func(i int) bool { return arr[i] >= 0 })
	fmt.Println(arr)
	fmt.Println(lastNeg)

	if firstPos == n {
		fmt.Println(1.00)
		fmt.Println(0.00)
		fmt.Println(0.00)
		return
	}

	if lastNeg == n {
		fmt.Println(0.00)
		fmt.Println(1.00)
		fmt.Println(0.00)
		return
	}

	pos := 0.00
	if arr[n-1-firstPos] == 0 {
		pos = float64(firstPos)
	} else if firstPos < n-1 {
		pos = float64(firstPos)
	}
	fmt.Println(pos)

	neg := 0.00
	if arr[lastNeg] == 0 {
		neg = float64(lastNeg)
	} else if lastNeg < n-1 {
		neg = float64(lastNeg)
	}
	fmt.Println(neg)

	total := float64(n)
	zer := total - pos - neg
	totalPos := pos / total
	totalNeg := neg / total
	totalZer := zer / total

	fmt.Println(totalPos)
	fmt.Println(totalNeg)
	fmt.Println(totalZer)
}

func main() {
	test := []int32{55, 48, 48, 45, 91, 97, 45, 1, 39, 54, 36, 6, 19, 35, 66, 36, 72, 93, 38, 21, 65, 70, 36, 63, 39, 76, 82, 26, 67, 29, 24, 82, 62, 53, 1, 50, 47, 65, 67, 19, 66, 90, 77}
	plusMinus(test)
}
