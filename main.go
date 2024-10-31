package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	testFunc(findSubstring, []any{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9})
	testFunc(
		findSubstring,
		[]any{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
		[]int{},
	)
	testFunc(
		findSubstring,
		[]any{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}},
		[]int{6, 9, 12},
	)
	testFunc(findSubstring, []any{"barfoobar", []string{"foo", "bar"}}, []int{0, 3})
}

func testFunc(fn any, inputs []any, expected any) {
	var (
		green = "\033[38;5;10m"
		red   = "\033[38;2;255;95;95m"

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
		resultMsg = fmt.Sprint(Colorize("Test passed!", green))
	} else {
		if reflect.TypeOf(expected).Kind() == reflect.String {
			resultMsg = fmt.Sprintf("Test failed! Expected %q, got %q", expected, actual)
		} else {
			resultMsg = fmt.Sprintf("Test failed! Expected %v, got %v", expected, actual)
		}

		resultMsg = Colorize(resultMsg, red)
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
