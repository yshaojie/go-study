package main

import "fmt"

func main() {
	dog := Dog{"xiaohei", 3}
	dog.eating()
	fmt.Println(dog)

	dog.PointEating()
	fmt.Println(dog)
}

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/**
Go中虽没有class，但依旧有method
通过显示说明receiver来实现与某个类型的组合
只能为同一个包中的类型定义方法
Receiver可以是类型的值或者指针
不存在方法重载
可以使用值或指针来调用方法，编译器会自动完成转换
从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第1个参数）
如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
类型别名不会拥有底层类型所附带的方法
方法可以调用结构中的非公开字段
*/
// 这里为Dog声明了一个方法
func (d Dog) eating() {
	fmt.Println(d.Name, " eating ....")
	d.Name = "eating new name"
}

// PointEating 这里为Dog声明了一个方法
func (d *Dog) PointEating() {
	fmt.Println(d.Name, " eating ....")
	d.Name = "PointEating new name"
}
