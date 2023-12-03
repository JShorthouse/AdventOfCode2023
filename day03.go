package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

type Position struct {
    x int
    y int
}

var OFFSETS = []Position{
    { x: -1, y: -1 },
    { x:  0, y: -1 },
    { x:  1, y: -1 },
    { x:  1, y:  0 },
    { x:  1, y:  1 },
    { x:  0, y:  1 },
    { x: -1, y:  1 },
    { x: -1, y:  0 },
}

func main() {
	file, err := os.Open("./input/03.txt")
	if err != nil { panic(err) }

	scanner := bufio.NewScanner(file)

    grid := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

        grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

    cur_row := 0
    cur_col := 0
    parsing_number := false
    number_string := "" 
    adjacent_symbol := false
    gear_position := (*Position)(nil)

    p1_score := 0

    gear_map := make(map[Position][]int)

    for cur_row < len(grid) {
        cur_char := grid[cur_row][cur_col]

        if cur_char >= '0' && cur_char <= '9' {
            parsing_number = true
            number_string += string(cur_char)

            // Check around number for adjacent symbols
            if !adjacent_symbol || gear_position == nil {
                for _, offset := range OFFSETS {
                    test_row := cur_row + offset.y
                    test_col := cur_col + offset.x
                    if test_row > 0 && test_row < len(grid) && test_col > 0 && test_col < len(grid[0]) {
                        symbol := grid[test_row][test_col]
                        if !(symbol >= '0' && symbol <= '9') && symbol != '.' {
                            adjacent_symbol = true;

                            if symbol == '*' {
                                gear_position = &Position{ x: test_col, y: test_row }
                            }
                        }
                    }
                }
            }
        } else {
            parsing_number = false
        }

        // At end of number
        if len(number_string) > 0 && (parsing_number == false || cur_col+1 >= len(grid[0])) {
            conv_num, err := strconv.Atoi(number_string)
            if err != nil { panic(err) }

            if adjacent_symbol {
                p1_score += conv_num
            }

            if gear_position != nil {
                if gear_map[*gear_position] == nil {
                    gear_map[*gear_position] = make([]int, 0, 2)
                }
                gear_map[*gear_position] = append(gear_map[*gear_position], conv_num)
            }

            number_string = ""
            adjacent_symbol = false
            gear_position = nil
        }

        cur_col += 1
        if cur_col >= len(grid[0]) {
            cur_row += 1
            cur_col = 0
        }
    }

    p2_score := 0;

    for _, nums := range gear_map {
        if len(nums) == 2 {
            p2_score += nums[0] * nums[1]
        }
    }

	fmt.Printf("Part 1: %v\n", p1_score)
	fmt.Printf("Part 2: %v\n", p2_score)
}
