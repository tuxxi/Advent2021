package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EXPECTED_INPUT_LEN = 2000
)

const (
	Up = "up"
	Down = "down"
	Forward = "forward"
)

type Command struct {
	dir string
	amt int64
}

func day1(commands []Command) int64 {
	var hpos int64
	var depth int64
	for _, command := range commands {
		switch command.dir {
			case Up: {
				depth -= command.amt
			}
			case Down: {
				depth += command.amt
			}
			case Forward: {
				hpos += command.amt
			}
			default: panic(fmt.Sprintf("Unknown direction: %v!", command.dir))
		}
	}
	return hpos * depth
}

func day2(commands []Command) int64 {
	var hpos int64
	var depth int64
	var aim int64
	for _, command := range commands {
		switch command.dir {
			case Up: {
				aim -= command.amt
			}
			case Down: {
				aim += command.amt
			}
			case Forward: {
				hpos += command.amt
				depth += (aim * command.amt)
			}
			default: panic(fmt.Sprintf("Unknown direction: %v!", command.dir))
		}
	}
	return hpos * depth

}

func main() {
	inp := bufio.NewScanner(os.Stdin)
	commands := make([]Command, 0, EXPECTED_INPUT_LEN)
	for inp.Scan() {
		parts := strings.Split(inp.Text(), " ")
		dir := parts[0]
		amt, _ := strconv.ParseInt(parts[1], 10, 64)
		commands = append(commands, Command{dir, amt})
	}
	fmt.Printf("Part 1: %v\n", day1(commands))
	fmt.Printf("Part 2: %v\n", day2(commands))
}
