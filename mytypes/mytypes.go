package mytypes

import "strings"

type MyInt int

// Twice multiplies its receiver by 2 and returns the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MyBuilder struct {
	Contents strings.Builder
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (mi *MyInt) Double() {
	*mi *= 2
}
