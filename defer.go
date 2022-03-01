package main

import "fmt"

func MyDefFunc(){
	defer fmt.Println("Done2")
}

func test(){
	//defer
	defer fmt.Println("Done")
	//defer fmt.Println("Done1")
	//defer MyDefFunc
	//Body
	fmt.Println("Do something here")

}

func main(){
	test()
}