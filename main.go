package main

import (
	"fmt"
	"reflect"
)

func main() {
	testFunc(candy, []any{[]int{1, 0, 2}}, 5)
	testFunc(candy, []any{[]int{1, 2, 2}}, 4)
	testFunc(candy, []any{[]int{1, 0, 2, 3, 2, 0, 5, 2, 1}}, 17)
	testFunc(candy, []any{[]int{1, 3, 2, 2, 1}}, 7)
	testFunc(candy, []any{[]int{1, 2, 87, 87, 87, 2, 1}}, 13)
	testFunc(candy, []any{[]int{1, 2, 3, 1, 0}}, 9)
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
		fmt.Printf("%+v ", i)
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
		fmt.Printf("Test failed! Expected %v, got %v\n", expected, actual)
	}

	fmt.Println("-------------")
}
