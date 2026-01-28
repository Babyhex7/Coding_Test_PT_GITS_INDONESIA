//go:build ignore

package main

import (
	"fmt"
	"strconv"
)


func calculateSloane(n int) int {
	return (n*(n-1))/2 + 1
}


func generateSloaneSequence(n, current int, result []int) []int {
	if current > n {
		return result
	}
	result = append(result, calculateSloane(current))
	return generateSloaneSequence(n, current+1, result)
}

func formatOutput(arr []int, index int, result string) string {
	if index >= len(arr) {
		if len(result) > 0 && result[len(result)-1] == '-' {
			return result[:len(result)-1]
		}
		return result
	}
	result = result + strconv.Itoa(arr[index])
	if index < len(arr)-1 {
		result = result + "-"
	}
	return formatOutput(arr, index+1, result)
}


func sloaneOEIS(n int) string {
	if n < 0 {
		return ""
	}
	sequence := generateSloaneSequence(n, 1, []int{})
	return formatOutput(sequence, 0, "")
}

func main() {
	fmt.Println(" A000124 Sloane's OEIS ")
	fmt.Println("rumus: a(n) = n(n-1)/2 + 1")
	fmt.Println()
	
	// Test case 1 (dari soal)
	input1 := 7
	output1 := sloaneOEIS(input1)
	fmt.Printf("Test Case 1:\n")
	fmt.Printf("Input: %d\n", input1)
	fmt.Printf("Output: %s\n", output1)
	fmt.Println()
	
	// Test case 2
	input2 := 5
	output2 := sloaneOEIS(input2)
	fmt.Printf("Test Case 2:\n")
	fmt.Printf("Input: %d\n", input2)
	fmt.Printf("Output: %s\n", output2)
	fmt.Println()
	
	// Test case 3
	input3 := 10
	output3 := sloaneOEIS(input3)
	fmt.Printf("Test Case 3:\n")
	fmt.Printf("Input: %d\n", input3)
	fmt.Printf("Output: %s\n", output3)
	fmt.Println()
}
