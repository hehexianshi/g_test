package main

import (
	"fmt"
)

type Str struct {
	value string
}

func (s *Str) Map(str string) {

	fmt.Println(str)
}

type Simple struct {
	Str
	s string
}

func main() {

	s := &Simple{Str: Str{value: "abc"}, s: "abc"}
	s.Map("abc")

}
