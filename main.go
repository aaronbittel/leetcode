package main

import (
	"fmt"
)

func main() {
	x := 121
	fmt.Println(x)
	res := isPalindrome(x)
	fmt.Println(res, res == true)

	x = -121
	fmt.Println(x)
	res = isPalindrome(x)
	fmt.Println(res, res == false)

	x = 10
	fmt.Println(x)
	res = isPalindrome(x)
	fmt.Println(res, res == false)

	x = 1231
	fmt.Println(x)
	res = isPalindrome(x)
	fmt.Println(res, res == false)
}
