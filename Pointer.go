package main

import "fmt"

func main(){
	
	var x int = 2 
	var PointerX *int 
	PointerX = &x
	fmt.Println("value of x ", x)
	fmt.Println("value of x ", *PointerX,"(retrieve from pointer)")
	fmt.Println("Address of x ", PointerX)
	fmt.Println("Address of x ", &x)
}