package main

func main(){
	// 编译器对未赋值的变量,默认初始化为0
	var x int32
	var s = "hello, world"
	// 在函数内部可以省略var关键字, 使用更简单的定义模式
	b := 100
	// 编译器将未使用的局部变量定义当作错误
	println(x, s, b)
}
