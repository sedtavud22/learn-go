package main

import "fmt"

func main() {
	var p *int32
	var i int32
	fmt.Println(p,i)

	// memory location
	var p2 *int32 = new(int32)
	fmt.Println(p2)

	fmt.Printf("The value p2 points to is: %v", *p2)
	fmt.Printf("\nThe value of i is: %v", i)
	
	// change values stored at memory location
	*p2 = 10
	fmt.Printf("\nThe value p2 points to now is: %v", *p2)

	var p3 *int32 = new(int32)
	// create pointer from address of another var
	// p3 now refers to memory address of i
	// p3 and i now reference the same int32 value
	p3 = &i
	*p3 = 1
	fmt.Printf("\nThe value p3 points to is: %v", *p3)
	fmt.Printf("\nThe value of i is: %v\n", i)

	// copying variables refer to the same data
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)

	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	var newThing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the newThing1 array is: %p", &newThing1)
	var result1 [5]float64 = square2(&newThing1)
	fmt.Printf("\nThe result is: %v", result1)

	fmt.Printf("\nThe value of newThing1 is: %v", newThing1)
}

// create a copy
func square(thing2 [5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)

	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

// take a pointer to an array instead
func square2(thing3 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing3 array is: %p", &thing3)

	for i := range thing3{
		thing3[i] = thing3[i]*thing3[i]
	}
	return *thing3
}
