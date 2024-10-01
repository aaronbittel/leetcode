package main

import (
	"fmt"
	"reflect"
)

func main() {
	testFunc(intToRoman, []any{3749}, "MMMDCCXLIX")
	testFunc(intToRoman, []any{58}, "LVIII")
	testFunc(intToRoman, []any{1994}, "MCMXCIV")
	testFunc(intToRoman, []any{8}, "VIII")
	testFunc(intToRoman, []any{10}, "X")
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
