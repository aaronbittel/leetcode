package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	GREEN = "\033[38;5;10m"
	RED   = "\033[38;2;255;95;95m"
)

func main() {
	testFunc(searchMatrix, []any{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3}, true)
	testFunc(searchMatrix, []any{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13}, false)
}

func testFunc(fn any, inputs []any, expected any) {
	var (
		fnValue = reflect.ValueOf(fn)
		fnType  = fnValue.Type()
	)

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
	var resultMsg string

	if reflect.DeepEqual(actual, expected) {
		resultMsg = fmt.Sprint(Colorize("Test passed!", GREEN))
	} else {
		if reflect.TypeOf(expected).Kind() == reflect.String {
			resultMsg = fmt.Sprintf("Test failed! Expected %q, got %q", expected, actual)
		} else {
			resultMsg = fmt.Sprintf("Test failed! Expected %v, got %v", expected, actual)
		}

		resultMsg = Colorize(resultMsg, RED)
	}

	fmt.Println(resultMsg)
	fmt.Println(strings.Repeat("-", StringLen(resultMsg)))
}

func Colorize(s, color string) string {
	return fmt.Sprintf("%s%s%s", color, s, "\033[m")
}

func StringLen(str string) int {
	return utf8.RuneCountInString(StripAnsi(str))
}

var (
	ansiRegex = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
)

func StripAnsi(str string) string {
	return ansiRegex.ReplaceAllString(str, "")
}

func inPlace() {
	input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}

	fmt.Println(input)

	gameOfLife(input)
	correct := true
	for i := 0; i < len(input); i++ {
		if !slices.Equal(input[i], expected[i]) {
			fmt.Println(
				Colorize(
					fmt.Sprintf("Test failed! Expected %v, got %v", expected[i], input[i]),
					RED,
				))
			correct = false
		}
	}

	if correct {
		fmt.Println(fmt.Sprint(Colorize("Test passed!", GREEN)))
	}

	fmt.Println()

	input = [][]int{{1, 1}, {1, 0}}
	expected = [][]int{{1, 1}, {1, 1}}

	fmt.Println(input)

	gameOfLife(input)
	correct = true
	for i := 0; i < len(input); i++ {
		if !slices.Equal(input[i], expected[i]) {
			fmt.Println(
				Colorize(
					fmt.Sprintf("Test failed! Expected %v, got %v", expected[i], input[i]),
					RED,
				))
			correct = false
		}
	}

	if correct {
		fmt.Println(fmt.Sprint(Colorize("Test passed!", GREEN)))
	}
}

func ColorizeIf(x, y any) string {
	var cond bool
	switch x.(type) {
	case int:
		cond = x == y
	default:
		log.Fatalf("unknown type: %v", reflect.TypeOf(x))
	}

	if cond {
		return Colorize(fmt.Sprintf("%v == %v", x, y), GREEN)
	}
	return Colorize(fmt.Sprintf("%v != %v", x, y), RED)
}
