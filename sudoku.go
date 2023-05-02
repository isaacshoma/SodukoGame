// Isaac Shoma
package main

// Import format for input and output operations
import (
	"fmt"
)

// Define size equal to 9 for the sudoku grid
const size = 9

// Define the sudoku type as a 9x9 array of integers
type sudoku [size][size]int

// Define a method for the sudoku type to print the puzzle in sukoko game format
// The print method outputs the current state of the sudoku puzzle in a 9x9 grid format
func (s sudoku) print() {
	// Prints separators for every 3rd coloumn
	fmt.Println("-------------------------")
	for i, row := range s {
		// Prints separators for every 3rd row
		for j, value := range row {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", value)
		}
		fmt.Println("|")
		if (i+1)%3 == 0 {
			fmt.Println("|-----------------------|")
		}
	}
}

// This method takes 3 integers and checks if number is a possible value in that position
func (s sudoku) possible(x, y, number int) bool {
	// Check row and columns if any of the cells contain a number
	for i := 0; i < size; i++ {
		// If cells contain a number, it was not a possible number
		if s[x][i] == number || s[i][y] == number {
			return false
		}
	}

	// Checking the individual cells in the boxes
	// The index of the first row of the 3x3 sub-grid containing cell
	rowIndex := (x / 3) * 3

	// The index of the first column of the 3x3 sub-grid containing cell
	columnIndex := (y / 3) * 3

	// Iterate through all the cells in the boxes to see if they contain number
	for i := rowIndex; i < rowIndex+3; i++ {
		// If cell in boxes already have number, Return false
		for j := columnIndex; j < columnIndex+3; j++ {
			if s[i][j] == number {
				return false
			}
		}
	}
	return true
}

// Define a method for the puzzle type to check if the puzzle is complete
func (s sudoku) complete() bool {
	// iterate through every row cell to check if a value is 0
	for i := 0; i < size; i++ {
		// iterate through every row cell to check if a value is 0
		for j := 0; j < size; j++ {
			// if all have non-zero values, the sudoku solver is complete
			if s[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// Define a method for the sudoku type to set a number in a certain position
func (s *sudoku) place(x, y, number int) bool {
	// Rule for the number being input
	if number < 1 || number > 9 {
		return false
	}
	// If the coordinates and the number are not a possible number
	if !s.possible(x, y, number) {
		return false
	}
	// If the number is a possible position
	// this function will update all the values of sudoku at that position
	(*s)[x][y] = number
	return true
}

// Define a function to get user input for the next move
func userInput() (x, y, number int) {
	fmt.Print("Enter row, column, and number seperated by spaces (e.g. 3 2 1): ")
	// Scan the users input, and takes the memory address x y number as arguments
	// assigns users input in them
	fmt.Scan(&x, &y, &number)
	// convert the users input to start at index 0
	return x - 1, y - 1, number
}

// Main function to run the Sudoku solver
func main() {
	s := sudoku{}
	for !s.complete() {
		s.print()
		row, column, number := userInput()
		if !s.place(row, column, number) {
			fmt.Println("Invalid input, please try again.")
		}
	}
	fmt.Println("Congratulations, you have solved the puzzle!")
	s.print()
}
