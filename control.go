package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	if a > 0 {
		fmt.Println("a > 0")
		if a == 1 {
			fmt.Println("a = 1")
		}
	} else if a == 0 {
		fmt.Println("a = 0")
	} else {
		fmt.Println("a < 0")
	}

	var swit interface{}
	swit = 32

	switch swit {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("none")
	}

	for i := 1; i < 10; i++ {
		fmt.Println("imooc")
		time.Sleep(1 * time.Microsecond)
	}

	arr := []string{"a", "b", "c"}
	for key, value := range arr {
		fmt.Print(key, " ")
		fmt.Println(value)
		break
	}

	goto One
	fmt.Print("middle code")
One:
	fmt.Println("code one")
}
