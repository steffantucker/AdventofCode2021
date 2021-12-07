package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

type fish struct {
	age   int
	birth int
}

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	//fmt.Printf("test 1: %v\n", puzzle1(testinput))
	//fmt.Printf("input 1: %v\n", puzzle1(inputinput))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func puzzle1(input []string) (out int) {
	fishes := loadFish(input[0])
	for i := 0; i < 80; i++ {
		newFish := []fish{}
		for i := range fishes {
			fishes[i].age--
			if fishes[i].age < 0 {
				fishes[i].age = 6
				newFish = append(newFish, fish{age: 8})
			}
		}
		fishes = append(fishes, newFish...)
	}
	return len(fishes)
}

func puzzle2(input []string) (out uint64) {
	fishes := loadFishMap(input[0])
	for i := 0; i < 256; i++ {
		for j := 0; j <= 9; j++ {
			fishes[j-1] = fishes[j]
		}
		fishes[6] += fishes[-1]
		fishes[8] += fishes[-1]
		fishes[-1] = 0
	}
	for _, f := range fishes {
		out += uint64(f)
	}
	return
}

func loadFish(in string) (out []fish) {
	f := strings.Split(in, ",")
	for _, a := range f {
		n, _ := strconv.Atoi(a)
		out = append(out, fish{age: n, birth: 0})
	}
	return
}

func loadFishMap(in string) (out map[int]int) {
	out = make(map[int]int, 11)
	f := strings.Split(in, ",")
	for _, a := range f {
		n, _ := strconv.Atoi(a)
		out[n]++
	}
	return
}
