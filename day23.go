package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
)

const (
    UP = iota
    DOWN
    LEFT
    RIGHT
)

type Position struct {
    x int
    y int
}

var OFFSETS = [4]Position {
    Position{ x:  0, y: -1 },  // UP
    Position{ x:  0, y:  1 },  // DOWN
    Position{ x: -1, y:  0 },  // LEFT
    Position{ x:  1, y:  0 },  // RIGHT
}

func main() {
    file, err := os.Open("./input/23.txt")
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

    var start_pos Position
    var end_pos Position

    for idx, char := range grid[0] {
        if char == '.' { 
            start_pos = Position{ y: 0, x: idx }
        }
    }
    for idx, char := range grid[len(grid)-1] {
        if char == '.' { 
            end_pos = Position{ y: len(grid)-1, x: idx }
        }
    }

    end_distances := make([]int, 0)

    position_stack := make([]Position, 1)
    direction_stack := make([]int, 1)

    position_stack[0] = start_pos
    direction_stack[0] = UP

    for len(position_stack) > 0 {
        cur_pos := position_stack[len(position_stack)-1]
        cur_dir := direction_stack[len(direction_stack)-1]

        if cur_dir >= 4 {
            // Exhasted all directions for this position, pop the stack
            position_stack = position_stack[:len(position_stack)-1]
            direction_stack = direction_stack[:len(direction_stack)-1]
            continue
        }
        
        cur_offset := OFFSETS[cur_dir]
        target_pos := Position{ x: cur_pos.x + cur_offset.x, y: cur_pos.y + cur_offset.y }

        valid_pos := true

        if target_pos == end_pos {
            valid_pos = false
            end_distances = append(end_distances, len(position_stack))
        } else if target_pos.y < 0 || target_pos.x < 0 || target_pos.y >= len(grid) || target_pos.x >= len(grid[0]) {
            valid_pos = false
        } else {
            char := grid[target_pos.y][target_pos.x]
            if char == '#' || (char == '>' && cur_dir != RIGHT) || (char == '<' && cur_dir != LEFT) ||
            (char == '^' && cur_dir != UP) || (char == 'v' && cur_dir != DOWN) {
                valid_pos = false
            } else if slices.Contains(position_stack, target_pos) {
                valid_pos = false
            }
        }

        cur_dir += 1
        direction_stack[len(direction_stack)-1] = cur_dir

        if valid_pos {
            // Set position to be visited next (depth first search)
            position_stack = append(position_stack, target_pos)
            direction_stack = append(direction_stack, UP)
        }
    }

    longest := 0
    for _, dist := range end_distances {
        if dist > longest {
            longest = dist
        }
    }

    fmt.Printf("Part 1: %v\n", longest)
    //fmt.Printf("Part 2: %v\n", p2_score)
}
