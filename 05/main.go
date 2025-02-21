package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func reader(input string) (string, string, error) {
	content, err := os.ReadFile(input)
	if err != nil {
		return "", "", err
	}
	text := string(content)
	split := strings.Split(text, "\n\n")
	return split[0], split[1], nil 
}

func parseInstructs(instructions string) (map[int][]int) {
	instruction := strings.Split(instructions, "\n")
	ins := make(map[int][]int)
	for i := 0; i < len(instruction); i++ {
		split := strings.Split(instruction[i], "|")
		first, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("Error parsing into integer %v\n", err)
		}
		sec, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Error parsing into integer %v\n", err)
		}
		_, ok := ins[first]
		if !ok {
			arr := []int{sec}
			ins[first] = arr 
		} else {
			ins[first] = append(ins[first], sec)
		}
	}
	return ins
}

func parsePages(input string) ([]string) {
	pages := strings.Split(input, "\n")
	return pages[:len(pages)-1]
}

func checkValid (pages string, instructions map[int][]int) (int) {
	page := strings.Split(pages, ",")
	for i := 0; i < len(page); i++ {
		tmp, err := strconv.Atoi(page[i])
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		val, ok := instructions[tmp]
		if ok {
			for j := 0; j < len(page[:i]); j++ {
				for k := 0; k < len(val); k++ {
					tmp, err := strconv.Atoi(page[j])
					if err != nil {
						log.Fatalf("%v\n", err)
					}
					if val[k] == tmp {
						return 0
					}
				}
			}
		}
	}
	output, err := strconv.Atoi(page[len(page)/2])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return output
}

func main() {
	file := "input.txt"
	instructions, pages, err := reader(file)
	if err != nil {
		log.Fatalf("Error reading file")
	}
	ins := parseInstructs(instructions)
	pag := parsePages(pages)
	val := 0
	for i := 0; i < len(pag); i++ {
		val += checkValid(pag[i], ins)
	}
	fmt.Printf("Value: %v\n", val)
}
