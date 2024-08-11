package main

import (
	"fmt"
	"time"
)

// Channels
//  - Hold data
//  - Thread Safe
//  - Listen for data

func main(){
	var c = make(chan int)
	go process(c)
	for i:= range c{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}

func process(c chan int){
	defer close(c)
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Exiting process")
}