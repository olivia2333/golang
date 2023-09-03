package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const NAME string = "imooc"

var a string = "imooc"

type imoocInt int

type Learn struct {
}

type Ilearn interface {
}

func learnImooc() {
	fmt.Println("leanImooc")
}

func main() {
	fmt.Println("Hello world!")
	learnImooc()

	var a int = 1

	// default zero
	var b int32
	var c int8
	var d uint = 8
	var e float32 = 3.2
	var f bool
	var g byte = 1
	var h string
	var i rune = 3
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(g))
	fmt.Println(unsafe.Sizeof(h))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(h)

	type imooc int32
	// j cannot perform operations with other int32 types
	var j imooc = 32
	fmt.Println(j)
	fmt.Println(reflect.TypeOf(j))
}
