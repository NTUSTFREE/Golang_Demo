package main

import "fmt"

func MyDefFunc(){
	defer fmt.Println("Done2")
}

func test(a int){
	//defer
	defer func(){
		if a == 1 {
			fmt.Println("Condition 1")
		}else if a == 2{ 
			fmt.Println("Condition 2")
		}
	}()

	fmt.Println("Do something here")

}

func main(){
	
	test(1)
	test(2)
}