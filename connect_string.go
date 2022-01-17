package main

import (
	"bytes"
	"fmt"
)
func main() {
	hammer := "eating a hammer"
	sickle := "death"
	//notice byte cache
	var string Builder bytes.Buffer
	//write string in cache
	srting Builder.Writer String(hammer)
	string Builder.Writer String(sickle)
	fmt.Println(string Builder.String())
}