package main

import "fmt"

func main() {
	ids := []int{23, 24, 56, 67, 76}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Sum of ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum is %d\n", sum)

	// Range with map
	employees := map[int]string{1: "John", 2: "Brad"}

	for key, value := range employees {
		fmt.Printf("%d: %s\n", key, value)
	}

	// Get only keys
	for key := range employees {
		fmt.Printf("key: %d\n", key)
	}
}
