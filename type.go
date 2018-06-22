package main

import (
	"fmt"
)

type FastInvoker interface {
	Invoker(string)
}

type handler func(string, string)

func (h handler) Invoker(abc string) {
	fmt.Println(abc)

}
func is(h interface{}) bool {
	_, ok := h.(FastInvoker)
	return ok
}

func t(h interface{}) {
	switch v := h.(type) {
	case handler:
		fmt.Println("h")
	default:
		fmt.Println(v)

	}
}

func run() func(string, string) {

	/*
		fmt.Println(is(handler(func(a, b string) {

		})))
	*/
	return handler(func(a, b string) {
		fmt.Println("handler")
	})

}

func main() {

	f := run()
	/*
		f("abc", "bcd")
	*/

	handler(f).Invoker("ttt")

	t(handler(f))
}
