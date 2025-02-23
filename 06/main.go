package main

import(
	"fmt"
	"log"
	"os"
	"strings"
)

func reader(input string) ([]byte, error) {
	content, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	return content, nil
}


func processInput (dat []byte) ([][]byte) {
	
	str := string(dat)
	output := make([][]byte, 0)
	cols := strings.Split(str, "\n")
	for col := 0; col < len(cols) - 1; col++ {
		tmp := make([]byte, 0)
		for row := 0; row < len(cols[0]); row++{
			tmp = append(tmp, cols[col][row])
		}
		output = append(output, tmp)
	}
	return output
}

func walk(input [][]byte) (int) {
	ans := 0
	for col := 0; col < len(input); col++ {
		for row := 0; row < len(input[0]); row++{	
			if input[col][row] == 94 {
				fmt.Printf("Starting loc is %v, %v\n", col, row)
				input[col][row] = 100
				ans = walkUp(input, row, col)
				break
			}
		}
	}
	return ans 
}

func walkUp(input [][]byte, row int, col int) (int) {
	if col - 1  < 0 {
		return countUnique(input) 
	}
	if input[col - 1][row] != 35 {
		input[col-1][row] = 100
		walkUp(input, row, col - 1)
	} else {
		fmt.Printf("Finsihed walking up\n")
		walkRight(input, row, col)
	}
	return countUnique(input)
}


func walkRight(input [][]byte, row int, col int) (int) {
	if row + 1 > len(input[0]) {
		return countUnique(input) 
	}
	if input[col][row + 1] != 35 {
		input[col][row + 1] = 100
		walkRight(input, row + 1, col)
	} else {
		fmt.Printf("Finished walking right\n")
		walkDown(input, row, col)
	}
	return countUnique(input)
}


func walkDown(input [][]byte, row int, col int) (int) {
	if col + 1 > len(input) - 1 {
		return countUnique(input) 
	}
	if input[col + 1][row] != 35 {
		input[col + 1][row] = 100
		walkDown(input, row, col + 1)
	} else {
		fmt.Printf("Finished walking down\n")
		walkLeft(input, row, col)
	}
	return countUnique(input)
}



func walkLeft(input [][]byte, row int, col int) (int) {
	if row - 1 < 0 {
		return countUnique(input) 
	}
	if input[col][row - 1] != 35 {
		input[col][row - 1] = 100
		walkLeft(input, row - 1, col)
	} else {
		fmt.Printf("Finished walking left\n")
		walkUp(input, row, col)
	}
	return countUnique(input)
}



func countUnique(input [][]byte) (int) {
	count := 0
	for col := 0; col < len(input); col++ {
		for row := 0; row < len(input[0]); row++{
			if input[col][row] == 100 {
				count += 1
			}
		}
	}
	return count
}

func main() {
	text := "input.txt"
	dat, err := (reader(text))
	if err != nil {
		log.Fatalf("Error with reader: %v\n", err)
	}
	content := processInput(dat)
	for i := 0; i < len(content); i ++ {
		fmt.Printf("%c\n", content[i])
	}

	steps := walk(content)
	for i := 0; i < len(content); i ++ {
		fmt.Printf("%c\n", content[i])
	}
	fmt.Printf("%v\n", steps)
}
