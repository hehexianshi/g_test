package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("test.log") //针对test.log文件
	if err != nil {
		log.Fatalf("Open file fail:%v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
			} else {
				break
			}
		}
		fmt.Print(string(line))
	}

}
