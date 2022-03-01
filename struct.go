package main

import "fmt"

func main(){
	
type Student struct {
	Age int
	height int
	weight int
}

p := Student {
	Age: 10,
	height: 180,
	weight: 50,
}

fmt.Println("Age is ", p.Age)
fmt.Println("height is ", p.height)
fmt.Println("weight is ", p.weight)

}