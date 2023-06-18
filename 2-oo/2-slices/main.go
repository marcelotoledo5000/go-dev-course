package main

import "fmt"

func main() {
	// make a slice of int, with 2 positions and capacity of 5
	slice := make([]int, 2, 5)
	slice = append(slice, 1, 2, 3, 5, 7, 11, 13, 17, 29, 31, 37)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
		fmt.Println(slice)
	}

	// pointer
	doubleSlice := slice
	slice[0] = 47
	slice = append(slice, 100, 101, 102)
	fmt.Println(slice)
	fmt.Println(doubleSlice)

	sliceString := []string{
		"Hello", "World", "Much", "Better", "Last",
	}
	fmt.Println(sliceString)
	fmt.Println(sliceString[0])   // first position
	fmt.Println(sliceString[:2])  // first 2 positions
	fmt.Println(sliceString[2:4]) // from index 2 to position 4
	fmt.Println(sliceString[2:])  // from index 2 to the end
}
