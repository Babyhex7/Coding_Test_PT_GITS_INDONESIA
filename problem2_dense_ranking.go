package main

import "fmt"


func removeDuplicates(scores []int, index int, result []int) []int {
	if index >= len(scores) {
		return result
	}
	
	if len(result) == 0 || result[len(result)-1] != scores[index] {
		result = append(result, scores[index])
	}
	return removeDuplicates(scores, index+1, result)
}

func findRank(uniqueScores []int, score, index int) int {
	if index >= len(uniqueScores) {
		return len(uniqueScores) + 1
	}
	if score >= uniqueScores[index] {
		return index + 1
	}
	return findRank(uniqueScores, score, index+1)
}

func processGitsScores(uniqueScores []int, gitsScores []int, index int, result []int) []int {
	if index >= len(gitsScores) {
		return result
	}
	rank := findRank(uniqueScores, gitsScores[index], 0)
	result = append(result, rank)
	return processGitsScores(uniqueScores, gitsScores, index+1, result)
}


func denseRanking(leaderboard []int, gitsScores []int) []int {
	uniqueScores := removeDuplicates(leaderboard, 0, []int{})
	return processGitsScores(uniqueScores, gitsScores, 0, []int{})
}

func printArray(arr []int, index int) {
	if index >= len(arr) {
		return
	}
	if index > 0 {
		fmt.Print(" ")
	}
	fmt.Print(arr[index])
	printArray(arr, index+1)
}

func printIntArray(arr []int, index int) {
	if index >= len(arr) {
		return
	}
	if index > 0 {
		fmt.Print(" ")
	}
	fmt.Print(arr[index])
	printIntArray(arr, index+1)
}

func main() {
	fmt.Println(" Dense Ranking Problem")
	fmt.Println("Sistem ranking pemain skor sama maka rankingnya sama")
	fmt.Println()
	
	// Test case 1 (dari soal)
	fmt.Println("Test Case 1:")
	leaderboard1 := []int{100, 100, 50, 40, 40, 20, 10}
	gitsScores1 := []int{5, 25, 50, 120}
	fmt.Print("Leaderboard: ")
	printIntArray(leaderboard1, 0)
	fmt.Println()
	fmt.Print("GITS Score: ")
	printIntArray(gitsScores1, 0)
	fmt.Println()
	result1 := denseRanking(leaderboard1, gitsScores1)
	fmt.Print("Rankings: ")
	printArray(result1, 0)
	fmt.Println()
	fmt.Println()
	
	// Test case 2
	fmt.Println("Test Case 2:")
	leaderboard2 := []int{100, 90, 90, 80, 75, 60}
	gitsScores2 := []int{50, 65, 77, 90, 102}
	fmt.Print("Leaderboard: ")
	printIntArray(leaderboard2, 0)
	fmt.Println()
	fmt.Print("GITS Score: ")
	printIntArray(gitsScores2, 0)
	fmt.Println()
	result2 := denseRanking(leaderboard2, gitsScores2)
	fmt.Print("Rankings: ")
	printArray(result2, 0)
	fmt.Println()
	fmt.Println()
	
	// Test case 3
	fmt.Println("Test Case 3:")
	leaderboard3 := []int{95, 85, 85, 70, 60, 60, 50}
	gitsScores3 := []int{55, 70, 88, 100}
	fmt.Print("Leaderboard: ")
	printIntArray(leaderboard3, 0)
	fmt.Println()
	fmt.Print("GITS Scores: ")
	printIntArray(gitsScores3, 0)
	fmt.Println()
	result3 := denseRanking(leaderboard3, gitsScores3)
	fmt.Print("Rankings: ")
	printArray(result3, 0)
	fmt.Println()
}
