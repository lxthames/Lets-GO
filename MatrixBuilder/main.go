package main

import (
	"fmt"
	"strings"
)

// ─────────────────────────────────────────────
//
//	Add a NEW row
//
// ─────────────────────────────────────────────
func addRow(matrix [][]int, row []int) [][]int {
	return append(matrix, row)
}

// ─────────────────────────────────────────────
//
//	Add a value to an existing row (column append)
//
// ─────────────────────────────────────────────
func addToRow(matrix [][]int, rowIndex int, value int) [][]int {
	if rowIndex < 0 || rowIndex >= len(matrix) {
		fmt.Println("Row index out of range")
		return matrix
	}
	matrix[rowIndex] = append(matrix[rowIndex], value)
	return matrix
}

// ─────────────────────────────────────────────
//
//	Pretty print matrix as a grid
//
// ─────────────────────────────────────────────
func printMatrix(matrix [][]int) {
	fmt.Println("\nCurrent Matrix:")
	for _, row := range matrix {
		// convert []int → "1 2 3"
		parts := make([]string, len(row))
		for i, v := range row {
			parts[i] = fmt.Sprintf("%d", v)
		}
		fmt.Println(strings.Join(parts, " "))
	}
	fmt.Println()
}

func main() {
	// Start with empty matrix ([][]int)
	var matrix [][]int

	// Add rows
	matrix = addRow(matrix, []int{1, 2, 3})
	matrix = addRow(matrix, []int{4, 5, 6})
	matrix = addRow(matrix, []int{7}) // uneven rows allowed

	printMatrix(matrix)

	// Add columns to specific rows
	matrix = addToRow(matrix, 0, 99) // add to row 0
	matrix = addToRow(matrix, 2, 88) // add to row 2
	matrix = addToRow(matrix, 2, 77) // add more

	printMatrix(matrix)

	// Add another completely new row
	matrix = addRow(matrix, []int{9, 9, 9, 9})
	printMatrix(matrix)
}
