package main

import (
	"fmt"
	"../test/str1"
	"strconv"
)

func main(){
	//println("hello,world111")
	/*
	var x int32
	x=100
	f:=test(x)
	f()
	*/
	var user str1.User
	user.Id=1
	user.Name="张三"
	user.Age=20
	user.Add="北京"
	fmt.Println(str1.ReturnStr())
	fmt.Printf("编号:%s,姓名:%s,年龄:%s,地址:%s\n",strconv.Itoa(user.Id),
		user.Name,strconv.Itoa(user.Age),user.Add)
	Copy1()

}
func Copy1(){
	s1:=[]int{1,2,3,4,5}
	s2:=[]int{6,7,8,9}
	copy(s1,s2)
	fmt.Println(s1)
}

