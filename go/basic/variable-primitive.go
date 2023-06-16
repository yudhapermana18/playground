package basic

import "fmt"

func VariablePrimitive() {
	var x interface{}

	// numeric types
	var i int           // integer default int32, int8, int16, int32, int64
	var i8 int8         // -128 ~ 127
	var i16 int16       // -32768 ~ 32767
	var i32 int32       // -2147483648 to 2147483647
	var i64 int64       // -9223372036854775808 to 9223372036854775807
	var ui uint         // unsigned integer default uint32, uint8, uint16, uint32, uint64
	var ui8 uint8       // 0 to 255
	var ui16 uint16     // 0 to 65535
	var ui32 uint32     // 0 to 4294967295
	var ui64 uint64     // 0 to 18446744073709551615
	var f32 float32     // IEEE-754 32-bit
	var f64 float64     // IEEE-754 64-bit
	var c64 complex64   // complex numbers with float32 real and imaginary parts
	var c128 complex128 // complex numbers with float64 real and imaginary parts
	var r rune          // alias for int32
	var b byte          // alias for uint8
	var uptr uintptr    // an unsigned integer large enough to store the uninterpreted bits of a pointer value

	var s string // sequence of character
	var p *int   // pointer integer

	fmt.Printf("interface: %v\n", x)
	fmt.Printf("integer: %v\n", i)
	fmt.Printf("integer 8: %v\n", i8)
	fmt.Printf("integer 16: %v\n", i16)
	fmt.Printf("integer 32: %v\n", i32)
	fmt.Printf("integer 64: %v\n", i64)
	fmt.Printf("unsigned integer: %v\n", ui)
	fmt.Printf("unsigned integer 8: %v\n", ui8)
	fmt.Printf("unsigned integer 16: %v\n", ui16)
	fmt.Printf("unsigned integer 32: %v\n", ui32)
	fmt.Printf("unsigned integer 64: %v\n", ui64)
	fmt.Printf("float 32: %v\n", f32)
	fmt.Printf("float 64: %v\n", f64)
	fmt.Printf("complex 64: %v\n", c64)
	fmt.Printf("complex 128: %v\n", c128)
	fmt.Printf("byte: %v\n", b)
	fmt.Printf("uninterpreted pointer: %v\n", uptr)
	fmt.Printf("rune: %v\n", r)

	fmt.Printf("string: %v\n", s)
	fmt.Printf("pointer: %v\n", p)

	x = "a"
	i = 1_2
	i8 = -128
	i16 = -32768
	i32 = -2147483648
	i64 = -9223372036854775808
	ui = 3_4
	ui8 = 255
	ui16 = 65535
	ui32 = 4294967295
	ui64 = 18446744073709551615
	f32 = -99999999999999999999999999999999999999
	f64 = -9999999999999999999999999999999999999999999999999999999999999999999999999999
	b = 23
	uptr = 0xc82000c290
	r = '\n'

	s = "Hello"
	p = &i

	fmt.Println("---After assigment---")

	fmt.Printf("interface: %v\n", x)
	fmt.Printf("integer: %v\n", i)
	fmt.Printf("integer 8: %v\n", i8)
	fmt.Printf("integer 16: %v\n", i16)
	fmt.Printf("integer 32: %v\n", i32)
	fmt.Printf("integer 64: %v\n", i64)
	fmt.Printf("unsigned integer: %v\n", ui)
	fmt.Printf("unsigned integer 8: %v\n", ui8)
	fmt.Printf("unsigned integer 16: %v\n", ui16)
	fmt.Printf("unsigned integer 32: %v\n", ui32)
	fmt.Printf("unsigned integer 64: %v\n", ui64)
	fmt.Printf("float 32: %v\n", f32)
	fmt.Printf("float 64: %v\n", f64)
	fmt.Printf("complex 64: %v\n", c64)
	fmt.Printf("complex 128: %v\n", c128)
	fmt.Printf("byte: %v\n", b)
	fmt.Printf("uninterpreted pointer: %d\n", uptr)
	fmt.Printf("rune: %v\n", r)

	fmt.Printf("string: %v\n", s)
	fmt.Printf("pointer: %v\n", p)

	var by []byte = []byte("Hello Brother")
	var rn []rune = []rune{1572, 100}
	var st string
	for _, v := range by {
		st += string(v)
	}

	fmt.Println(st)

	fmt.Println(by)
	fmt.Printf("%c\n", by)

	fmt.Println(rn)
	fmt.Printf("%c\n", rn)
}
