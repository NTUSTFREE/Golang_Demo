package main

import "fmt"

func AddTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals{
		out = append(out, base+v)
	}
	return out	
}

func main(){

	fmt.Println(AddTo(3))
	fmt.Println(AddTo(3,2))
	fmt.Println(AddTo(3,2,4,6,8))
	a := []int{4,3}
	fmt.Println(AddTo(3,a...))
	fmt.Println(AddTo(3,[]int{1,2,3,4,5}...))

}