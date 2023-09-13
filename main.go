package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const NAME string = "basic"

var a string = "basic"

type imoocInt int

type Learn struct {
}

type Ilearn interface {
}

const imooc string = "imooc"
const name = "name"

const (
	cat string = "cat"
	dog        = "dog"
)

const apple, num = "apple", 123
const length = len(apple)

const testa = iota
const testb = iota
const (
	testc = iota + 2
	_     = iota
	teste = 3.14
	testd = iota
)

const (
	iotaa = iota * 2
	// with be automatically set to iota * 2 if not declared
	iotab
	iotac
	iotad, iotae = iota, iota * 3
	iotaf, iotag
)

func learnImooc() {
	fmt.Println("leanImooc")
}

func main() {
	fmt.Println("Hello world!")
	learnImooc()
	fmt.Println(imooc)
	fmt.Println(name)

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

	var (
		in    = 123
		float float32
		str   = "123"
	)

	var x, y, z = 1.1, 2, 3
	fmt.Println(in, float, str)
	fmt.Println(x, y, z)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))

	l, _, n := 1, 2, 3
	fmt.Println(l, n)

	m := float32(y)
	fmt.Println(m)

	fmt.Println(testa)
	fmt.Println(testb)
	fmt.Println(testc)
	fmt.Println(testd)
}
