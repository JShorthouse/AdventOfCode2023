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

func abs(x int) int {
    if x >= 0 {
        return x
    } else {
        return -x
    }
}

func main() {
    file, err := os.Open("./input/11.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    grid := make([][]bool, 0)

    for scanner.Scan() {
        line := scanner.Text()

        row := make([]bool, 0)
        for _, c := range line {
            row = append(row, c == '#')
        }

        grid = append(grid, row)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    galaxies := make([]Position, 0)

    row_expansion := make([]int, 0)
    col_expansion := make([]int, 0)

    total_row_expansion := 0
    total_col_expansion := 0

    for y := range grid {
        empty := true
        for x := range grid[0] {
            if grid[y][x] {
                empty = false
                galaxies = append(galaxies, Position{ x, y })
            }
        }
        if empty {
            total_row_expansion += 1
        }
        row_expansion = append(row_expansion, total_row_expansion)
    }

    for x := range grid[0] {
        empty := true
        for y := range grid[0] {
            if grid[y][x] {
                empty = false
            }
        }
        if empty {
            total_col_expansion += 1
        }
        col_expansion = append(col_expansion, total_col_expansion)
    }

    p1_distance := 0
    p2_distance := 0

    for i := 0; i < len(galaxies)-1; i++ {
        for j := i+1; j < len(galaxies); j++ {
            first := galaxies[i]
            second := galaxies[j]
            x_expansion := abs(col_expansion[first.x] - col_expansion[second.x])
            y_expansion := abs(row_expansion[first.y] - row_expansion[second.y])
            distance := abs(first.x - second.x) + abs(first.y - second.y)

            p1_distance += distance + x_expansion + y_expansion
            p2_distance += distance + 999999 * (x_expansion + y_expansion)
        }
    }

    fmt.Printf("Part 1: %v\n", p1_distance)
    fmt.Printf("Part 2: %v\n", p2_distance)
}
