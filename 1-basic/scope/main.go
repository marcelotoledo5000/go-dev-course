package scope

import (
	"fmt"
	"scope"
)

var Y int = 20

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(Y * x)
	PrintNum()
	fmt.Println(scope.Z)
	scope.PrintZ()
}

func PrintNum() {
	fmt.Println(Y)
}
