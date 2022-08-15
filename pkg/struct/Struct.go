package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	student := Student{"yue", "bg", 12}
	bytes, _ := json.Marshal(student)
	println(string(bytes))
	fmt.Println(student)
	//println(student)

	student1 := &Student{Name: "yueyue"}
	bytes1, _ := json.Marshal(student1)
	println(string(bytes1))
	println(student1)

	ValueCopy(student)
	fmt.Println(student)

	//这里会修改student的内容
	Point(&student)
	//更新后的值
	fmt.Println(student)

}

type Student struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Age    int    `json:"age"`
}

// ValueCopy 不带指针就是值拷贝
func ValueCopy(student Student) {
	student.Name = "new name"
	fmt.Println(student)
}

// Point 带指针的话会修改内容
func Point(student *Student) {
	student.Name = "new name"
	fmt.Println(student)
}
