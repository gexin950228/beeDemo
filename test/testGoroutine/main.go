package main

import (
	"fmt"
	"time"
)

var ch1 chan string

func test11(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
		time.Sleep(1 * time.Second)
	}
}

func test1() {
	test11("hi test1")
	ch1 <- "666"
}

func test2() {
	test11("hello test2")
	ch1 <- "hello"
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	go test1()
	go test2()
	for {
		time.Sleep(1 * time.Second)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
