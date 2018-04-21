package main

import (
	"fmt"
	"../test/str1"
)

func main(){
	//println("hello,world111")
	/*
	var x int32
	x=100
	f:=test(x)
	f()
	*/
	fmt.Println(str1.ReturnStr())
	Copy1()

}
func Copy1(){
	s1:=[]int{1,2,3,4,5}
	s2:=[]int{6,7,8,9}
	copy(s1,s2)
	fmt.Println(s1)
}

