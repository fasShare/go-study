package main

import (
	"fmt"
	"reflect"
)

// 但类型反射操作
func SimpleReflection() {
	var fv float64 = 9000

	// 简单的获取值或者类型
	rfvp := reflect.ValueOf(&fv)
	fmt.Printf("type = %s|%s\n", rfvp.Kind(), rfvp.Type())

	// panic: reflect: call of reflect.Value.Float on ptr Value
	// fmt.Printf("value=%f\n", rfvp.Float())

	// 指针反射时访问元素的正确方式
	rfv1 := rfvp.Elem()
	fmt.Printf("type = %s|%s value = %f\n", rfv1.Kind(), rfv1.Type(), rfv1.Float())

	if rfv1.CanSet() {
		// 修改值
		rfv1.SetFloat(4.7)
		fmt.Printf("type = %s|%s value = %f\n", rfv1.Kind(), rfv1.Type(), rfv1.Float())
		// 貌似不可以调用
		// rfv1.Set(8888.0)
		// fmt.Printf("type = %s|%s value = %f\n", rfv1.Kind(), rfv1.Type(), rfv1.Float())
	}
}

type Book struct {
	Price float32
	Pages int
	Name  string
}

func DataStructReflection() {
	var book Book = Book{23.5, 199, "The source code analysis of etcd 2.7"}
	fmt.Println("book =", book)

	// panic: reflect: call of reflect.Value.Elem on struct Value
	// 传入book的地址，不然会出现以上panic
	bv := reflect.ValueOf(&book)
	bt := reflect.TypeOf(&book)
	fmt.Println(bv, bt)
	ebv := bv.Elem()
	ebvt := ebv.Type()
	for i := 0; i < ebv.NumField(); i++ {
		ef := ebv.Field(i)
		eft := ebvt.Field(i)
		fmt.Printf("index = %d, Filed = %s, Type = %s, value = %v\n", i, eft.Name, ef.Type(), ef.Interface())
	}
}

func main() {
	SimpleReflection()
	DataStructReflection()
}
