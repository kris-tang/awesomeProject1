package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello,world")
	theme := "狙击start"
	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
	for i :=0; i < len(theme); i++{
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}

	tracer := "死神来了， 死神bye bye"
	comma := strings.Index(tracer, "， ")
	pos := strings.Index(tracer[comma:], "死神")

	 fmt.Println(comma, pos, tracer[comma+pos:])
	angle := "Heros nerver die"

	angele Bytes := []byte(angel)
	for i :=5; i<=10; i++{
		angle Bytes[i] = ''
	}
	fmt.Println(string(angle Bytes))
}