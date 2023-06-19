package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const tagName = "example"

type People struct {
	Name string `json:"name" example:"ename"`
	Age  uint8  `json:"age" example:"eage"`
}

func StructGo() {
	/*
		Struct
		sequence of fields
	*/
	// definition
	type Sexa struct {
		x, y int
		u    float32
		_    float32 // padding
		A    *[]int
		F    func()
	}

	// declaration standart
	var sr Sexa

	// short hand declaration
	sexa := Sexa{1, 2, 3.3, 1.2, &[]int{1, 2}, func() { fmt.Println("Func") }}
	sexa2 := Sexa{
		x: 1,
		y: 2,
		u: 3.3,
		A: &[]int{1, 2},
		F: func() {
			fmt.Println("Func")
		},
	}

	// anonymous
	anony := struct {
		name string
		no   int
	}{
		name: "Test",
		no:   2,
	}

	// nested
	nest := struct {
		r    bool
		anom Sexa
	}{
		r:    true,
		anom: sexa2,
	}

	fmt.Println(sr.x, sr.y, sr.u)
	fmt.Println(sexa)
	fmt.Println(sexa2)
	fmt.Println(sexa2.x, sexa2.y, sexa2.A, &sexa2.F)
	fmt.Println(anony)
	fmt.Println(nest)
	fmt.Println(&nest)

	// struct tag
	structTag()
	getStructTag()
}

func structTag() {
	people := People{
		Name: "Man",
		Age:  2,
	}

	fmt.Println(people)

	peopleJson, err := json.MarshalIndent(people, "", "	")
	if err != nil {
		panic("error")
	}

	fmt.Println(string(peopleJson))
	fmt.Println("struct tag")
}

func getStructTag() {
	people := People{
		Name: "Bro",
		Age:  3,
	}

	fmt.Println(people)

	ref := reflect.TypeOf(people)

	fmt.Println(ref.Name())
	fmt.Println(ref.Kind())
	fmt.Println(ref.NumField())

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		// fmt.Println(field)
		// fmt.Println(field.Tag)
		// fmt.Println(field.Type)

		tag := field.Tag.Get(tagName)
		// fmt.Println(tag)

		fmt.Printf("Field %s have tag %s\n", field.Name, tag)
	}
}
