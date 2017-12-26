package main

import (
	"fmt"
)

// Search ...
func Search(val int, arr []int) bool {
	for _, a := range arr {
		if val == a {
			return true
		}
	}
	return false
}

func main() {
	sli := []int{1, 2, 3}
	ret := Search(2, sli)

	fmt.Println(ret)
}
