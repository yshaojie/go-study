package main

import (
	"fmt"
)

func main() {
	//DeferFunc()
	//PanicRecover()
	//AnonymousFunc()()
	ClosureFunc()()
}

func DeferFunc() {
	//defer相当于java的finally，就算函数执行错误也能执行defer
	//defer在函数体内可以设置多个,执行时按照倒叙来执行,即由下到上
	defer func() {
		println("第二个执行...")
	}()
	defer func() {
		println("第一个执行...")
	}()
	println("function body")

	panic("我出错了...")
}

func PanicRecover() {
	//Panic和Recover可以用来发出和捕获异常
	//recover只能在defer里使用
	//如果存在多个defer里都有recover,那么只有第一个才能捕获到
	defer func() {
		println("first...")
		err := recover()

		if err != nil {
			println(err)
			fmt.Println("我是捕捉者二号")
		}
	}()
	defer func() {
		err := recover()
		if err != nil {
			println(err)
			fmt.Println("我是捕捉者一号")
		}
		println("second...")
	}()
	println("before panic...")
	panic(12)
	println("after panic...")

}

func AnonymousFunc() func() {
	//匿名函数
	af := func() {
		println("我是匿名函数")
	}
	af()
	return af
}

func ClosureFunc() func() {
	//匿名函数
	af := func() {
		println("我是闭包函数")
	}
	return af
}
