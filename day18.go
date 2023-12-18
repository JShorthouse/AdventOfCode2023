package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

const (
    UP = iota
    DOWN
    LEFT
    RIGHT
)

type Vector struct {
    x int
    y int
}

var OFFSETS = [4]Vector {
    Vector{ x:  0, y: -1 },  // UP
    Vector{ x:  0, y:  1 },  // DOWN
    Vector{ x: -1, y:  0 },  // LEFT
    Vector{ x:  1, y:  0 },  // RIGHT
}

type Instruction struct {
    dir int
    distance int
    colour string
}

// Returns (start offset, maximum bound)
func calculateBounds(instructions []Instruction) (Vector, Vector) {
    min_x, min_y := math.MaxInt32, math.MaxInt32
    max_x, max_y := 0, 0

    pos_x, pos_y := 0, 0
    for _, ins := range instructions {
        pos_x += OFFSETS[ins.dir].x * ins.distance
        pos_y += OFFSETS[ins.dir].y * ins.distance

        if pos_x > max_x { max_x = pos_x }
        if pos_x < min_x { min_x = pos_x }
        if pos_y > max_y { max_y = pos_y }
        if pos_y < min_y { min_y = pos_y }
    }

    offset := Vector{ x: -min_x, y: -min_y }
    bound := Vector{ x: max_x - min_x, y: max_y - min_y }

    return offset, bound
}

// Detect if blocks are inside using ray algorithm from wikipedia
// https://en.wikipedia.org/wiki/Point_in_polygon
// Offset the rays by 0.5 to avoid them being co-linear with horizontal edge lines,
// this is simulated by considering two blocks at once, as if the ray was between them
func fillGrid(grid [][]bool) [][]bool {
    output := make([][]bool, 0)

    for _, row := range grid {
        row_copy := make([]bool, len(row))
        copy(row_copy, row)
        output = append(output, row_copy)
    }

    for y := 0; y < len(grid)-1; y++ {
        // Start new ray
        inside := false
        for x := 0; x < len(grid[0]); x++ {
            if grid[y][x] && grid[y+1][x] {
                inside = !inside
            }
            if inside {
                output[y][x] = true
                output[y+1][x] = true
            }
        }
    }

    return output
}

func main() {
    file, err := os.Open("./input/18.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    instructions := make([]Instruction, 0)

    for scanner.Scan() {
        line := scanner.Text()

        split := strings.Fields(line)
        var dir int 
        switch split[0] {
            case "U": dir = UP
            case "D": dir = DOWN
            case "L": dir = LEFT
            case "R": dir = RIGHT
        }
        distance, err := strconv.Atoi(split[1])
        if err != nil { panic(err) }

        colour := split[2][2:8]

        ins := Instruction{ dir, distance, colour }
        instructions = append(instructions, ins)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    offset, bound := calculateBounds(instructions)

    grid := make([][]bool, 0)
    for i := 0; i < bound.y+1; i++ {
        grid = append(grid, make([]bool, bound.x+1))
    }

    cur_pos := offset
    grid[cur_pos.y][cur_pos.x] = true

    for _, ins := range instructions {
        dir_offset := OFFSETS[ins.dir]
        for n := 0; n < ins.distance; n++ {
            cur_pos.x += dir_offset.x
            cur_pos.y += dir_offset.y
            grid[cur_pos.y][cur_pos.x] = true
        }
    }

    filled_grid := fillGrid(grid)

    p1_score := 0
    for _, row := range filled_grid {
        for _, filled := range row {
            if filled { p1_score++ }
        }
    }

    fmt.Printf("Part 1: %v\n", p1_score)
}
