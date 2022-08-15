package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	user := User{
		Name: "yue",
		Age:  12,
	}
	reflection(user)
	value := reflect.ValueOf(user)
	method := value.MethodByName("Goto")
	args := []reflect.Value{reflect.ValueOf("Haidian ...")}
	method.Call(args)

}

type User struct {
	Name string
	Age  int
}

func (user User) Goto(address string) time.Duration {
	fmt.Println(user.Name, " go to ", address)
	return time.Duration(12)
}

func (user User) Eating() {
	fmt.Println(" eating .....")
}

func reflection(any2 any) {
	t := reflect.TypeOf(any2)
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(any2)
	fmt.Println("Fields:", v)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method)
		fmt.Println(method.PkgPath)
		fmt.Println(method.Name)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
	}
}
