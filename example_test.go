package base62_test

import (
	"fmt"
	"github.com/catinello/base62"
)

func ExampleEncode() {
	n := 89569285645
	fmt.Println(base62.Encode(n))
	// Output: 1ZlfarV
}

func ExampleDecode() {
	fmt.Println(base62.Decode("1ZlfarV"))
	// Output: 89569285645 <nil>
}
