//go:build ignore

package main

import "fmt"


func stringToIntArray(s string, index int, result []int) []int {
	if index >= len(s) {
		return result
	}
	digit := int(s[index] - '0')
	result = append(result, digit)
	return stringToIntArray(s, index+1, result)
}


func intArrayToString(arr []int, index int, result string) string {
	if index >= len(arr) {
		return result
	}
	result = result + string(rune(arr[index]+'0'))
	return intArrayToString(arr, index+1, result)
}

func copyArray(src []int, index int, dest []int) []int {
	if index >= len(src) {
		return dest
	}
	dest = append(dest, src[index])
	return copyArray(src, index+1, dest)
}


func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func makePalindromeRecursive(arr []int, left, right, k int, changed []bool) ([]int, int, []bool, bool) {
	if left >= right {
		return arr, k, changed, true
	}
	
	if arr[left] != arr[right] {
		if k > 0 {
			maxDigit := maxInt(arr[left], arr[right])
			arr[left] = maxDigit
			arr[right] = maxDigit
			changed[left] = true
			changed[right] = true
			k--
		} else {
			return nil, -1, nil, false
		}
	}
	
	return makePalindromeRecursive(arr, left+1, right-1, k, changed)
}


func maximizePalindromeRecursive(arr []int, left, right, k int, changed []bool) ([]int, int) {
	if left >= right {
		return arr, k
	}
	
	if arr[left] < 9 {
		bothChanged := changed[left] && changed[right]
		bothUnchanged := !changed[left] && !changed[right]
		
		if bothChanged && k >= 1 {
			
			arr[left] = 9
			arr[right] = 9
			k--
		} else if bothUnchanged && k >= 2 {
		
			arr[left] = 9
			arr[right] = 9
			k -= 2
		}
	}
	
	return maximizePalindromeRecursive(arr, left+1, right-1, k, changed)
}


func handleMiddleDigit(arr []int, k int) []int {
	if len(arr)%2 == 1 && k > 0 {
		mid := len(arr) / 2
		if arr[mid] < 9 {
			arr[mid] = 9
		}
	}
	return arr
}


func highestPalindrome(s string, k int) string {
	if s == "" || k < 0 {
		return "-1"
	}
	
	n := len(s)
	arr := stringToIntArray(s, 0, []int{})
	changed := make([]bool, n)
	
	result, remainingK, changedArr, success := makePalindromeRecursive(copyArray(arr, 0, []int{}), 0, n-1, k, changed)
	
	if !success || remainingK < 0 {
		return "-1"
	}
	
	changed = changedArr
	
	
	result, remainingK = maximizePalindromeRecursive(result, 0, n-1, remainingK, changed)

	result = handleMiddleDigit(result, remainingK)
	
	return intArrayToString(result, 0, "")
}

func main() {
	fmt.Println("Highest Palindrome Problem ")
	fmt.Println("Cari palindrome terbesar  maksimal k perubahan")
	fmt.Println()
	
	// Test case 1 (dari soal)
	fmt.Println("Test Case 1:")
	input1 := "3943"
	k1 := 1
	output1 := highestPalindrome(input1, k1)
	fmt.Printf("Input: %s, k: %d\n", input1, k1)
	fmt.Printf("Output: %s\n", output1)
	fmt.Println("Expected: 3993")
	fmt.Println()
	
	// Test case 2 dri soal
	fmt.Println("Test Case 2:")
	input2 := "932239"
	k2 := 2
	output2 := highestPalindrome(input2, k2)
	fmt.Printf("Input: %s, k: %d\n", input2, k2)
	fmt.Printf("Output: %s\n", output2)
	fmt.Println("Expected: 992299")
	fmt.Println()
	
	// Test case 3
	fmt.Println("Test Case 3:")
	input3 := "12345"
	k3 := 3
	output3 := highestPalindrome(input3, k3)
	fmt.Printf("Input: %s, k: %d\n", input3, k3)
	fmt.Printf("Output: %s\n", output3)
	fmt.Println("Expected: 99399 atau 99999")
	fmt.Println()
}
