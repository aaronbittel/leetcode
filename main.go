package main

import (
	"fmt"
)

func main() {
	num := 240
	k := 2
	res := divisorSubstrings(num, k)
	fmt.Println(num, k)
	fmt.Println(res, "sol", res == 2)
	fmt.Println("---------")

	num = 430043
	k = 2
	res = divisorSubstrings(num, k)
	fmt.Println(num, k)
	fmt.Println(res, "sol", res == 2)
	fmt.Println("---------")
}
