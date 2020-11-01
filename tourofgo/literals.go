package main

import (
	"fmt"
	"math"
)

func main() {
	// 一：array literal
	array2 := [3]string{"a", "b", "c"}
	array2[2] = "d"
	// array
	var myArray [3]string
	myArray[0] = "cn"
	myArray[1] = "中国"

	// 二：slice literal
	slice2 := []string{"a", "b", "c"}
	slice2[2] = "d"
	// make slice, len为3， capacity为10的slice
	mySlice := make([]string, 3, 10)
	for i, s := range mySlice {
		fmt.Println(i, s)
	}

	// 三：map literal
	map2 := map[string]int{
		"x": 9,
		"y": 8,
		"z": 7,
	}
	// make map, size为3
	myMap := make(map[string]int, 2)
	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3
	for k, v := range myMap {
		fmt.Println(k, v)
	}
	for k, v := range map2 {
		fmt.Println(k, v)
	}
	// 四：struct literal
	v7 := Vertex2{1, 2}
	v8 := Vertex2{}
	p9 := &Vertex2{X: 9}
	fmt.Println(v7, v8, p9)

	interfaces()
}

type Vertex2 struct {
	X, Y float64
}
type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex2 implements Abser

	// In the following line, v is a Vertex2 (not *Vertex2)
	// and does NOT implement Abser.
	//a = v
	fmt.Println(a.Abs())
}
