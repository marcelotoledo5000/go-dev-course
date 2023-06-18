package main

import "fmt"

func main() {
	// array definition
	var arr [10]int

	fmt.Println(arr)
	fmt.Println(len(arr))

	arr[0] = 9
	arr[1] = 3
	arr[2] = 7

	fmt.Println(arr)
	fmt.Println(arr[2])

	// array definition and attribution
	strArr := [3]string{"A", "b", "C"}
	fmt.Println(strArr)
}
