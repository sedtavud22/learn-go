package main

import (
	"fmt"
	"time"
)

func main() {
	// array == fixed length, same type, indexable, contiguous in memory
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [...]int32{1,2,3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	// [4, 5, 6, 7, *, *]
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice2, len(intSlice2), cap(intSlice2))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	/*
	second value == bool
	true if value is in the map
	*/
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	for name, age:= range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i:= 0; i < 3; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("\nTotal time without preallocation: %v", timeLoop(testSlice, n))
	fmt.Printf("\nTotal time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}