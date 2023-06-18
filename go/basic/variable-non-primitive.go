package basic

import (
	"fmt"
	"math"
	"reflect"
)

func VariableNonPrimitive() {
	/*
		Array
		sequence of variable primitive with single type
		- have size
		- cannot resize (increase / decrease)
	*/

	// standart declaration
	var a [3]int = [3]int{1, 2, 3}

	// shorthand declaration
	ash := [5]int8{1, 2, 3}

	// standart declaration without initial, value 0
	var ar [5]int16

	// declaration with index
	var ai [5]int32 = [5]int32{1: 3, 3: 4}

	// declaration infer
	var ain [6]int64 = [...]int64{1, 2, 3, 4, 5, 6}

	// max length array
	var amax [math.MaxInt32]int

	fmt.Println(a)
	fmt.Println(ash)
	fmt.Println(ar[:])
	fmt.Println(len(amax))
	fmt.Println(ai)
	fmt.Println(ain)
	fmt.Println(len(ain))
	fmt.Println(cap(ain))
	fmt.Println(reflect.TypeOf(ash).Size())
	fmt.Println(reflect.TypeOf(ar).Size())
	fmt.Println(reflect.TypeOf(ain).Size())
	fmt.Println(reflect.TypeOf(ain).Kind())

	// for _, v := range amax {
	// 	go fmt.Println(v)
	// }

	// for i := 0; i < len(amax); i++ {
	// 	fmt.Println(amax[i])
	// }
	maxi := math.MaxInt64
	fmt.Println(maxi)

	fmt.Println("address a index 1", &a[0])
	for i, v := range a {
		fmt.Println("type data i", reflect.TypeOf(i))
		fmt.Println("type data v", reflect.TypeOf(v))
		fmt.Println("index", i+1, "address", &v)
	}

	// for j := 0; j < maxi; j++ {
	// 	h := j
	// 	if h > 2 {
	// 		h = 2
	// 	}
	// 	go fmt.Println("index", h+1, "address", &a[h])
	// }

	/*
		Slice
		sequence of variable primitive with single type
		- cannot have size
		- can resize
	*/

	// standart declaration
	var s []int

	s = append(s, 20)

	fmt.Println(s)

	si := make([]int16, 0, 2)
	si = append(si, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9)

	sa := new([2]int16)[0:0]
	sa = append(sa, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9)

	fmt.Println(reflect.TypeOf(si).Size())
	fmt.Println(si)
	fmt.Println(len(si))
	fmt.Println(cap(si))
	fmt.Println(reflect.TypeOf(s).Kind())
	fmt.Println(reflect.TypeOf(si).Kind())

	fmt.Println(reflect.TypeOf(sa).Size())
	fmt.Println(sa)
	fmt.Println(len(sa))
	fmt.Println(cap(sa))
	fmt.Println(reflect.TypeOf(sa).Kind())
}
