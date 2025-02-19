package main

import (
	"fmt"
	"bufio"	
	"log"
	"os"
	"strings"
	"strconv"
)

func convertor(line string) ([]int, error){
	// line is the output of one line from reader
	output := make([]int, 0)
	tmp := strings.Split(line, " ")
	for i := 0; i < len(tmp); i++ {
		num, err := strconv.Atoi(tmp[i])
		if err != nil {
			return make([]int, 0), err
		}
		output = append(output, num)
	}

	return output, nil
}

func analysis(line []int) int {
	// If anything has a dif greater than 3 return early
	incCheck := 0
	dif := line[0] - line[1]
	if (dif >= 1) {
		incCheck = 1
	} else if (dif <= -1) {
		incCheck = -1
	} else {
		fmt.Printf("first two numbers match %v, %v\n", line[0], line[1])
		return 0 
	}	
	for i := 0; i < len(line) - 1; i++ {	
		dif = line[i] - line[i+1]	
		if ((dif > 3 || dif < -3) || (dif * incCheck < 1)) {
		return 0 
		}	
	}
	fmt.Println("Valid")
	return 1
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func Verifier(list []int) int {
	for i := 0; i < len(list); i++ {	
		res := analysis(RemoveIndex(list, i))
		if res == 1 {
			return 1
		}
	}
	return 0

}

func reader(file string) error {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file; %v\n", err)
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validLines := 0
	lineCount := 0
	for scanner.Scan() {
		line, err := convertor(scanner.Text())
		if err != nil {
			fmt.Printf("Error converting line to ints %v\n", scanner.Text())
		}
		valid := Verifier(line)
		lineCount += 1
		validLines += valid
		fmt.Printf("%v valid lines out of %v\n", validLines, lineCount)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error closing file; %v\n", err)
		return err
	}
	fmt.Printf("Valid lines: %v\n", validLines)
	return nil
}

func main() {
	err := reader("input.txt")

	if err != nil {
		log.Fatal(err)
	}
}
