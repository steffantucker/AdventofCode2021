package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

var (
	gammabin = []int{}
	epsbin   = []int{}
)

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\ntest 2: %v\n", puzzle1(toBin(testinput)), puzzle2(toBin(testinput)))
	fmt.Printf("input 1: %v\ninput 2: %v\n", puzzle1(toBin(inputinput)), puzzle2(toBin(inputinput)))
}

func toBin(in []string) (out [][]int) {
	for _, b := range in {
		x := []int{}
		for _, c := range b {
			if c == '1' {
				x = append(x, 1)
			} else {
				x = append(x, 0)
			}
		}
		out = append(out, x)
	}
	return
}

func puzzle1(input [][]int) (out int) {
	gammabin = []int{}
	epsbin = []int{}
	counts := make([]int, len(input[0]))
	for _, b := range input {
		for i, n := range b {
			if n == 1 {
				counts[i]++
			}
		}
	}
	for _, n := range counts {
		if n > len(input)/2 {
			gammabin = append(gammabin, 1)
			epsbin = append(epsbin, 0)
		} else {
			gammabin = append(gammabin, 0)
			epsbin = append(epsbin, 1)
		}
	}
	return binaryNumber(epsbin) * binaryNumber(gammabin)
}

func counts(input [][]int) (out []int) {
	counts := make([]float32, len(input[0]))
	for _, b := range input {
		for i, n := range b {
			if n == 1 {
				counts[i]++
			}
		}
	}
	half := float32(len(input)) / 2
	for _, n := range counts {
		if n == half || n > half {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return
}

func invertCounts(input [][]int) (out []int) {
	counts := make([]float32, len(input[0]))
	for _, b := range input {
		for i, n := range b {
			if n == 1 {
				counts[i]++
			}
		}
	}

	half := float32(len(input)) / 2
	for _, n := range counts {
		if n == half || n > half {
			out = append(out, 0)
		} else {
			out = append(out, 1)
		}
	}
	return
}

func binaryNumber(in []int) int {
	s := []string{}
	for _, n := range in {
		s = append(s, strconv.Itoa(n))
	}
	n, _ := strconv.ParseInt(strings.Join(s, ""), 2, 32)
	return int(n)
}

func copy(in [][]int) (out [][]int) {
	for _, i := range in {
		out = append(out, i)
	}
	return
}

func puzzle2(input [][]int) (out int) {
	o2 := copy(input)
	co2 := copy(input)

	o2numb := scrubbers(o2)
	co2numb := invscrubbers(co2)
	fmt.Printf("%v\n%v\n", o2numb, co2numb)
	return o2numb * co2numb
}

func scrubbers(o2 [][]int) int {
	com := counts(o2)
	for j := range o2[0] {
		for i := 0; i < len(o2); i++ {
			if o2[i][j] != com[j] {
				//fmt.Printf("\n%v\t%v\t%v\n%v\n", com, j, i, o2)
				o2[i] = o2[len(o2)-1]
				o2 = o2[:len(o2)-1]
				if i >= 0 {
					i--
				}
				if len(o2) == 1 {
					return binaryNumber(o2[0])
				}
			}
		}
		com = counts(o2)
	}
	return binaryNumber(o2[0])
}

func invscrubbers(o2 [][]int) int {
	com := invertCounts(o2)
	for j := range o2[0] {
		for i := 0; i < len(o2); i++ {
			if o2[i][j] != com[j] {
				o2[i] = o2[len(o2)-1]
				o2 = o2[:len(o2)-1]
				if i >= 0 {
					i--
				}
				if len(o2) == 1 {
					return binaryNumber(o2[0])
				}
			}
		}
		com = invertCounts(o2)
	}
	return binaryNumber(o2[0])
}
