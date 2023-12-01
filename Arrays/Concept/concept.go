package main

import "fmt"

func main() {
	// Declaration and Init od slice
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array elements : %v\n", arr)

	// Accessing elements --- O(1)
	fmt.Printf("Element at index 2: %d \n", arr[2])

	// Modifying elements -- O(1)
	arr[2] = 6
	fmt.Printf("After modifying element at index 2 : %v\n", arr)

	// Inserting an element at position 2 -- O(n)
	index := 2
	value := 7

	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	fmt.Printf("After inserting 7 at index 2 : %v\n", arr)

	// Delete elements at position 2 -- O(n)
	index = 2
	arr = append(arr[:index], arr[index+1:]...)

	fmt.Printf("After deleting element at index 2 : %v", arr)

}
