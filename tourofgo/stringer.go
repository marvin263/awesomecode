package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
// one of the most ubiquitous interface is Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	a := Person{"Marvin", 38}
	b := Person{"Susan", 37}
	c := Person{"Doudou", 7}
	fmt.Println(a, b, c)
}
