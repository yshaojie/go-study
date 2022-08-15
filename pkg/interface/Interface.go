package main

import "fmt"

func main() {
	var human Human
	human = new(Man)
	human.hello()
}

type Human interface {
	hello()
	//wc()
}

type Man struct {
}

type Women struct {
}

func (w Women) hello() {
	fmt.Println("Hello, I am a women")
}

func (w Man) hello() {
	fmt.Println("Hello, I am a man")
}
