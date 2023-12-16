package main

import (
    "bufio"
    "fmt"
    "os"
)

type Vector struct {
    x int
    y int
}

type BeamHead struct {
    pos Vector
    dir int
}

const (
    UP = iota
    DOWN
    LEFT
    RIGHT
)

var DIR_OFFSETS = [4]Vector {
    Vector{ x:  0, y: -1 }, // UP
    Vector{ x:  0, y:  1 }, // DOWN
    Vector{ x: -1, y:  0 }, // LEFT
    Vector{ x:  1, y:  0 }, // RIGHT
}

func simulateBeams(grid [][]rune, start_beam BeamHead) int {
    beam_heads := make([]BeamHead, 1)
    beam_heads[0] = start_beam

    new_beams := make([]BeamHead, 0)

    visited := make([][][4]bool, 0)
    for _, grid_row := range grid {
        row := make([][4]bool, len(grid_row))
        visited = append(visited, row)
    }

    for len(beam_heads) > 0 {
        for _, beam := range beam_heads {
            x := beam.pos.x
            y := beam.pos.y

            if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
                continue
            }
            if visited[y][x][beam.dir] {
                continue
            }

            visited[y][x][beam.dir] = true

            if grid[y][x] == '.' {
                offset := DIR_OFFSETS[beam.dir]
                beam.pos.x = x + offset.x
                beam.pos.y = y + offset.y
                new_beams = append(new_beams, beam)
            } else if grid[y][x] == '\\' {
                switch beam.dir {
                    case UP: 
                        beam.dir = LEFT
                        beam.pos.x = x-1
                    case DOWN:
                        beam.dir = RIGHT
                        beam.pos.x = x+1
                    case LEFT:
                        beam.dir = UP
                        beam.pos.y = y-1
                    case RIGHT:
                        beam.dir = DOWN
                        beam.pos.y = y+1
                }
                new_beams = append(new_beams, beam)

            } else if grid[y][x] == '/' {
                switch beam.dir {
                    case UP:
                        beam.dir = RIGHT
                        beam.pos.x = x+1
                    case DOWN:
                        beam.dir = LEFT
                        beam.pos.x = x-1
                    case LEFT:
                        beam.dir = DOWN
                        beam.pos.y = y+1
                    case RIGHT:
                        beam.dir = UP
                        beam.pos.y = y-1
                }
                new_beams = append(new_beams, beam)

            } else if grid[y][x] == '|' {
                if beam.dir == UP || beam.dir == DOWN {
                    offset := DIR_OFFSETS[beam.dir]
                    beam.pos.y = y + offset.y
                    new_beams = append(new_beams, beam)
                } else { // LEFT or RIGHT
                    up_beam   := BeamHead{ pos: Vector{ x: x, y: y-1 }, dir: UP }
                    down_beam := BeamHead{ pos: Vector{ x: x, y: y+1 }, dir: DOWN }
                    new_beams = append(new_beams, up_beam, down_beam)
                }

            } else if grid[y][x] == '-' {
                if beam.dir == LEFT || beam.dir == RIGHT {
                    offset := DIR_OFFSETS[beam.dir]
                    beam.pos.x = x + offset.x
                    new_beams = append(new_beams, beam)
                } else { // UP or DOWN
                    left_beam  := BeamHead{ pos: Vector{ x: x-1, y: y }, dir: LEFT }
                    right_beam := BeamHead{ pos: Vector{ x: x+1, y: y }, dir: RIGHT }
                    new_beams = append(new_beams, left_beam, right_beam)
                }
            }
        }
        beam_heads = new_beams
        new_beams = make([]BeamHead, 0)
    }

    visited_count := 0
    for _, row := range visited {
        for _, seen_dirs := range row {
            for _, seen := range seen_dirs {
                if seen {
                    visited_count += 1
                    break
                }
            }
        }
    }

    return visited_count
}

func main() {
    file, err := os.Open("./input/16.txt")
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

    p1_score := simulateBeams(grid, BeamHead{ pos: Vector{ x: 0, y: 0 }, dir: RIGHT })

    p2_max := 0
    for x := 0; x < len(grid[0]); x++ {
        score_down := simulateBeams(grid, BeamHead{ pos: Vector{ x: x, y: 0           }, dir: DOWN })
        score_up   := simulateBeams(grid, BeamHead{ pos: Vector{ x: x, y: len(grid)-1 }, dir: UP })
        if score_down > p2_max { p2_max = score_down }
        if score_up   > p2_max { p2_max = score_up }
    }
    for y := 0; y < len(grid); y++ {
        score_right := simulateBeams(grid, BeamHead{ pos: Vector{ x: 0,              y: y }, dir: RIGHT })
        score_left  := simulateBeams(grid, BeamHead{ pos: Vector{ x: len(grid[0])-1, y: y }, dir: LEFT })
        if score_right > p2_max { p2_max = score_right }
        if score_left  > p2_max { p2_max = score_left }
    }

    fmt.Printf("Part 1: %v\n", p1_score)
    fmt.Printf("Part 2: %v\n", p2_max)
}
