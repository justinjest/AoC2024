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
func findDo(input []byte) ([][]int) {
	re := regexp.MustCompile("do\\(\\)")	
	matches := re.FindAllIndex(input, -1)
	return matches
}

func findDont(input []byte) ([][]int) {
	re := regexp.MustCompile("don't\\(\\)")
	matches := re.FindAllIndex(input, -1)
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

func listStartOnly(input [][]int) ([]int) {
	output := make([]int, 0)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i][0])
	}
	return output
}
func splitData(input []byte, startPoints []int, endPoints[]int) ([]byte) {
	start := 0
	end := endPoints[0]	
	data := input[start:end]
	length := len(input)
	lenDo := len(startPoints)
	lenDont := len(endPoints)
	fmt.Printf("length %v, do %v, dont %v\n", length, lenDo, lenDont)
	fmt.Printf("%s\n", data)

	for i := 0; i < lenDo; i++ {
		if startPoints[i] > end {
			end = nextLargest(endPoints, startPoints[i])
			if end == 0 {
				end = length
			}
			start = startPoints[i]
			data =append(data, input[start:end]...)
		}
	}
	return data 
}

func nextLargest(data []int, val int) int {
	for i := 0; i < len(data); i++ {
		if data[i] > val {
			return data[i]
		}
	}
	return 0
}
func main() {
	input := "input.txt"
	data , err := reader(input)
	if err != nil {
		log.Fatalf("Error reading file %v: %v", input, err)
	}
	ans := 0
	doLocs := listStartOnly(findDo(data))
	dontLocs:=  listStartOnly(findDont(data))
	dataCommands := splitData(data, doLocs, dontLocs)
	matches := findAll(dataCommands)
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

