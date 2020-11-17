package main

import (
	"fmt"
	"libs"
	"libs/lib2"
)


func main(){
	var i1, i2 int = 1, 2
	var i3 = 3
	i4 := 4;
	numbers := [5]int{1, 2, 3, 4, 5}
	numbers2 := [...]int{1, 2, 3, 4}
	numbers2 = numbers
	fmt.Println(libs.Func1("world"),libs.Add(i1,i2))
	fmt.Println(libs.Func1("world"),lib2.Neg(i3,i4))
}
