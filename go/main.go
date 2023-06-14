package main

import (
	"fmt"
	"runtime"
)

// showNo - Shows no from 0 to 99
func showNo() {
	for i := 0; i < 2; i++ {
		fmt.Println("value of i=", i)
	}
}

// showAlphabets - shows alphabets from a-z
func showAlphabets() {
	for j := 'a'; j <= 'b'; j++ {
		fmt.Println("value of j=", string(j))
	}
}

func main() {

	/*
		In newer versions of go by default value GOMAXPROCS is set to no of logical processor you have
	*/
	runtime.GOMAXPROCS(1)

	for i := 0; i < 100; i++ {
		fmt.Println("No of i=", i)
		go showNo()
		go showAlphabets()
	}

	a := 0
	fmt.Scanf("%d", &a)
}
