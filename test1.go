package main

import "fmt"

func main() {
	var a int = 30

	if a < 20 {
		fmt.Print("a<20")
	}
	else {
		fmt.Print("a>=20")
	}
}