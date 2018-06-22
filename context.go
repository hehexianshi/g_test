package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleSearch)
	http.ListenAndServe(":8080", nil)

}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration("5")
	if err != nil {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}

	defer cancel()

	newContext := NewContext(ctx, "100")
	fmt.Println(FromContext(newContext))
}

func NewContext(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, "abc", value)
}

func FromContext(ctx context.Context) (string, bool) {

	abc, ok := ctx.Value("abc").(string)
	fmt.Println(abc, ok)
	return abc, ok
}

/*
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {

}
*/
