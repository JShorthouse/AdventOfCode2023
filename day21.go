package main

import (
    "bufio"
    "fmt"
    "os"
)

type Position struct {
    x int
    y int
}

var OFFSETS = [4]Position {
    Position{ x:  0, y: -1 },
    Position{ x:  0, y:  1 },
    Position{ x: -1, y:  0 },
    Position{ x:  1, y:  0 },
}

func main() {
    file, err := os.Open("./input/21.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    grid := make([][]rune, 0)

    var start_pos Position

    for scanner.Scan() {
        line := scanner.Text()

        row := make([]rune, 0)
        y := 0
        for x, char := range line {
            if char == 'S' {
                start_pos = Position{ x, y }
                row = append(row, '.')
            } else {
                row = append(row, char)
            }
            y += 1
        }
        grid = append(grid, row)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    search_positions := make(map[Position]bool)
    search_positions[start_pos] = true

    for i:=0; i<64; i++ {
        next_positions := make([]Position, 0)
        for pos := range search_positions {
            for _, offset := range OFFSETS {
                sample_x := pos.x + offset.x
                sample_y := pos.y + offset.y
                if grid[sample_y][sample_x] == '.' {
                    next_positions = append(next_positions, Position{ x: sample_x, y: sample_y })
                }
            }
        }
        clear(search_positions)
        for _, pos := range next_positions {
            search_positions[pos] = true
        }
    }

    p1_score := len(search_positions)
    fmt.Printf("Part 1: %v\n", p1_score)
    //fmt.Printf("Part 2: %v\n", p2_score)
}
