package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func reader(file string) ([]int, []int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)	
	
	list1 := make([]int, 0)	
	list2 := make([]int, 0)
	bit := 0

	for scanner.Scan() {
		text := scanner.Text()	
		i, err := strconv.Atoi(text)
		fmt.Printf("%s\n", i)
		if err != nil {
			log.Fatal(err)
		}
		if (bit % 2 == 0) {
			list1 = append(list1, i)
		} else {
			list2 = append(list2, i)
		}
		bit ++
	}	
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}

	slices.Sort(list1)
	slices.Sort(list2)
	return list1, list2
}

type Counters struct {
	counter []Counter
}

type Counter struct {
	number int 
	count int
}
func main() {
	list1, list2 := reader("input.txt")
	length := len(list1)
	val := 0
	for i := 0; i < length; i++ {
		a := list1[i]
		b := list2[i]
		if a > b {
			val = val + (a - b)
		} else {
			val = val + (b - a)
		}
		fmt.Println(val)
	}
}
