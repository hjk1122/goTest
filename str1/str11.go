package str1

import (
	"fmt"
)

var Pack1Int int = 42
var PackFloat = 3.14
type User struct {
	Id int
	Name string
	Age int
	Add string
}
func ReturnStr() int {
	return Pack1Int
}
func Ptr()  {
	s := "good bye"
	var p *string = &s
	*p = "你好"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
/*
func Cache(){
	c, err := cache.NewCache("memory", `{"interval":60}`)
	err = c.Put("nums", 12, 0)
	if err != nil {
		log.Fatal("err")
	}
	v := c.Get("nums")
	log.Println(v)
	c.Incr("nums")
	v = c.Get("nums")
	log.Println(v)
	c.Decr("nums")
	v = c.Get("nums")
	log.Println(v)
	c.Delete("nums")
	if c.IsExist("nums") {
		log.Fatal("delete err")
	}
}
*/