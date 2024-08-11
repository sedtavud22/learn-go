package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(isEmpty(intSlice))
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(isEmpty(float32Slice))
	fmt.Println(sumSlice(float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}