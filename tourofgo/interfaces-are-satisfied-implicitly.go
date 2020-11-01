package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// "var" myName T
	// "var" myName T = T{}
	// "type" myName "struct|interface"{}
	var myArray []string = []string{"ab", "cd", "ef"}
	var str string = "abc"
	// 看上面，知道这个语法啥意思了吗
	// 给定类型，并且，为其赋值。i就是interface的value
	var i I = T{"hello"}
	fmt.Println(str, myArray)

	i.M()
}
