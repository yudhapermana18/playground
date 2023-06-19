package basic

import "fmt"

func Function() {
	var num int8 = 3
	ret := escape(num)
	fmt.Println(ret)
}

func escape(num int8) int8 {
	return num
}
