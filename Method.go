package main

import "fmt"
import "time"

type Counter struct{
	total int
	lastUpdate time.Time
}

func (c *Counter) Increment(){
	c.total++
	c.lastUpdate = time.Now()
}


func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdate)
}

func main(){

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

}