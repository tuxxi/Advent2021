package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	EXPECTED_INPUT_LEN = 2000
)

func part1(report []string) int {
	// assume that all numbers in report are the same size
	bits := len(report[0])
	var gamma, epsilon string

	for bit := 0; bit < bits; bit++ {
		var num0s, num1s int
		for _, num := range report {
			if num[bit] == '0' {
				num0s++
			} else {
				num1s++
			}
		}
		// calc bit at this pos
		var mcb int
		if num0s > num1s {
			mcb = 0
		} else if num0s < num1s {
			mcb = 1
		} else {
			panic("Number of 0s == Number of 1s")
		}
		gamma += fmt.Sprintf("%d", mcb)
		epsilon += fmt.Sprintf("%d", ^(mcb & 1) & 1)
	}
	g, _ := strconv.ParseInt(gamma, 2, 32)
	e, _ := strconv.ParseInt(epsilon, 2, 32)
	return int(g * e)
}

func part2(report []string) int {
	calc := func(oxygen bool) string {
		// slice onto report to track which entries to look at 
		slc := report
		// start at bit n, go left to bit 0
		bit := 0
		for len(slc) > 1 {
			zeroes := make([]string, 0, cap(slc))
			ones := make([]string, 0, cap(slc))
			for _, num := range slc {
				if num[bit] == '0' {
					zeroes = append(zeroes, num)
				} else {
					ones = append(ones, num)
				}
			}
			var greater, lesser []string
			if len(zeroes) > len(ones) {
				greater = zeroes
				lesser = ones
			} else {
				greater = ones
				lesser = zeroes
			}
			if oxygen {
				slc = greater
			} else {
				slc = lesser
			}

			bit++
		}
		return slc[0]
	}

	oxygen, _ := strconv.ParseInt(calc(true), 2, 32)
	co2, _ := strconv.ParseInt(calc(false), 2, 32)
	return int(oxygen * co2)
}

func main() {
	inp := bufio.NewScanner(os.Stdin)
	nums := make([]string, 0, EXPECTED_INPUT_LEN)
	for inp.Scan() {
		nums = append(nums, inp.Text())
	}
	fmt.Printf("Part 1: %v\n", part1(nums))
	fmt.Printf("Part 2: %v\n", part2(nums))
}
