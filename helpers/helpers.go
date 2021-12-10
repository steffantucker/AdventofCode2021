package helpers

import (
	"bufio"
	"os"
	"strconv"
)

// LoadInputLines loads the input assuming each line is a string
func LoadInputLines(f string) (lines []string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func LoadInputNumbers(f string) (lines []int) {
	l := LoadInputLines(f)
	for _, n := range l {
		num, _ := strconv.Atoi(n)
		lines = append(lines, num)
	}
	return
}

func MustAtoI(in string) int {
	n, _ := strconv.Atoi(in)
	return n
}
