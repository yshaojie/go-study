package main

func main() {
	Range()
}

func Point() {
	age := 12
	println(&age)
	println(age)
	p := &age
	println(p)
	var pp = &age
	println(pp)
}

func Range() {
	nums := []int{2, 3, 4}
	for index, num := range nums {
		println("index=", index, " value=", num)
	}
}
