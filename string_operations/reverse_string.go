package string_operations

import "fmt"

func Reverse() {
	input := "akshay"
	var out string
	for i := len(input) - 1; i >= 0; i-- {
		out += string(input[i])
	}
	fmt.Println(out)
}
