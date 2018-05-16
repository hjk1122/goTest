package str1

var m=make(map[int32]string)
func user(id int32,name string) map[int32]string {
	m[1]="张三"
	m[2]="李四"
	return m
}

