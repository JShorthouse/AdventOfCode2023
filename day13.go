package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    VERTICAL = iota
    HORIZONTAL
)

type Mirror struct {
    idx int
    dir int
}

func scoreMirror(grid [][]bool, start_pos int, dir int) int {
    mismatched_count := 0

    if dir == VERTICAL {
        left_x := start_pos
        right_x := start_pos+1

        for left_x >= 0 && right_x < len(grid[0]) {
            for y := 0; y < len(grid); y++ {
                if grid[y][left_x] != grid[y][right_x] {
                    mismatched_count += 1
                }
            }

            left_x  -= 1
            right_x += 1
        }
    }
    if dir == HORIZONTAL {
        top_y := start_pos
        bottom_y := start_pos + 1

        for top_y >= 0 && bottom_y < len(grid) {
            for x := 0; x < len(grid[0]); x++ {
                if grid[top_y][x] != grid[bottom_y][x] {
                    mismatched_count += 1
                }
            }

            top_y -= 1
            bottom_y += 1
        }
    }

    return mismatched_count
}

func findMirrorLines(grid [][]bool) (*Mirror, *Mirror) {
    var perfect_mirror *Mirror = nil
    var smudged_mirror *Mirror = nil

    for y := 0; y < len(grid)-1; y++ {
        mismatched_count := scoreMirror(grid, y, HORIZONTAL)
        if mismatched_count == 0 {
            perfect_mirror = &Mirror{ idx: y, dir: HORIZONTAL }
        } else if mismatched_count == 1 {
            smudged_mirror = &Mirror{ idx: y, dir: HORIZONTAL }
        }
    }

    for x :=0; x < len(grid[0])-1; x++ {
        mismatched_count := scoreMirror(grid, x, VERTICAL)
        if mismatched_count == 0 {
            perfect_mirror = &Mirror{ idx: x, dir: VERTICAL }
        } else if mismatched_count == 1 {
            smudged_mirror = &Mirror{ idx: x, dir: VERTICAL }
        }
    }

    return perfect_mirror, smudged_mirror
}

func main() {
	file, err := os.Open("./input/13.txt")
	if err != nil { panic(err) }

	scanner := bufio.NewScanner(file)

    grids := make([][][]bool, 0)

    cur_grid := make([][]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
        if line == "" {
            grids = append(grids, cur_grid)
            cur_grid = make([][]bool, 0)
            continue
        }

        cur_row := make([]bool, 0)
        for _, char := range line {
            cur_row = append(cur_row, char == '#')
        }
        cur_grid = append(cur_grid, cur_row)
	}
    grids = append(grids, cur_grid)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

    p1_score := 0
    p2_score := 0

    for _, grid := range grids {
        mirror, smudged := findMirrorLines(grid)

        if mirror.dir == VERTICAL {
            p1_score += mirror.idx + 1
        } else {
            p1_score += (mirror.idx + 1) * 100
        }

        if smudged.dir == VERTICAL {
            p2_score += smudged.idx + 1
        } else {
            p2_score += (smudged.idx + 1) * 100
        }

    }

	fmt.Printf("Part 1: %v\n", p1_score)
	fmt.Printf("Part 2: %v\n", p2_score)
}
