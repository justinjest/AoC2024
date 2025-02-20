package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func reader(input string) ([]byte, error) {
	
	content, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func findAll(input []byte) ([][]byte) {
	re := regexp.MustCompile("mul\\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\\)")
	matches := re.FindAll(input, -1)
	return matches 
}

func multiply(input []byte) (int, error) {
	nums := input[4: len(input) -1]
	split := strings.Split(string(nums), ",")
	firstNum, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, err
	}
	secNum, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, err
	}
	res:= firstNum * secNum	
	return res, nil
}

func main() {
	input := "input.txt"
	data , err := reader(input)
	if err != nil {
		log.Fatalf("Error reading file %v: %v", input, err)
	}
	ans := 0
	matches := findAll(data)
		for i := 0; i < len(matches); i++ {	
		res, err:= multiply(matches[i])
		if err != nil {
			log.Fatalf("Error parsing multiplication %v\n", err)
		}
		ans += res
		fmt.Printf("%v\n", res)
	}
	fmt.Println(ans)
}

