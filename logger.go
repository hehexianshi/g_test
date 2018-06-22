package main

import "log"
import "io"
import "os"

func main() {
	l := logger(os.Stdout)
	l.Println("abc")
}

func logger(out io.Writer) *log.Logger {
	l := log.New(out, "[simple]", 0)
	return l
}
