package main

import "fmt"

func Swap(a, b int)int{
	var temp int
	temp = a
	a = b
	b = temp
	return 0
}

func main(){
	var num1 ,num2 int = 1,2
	fmt.Println("a is",num1,"b is ",num2)
	Swap(num1, num2)
	fmt.Println("a is",num1,"b is ",num2)
}