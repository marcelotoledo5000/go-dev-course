package scope

import "fmt"

var z string = "Hello"

// func main() {
// 	printZ()
// }

func PrintZ() { // function exported
	fmt.Println(z)
}
