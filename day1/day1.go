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

func part1(nums []int64) int64 {
	var prev int64 = -1
	var delta_count int64
	for _, num := range nums {
		if (prev != -1) {
			// we have prev measurement
			if num - prev > 0 {
				// increased
				delta_count++
			}
			
		}
		prev = num
	}
	return delta_count
}

func part2(nums []int64) int64 {
	sliding := make([]int64, 0, len(nums))
	for i := range nums {
		// stop sliding window when slice gets to end
		if i >= len(nums) - 2 {
			break
		}

		r := nums[i:i+3]
		sum := r[0] + r[1] + r[2]
		sliding = append(sliding, sum)
	}
	// now apply the same delta measurement from part 1
	return part1(sliding)
}

func main() {
	inp := bufio.NewScanner(os.Stdin)
	nums := make([]int64, 0, EXPECTED_INPUT_LEN)
	for inp.Scan() {
		i, err := strconv.ParseInt(inp.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	fmt.Printf("Part 1: %v\n", part1(nums))
	fmt.Printf("Part 2: %v\n", part2(nums))
}
