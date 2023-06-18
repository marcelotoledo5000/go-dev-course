package scope

import "fmt"

var Z string = "Hello"

func PrintZ() { // function exported
	fmt.Println(Z)
}
