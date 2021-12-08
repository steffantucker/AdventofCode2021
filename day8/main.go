package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\n", puzzle1(testinput))
	fmt.Printf("input 1: %v\n", puzzle1(inputinput))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func splitInput(in []string) (sig, out [][]string) {
	for _, i := range in {
		s := strings.Split(i, " | ")
		ssig := strings.Split(s[0], " ")
		sout := strings.Split(s[1], " ")
		sig = append(sig, ssig)
		out = append(out, sout)
	}
	return
}

func splitInput2(in []string) (sig []map[int][]string, out [][]string) {
	for _, i := range in {
		s := strings.Split(i, " | ")
		ssig := strings.Split(s[0], " ")
		sout := strings.Split(s[1], " ")
		t := make(map[int][]string)
		for _, part := range ssig {
			t[len(part)] = append(t[len(part)], part)
		}
		sig = append(sig, t)
		out = append(out, sout)
	}
	return
}
func puzzle1(input []string) (out int) {
	_, outs := splitInput(input)
	for _, row := range outs {
		for _, disp := range row {
			if len(disp) == 2 || len(disp) == 3 || len(disp) == 4 || len(disp) == 7 {
				out++
			}
		}
	}
	return
}

func puzzle2(input []string) (out int) {
	lentosig, outs := splitInput2(input)
	for i, l := range lentosig {
		wires := wire(l)
		out += solve(wires, outs[i])
	}

	return
}

func solve(w map[string]int, outs []string) (sum int) {
	sumstring := ""
	for _, n := range outs {
		num, ok := w[order(n)]
		if !ok {
			panic("number not found")
		}
		sumstring = fmt.Sprintf("%v%v", sumstring, num)
	}
	number, _ := strconv.Atoi(sumstring)
	return number
}

func wire(lentosig map[int][]string) map[string]int {
	sigstonum := make(map[string]int)
	for n, s := range lentosig {
		ordered := []string{}
		for _, o := range s {
			ordered = append(ordered, order(o))
		}
		lentosig[n] = ordered
	}
	one := lentosig[2][0]
	sigstonum[one] = 1
	seven := lentosig[3][0]
	sigstonum[seven] = 7
	four := lentosig[4][0]
	sigstonum[four] = 4
	eight := lentosig[7][0]
	sigstonum[eight] = 8
	ttf := twothreefive(lentosig[5], one, four)
	sigstonum[ttf[2]] = 2
	sigstonum[ttf[3]] = 3
	sigstonum[ttf[5]] = 5
	sn := sixtyninezero(lentosig[6], four, ttf[5])
	sigstonum[sn[6]] = 6
	sigstonum[sn[9]] = 9
	sigstonum[sn[0]] = 0
	return sigstonum
}

func sixtyninezero(ins []string, four, five string) map[int]string {
	if len(ins) > 3 {
		panic("too many with 6 segments")
	}
	out := make(map[int]string)
	zeroi, zero := findAinBwithCount(five, ins, 2)
	ninei, nine := findAinBwithCount(four, ins, 2)
	out[0] = zero
	out[9] = nine
	for i, s := range ins {
		if i != zeroi && i != ninei {
			out[6] = s
			return out
		}
	}
	panic("no 0!")
}

func findAinBwithCount(a string, b []string, count int) (int, string) {
	for i, in := range b {
		if len(removeAfromB(a, in)) == count {
			return i, in
		}
	}
	panic("not found!")
}

func twothreefive(ins []string, one, four string) map[int]string {
	if len(ins) > 3 {
		panic("too many with 5 segments")
	}
	out := make(map[int]string)
	threei, three := find3(ins, one)
	fivei, five := find5(ins, one, four)
	out[3] = three
	out[5] = five
	for i, s := range ins {
		if i != threei && i != fivei {
			out[2] = s
			return out
		}
	}
	panic("no 2!")
}

func removeAfromB(a, b string) string {
	of := strings.Split(b, "")
	c := []string{}
	for _, o := range of {
		if !strings.Contains(a, o) {
			c = append(c, o)
		}
	}
	return strings.Join(c, "")
}

func find5(ins []string, one, four string) (int, string) {
	comp := removeAfromB(one, four)
	for i, in := range ins {
		if len(removeAfromB(comp, in)) == 3 {
			return i, in
		}
	}
	panic("no 5")
}
func find3(ins []string, one string) (int, string) {
	for i, in := range ins {
		if len(removeAfromB(one, in)) == 3 {
			return i, in
		}
	}
	panic("no 3")
}

func order(input string) string {
	in := []rune(input)
	for i := range in {
		j := i + 1
		for j = range in {
			if in[i] < in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
	return string(in)
}
