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

var NORTH = Position{ x:  0, y: -1 }
var SOUTH = Position{ x:  0, y:  1 }
var EAST  = Position{ x:  1, y:  0 }
var WEST  = Position{ x: -1, y:  0 }

type Tile rune

var PIPE_OFFSETS = make(map[Tile][]Position)
var grid = make([][]Tile, 0)

func getTile(y, x int) Tile {
    if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
        return '.'
    }
    return grid[y][x]
}

func main() {
    PIPE_OFFSETS['|'] = []Position{ NORTH, SOUTH }
    PIPE_OFFSETS['-'] = []Position{ WEST, EAST }
    PIPE_OFFSETS['L'] = []Position{ NORTH, EAST }
    PIPE_OFFSETS['J'] = []Position{ NORTH, WEST }
    PIPE_OFFSETS['7'] = []Position{ WEST, SOUTH }
    PIPE_OFFSETS['F'] = []Position{ SOUTH, EAST }
    PIPE_OFFSETS['S'] = []Position{ NORTH, EAST, SOUTH, WEST }

    file, err := os.Open("./input/10.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, []Tile(line))
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    var start_pos Position
    posloop: for y := range grid {
        for x, tile := range grid[y] {
            if tile == 'S' {
                start_pos = Position{ x, y }
                break posloop
            }
        }
    }

    visited := make([]Position, 0)

    visited = append(visited, start_pos)
    cur_pos := start_pos

    for {
        cur_tile := getTile(cur_pos.y, cur_pos.x)

        pipeloop:
        for _, offset := range PIPE_OFFSETS[cur_tile] {

            target_pos := Position{ x: cur_pos.x + offset.x, y: cur_pos.y + offset.y }
            if len(visited) > 1 {
                last_pipe := visited[len(visited)-2]
                if target_pos == last_pipe {
                    continue
                }
            }

            target_tile := getTile(target_pos.y, target_pos.x)

            if target_tile != '.' {
                // Pipe found, check for connection to current tile
                for _, pipe_end := range PIPE_OFFSETS[target_tile] {
                    if target_pos.x + pipe_end.x == cur_pos.x && 
                       target_pos.y + pipe_end.y == cur_pos.y {
                        visited = append(visited, target_pos)
                        cur_pos = target_pos
                        break pipeloop
                    }
                }
            }
        }

        if cur_pos == start_pos {
            break
        }
    }

    p1_score := len(visited) / 2
    fmt.Printf("Part 1: %v\n", p1_score)
    //fmt.Printf("Part 2: %v\n", p2_score)
}
