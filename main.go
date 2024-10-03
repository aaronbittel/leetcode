package main

import (
	"fmt"
	"reflect"
)

func main() {
	testFunc(twoSum, []any{[]int{2, 7, 11, 15}, 9}, []int{1, 2})
	testFunc(twoSum, []any{[]int{2, 3, 4}, 6}, []int{1, 3})
	testFunc(twoSum, []any{[]int{-1, 0}, -1}, []int{1, 2})
	testFunc(twoSum, []any{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8}, []int{4, 5})
}

func testFunc(fn any, inputs []any, expected any) {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()

	if fnType.Kind() != reflect.Func {
		fmt.Println("Error: Provided argument is not a function")
		return
	}

	if len(inputs) != fnType.NumIn() {
		fmt.Printf("Error: Expected %d inputs, got %d\n", fnType.NumIn(), len(inputs))
		return
	}

	in := make([]reflect.Value, len(inputs))
	for i, input := range inputs {
		in[i] = reflect.ValueOf(input)
	}

	fmt.Print("Inputs: ")
	for _, i := range in {
		if reflect.TypeOf(i).Kind() == reflect.String {
			fmt.Printf("%q ", i)
		} else {
			fmt.Printf("%+v ", i)
		}
	}
	fmt.Println()

	result := fnValue.Call(in)

	if len(result) != 1 {
		fmt.Println("Error: Function returned more than one value")
		return
	}

	actual := result[0].Interface()

	if reflect.DeepEqual(actual, expected) {
		fmt.Println("Test passed!")
	} else {
		if reflect.TypeOf(expected).Kind() == reflect.String {
			fmt.Printf("Test failed! Expected %q, got %q\n", expected, actual)
		} else {
			fmt.Printf("Test failed! Expected %v, got %v\n", expected, actual)
		}
	}

	fmt.Println("-------------")
}
