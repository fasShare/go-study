package main

import (
	"fmt"
)

func try(panic_source func(), panic_handle func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			panic_handle(err)
		}
	}()
	defer panic_source()
}

func main() {
	fmt.Println("start main!")
	try(func() {
		panic("panic test")
	}, func(e interface{}) {
		fmt.Println(e)
	})
	fmt.Println("exit main!")
}
