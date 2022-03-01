package main

import "fmt"

func stringp(s string) *string{
	return &s
}

func main(){
	
type Person struct{
	FirstName string
	MiddleName *string
	LastName string
}

p := Person{
	FirstName: "Pat",
	MiddleName: stringp("Perry"),
	LastName: "Peterson",
}
fmt.Println(p.FirstName,*(p.MiddleName),p.LastName)
}


//stringp("Perry"),