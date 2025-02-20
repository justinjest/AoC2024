package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
// x = 88
// m = 77
// a = 65
// s = 83
func reader(input string) ([][]byte, error) {
	content, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	data := make([][]byte, 0)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func findXMAS (input [][]byte, x, y int) (result int) {

	ltor := ltor(input, x, y)
	rtol := rtol(input, x, y)
	utod := utod(input, x, y)
	dtou := dtou(input, x, y)
	diaur := diaur(input, x, y)
	diaul := diaul(input, x, y)
	diadr := diadr(input, x, y)
	diadl := diadl(input, x, y)
	return ltor + rtol + utod + dtou + diaur + diaul + diadr + diadl
}

func diadr (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x < 3 || y > len(input[0]) - 4 {
		return 0
	}
	if input[x-1][y+1] == 77 && input[x-2][y+2] == 65 && input[x-3][y+3] == 83 {

		return 1
	}
	return 0

}

func diadl (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x < 3 || y < 3 {
		return 0
	}
	if input[x-1][y-1] == 77 && input[x-2][y-2] == 65 && input[x-3][y-3] == 83 {
		return 1
	}
	return 0

}

func diaul (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x > len(input) - 4 || y < 3 {
		return 0
	}
	if input[x+1][y-1] == 77 && input[x+2][y-2] == 65 && input[x+3][y-3] == 83 {
		return 1
	}
	return 0

}
func diaur (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x > len(input) - 4 || y > len(input[0]) - 4 {	
		return 0
	}
	if input[x+1][y+1] == 77 && input[x+2][y+2] == 65 && input[x+3][y+3] == 83 {
		return 1
	}
	return 0

}
func ltor (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x > len(input) || y > len(input[0]) - 4 {
		return 0
	}
	if input[x][y+1] == 77 && input[x][y+2] == 65 && input[x][y+3] == 83 {
		return 1
	}
	return 0

}

func rtol (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x > len(input) || y <  3 {	
		return 0
	}
	if input[x][y-1] == 77 && input[x][y-2] == 65 && input[x][y-3] == 83 {
		return 1
	}
	return 0

}


func utod (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x > len(input) - 4 || y > len(input[0]) {
		return 0
	}
	if input[x+1][y] == 77 && input[x+2][y] == 65 && input[x+3][y] == 83 {
		return 1
	}
	return 0

}

func dtou (input [][]byte, x, y int) (result int) {

	// Searching only for left to right
	if x <  3 || y > len(input[0]) {
		return 0
	}
	if input[x-1][y] == 77 && input[x-2][y] == 65 && input[x-3][y] == 83 {
		return 1
	}
	return 0

}
func main() {
	file := "input.txt"
	result := 0
	data, err := reader(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	fmt.Printf("Rows: %v Cols: %v\n", len(data), len(data[0]))
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {
			if data[x][y] == 88 {
				result += findXMAS(data, x, y)	
			}
		}
	}
	fmt.Printf("Found %v xmas\n", result)
}
