package main

import (
   "bufio"
   "fmt"
   "os"
)

func copyGrid(grid [][]rune) [][]rune {
   clone := make([][]rune, len(grid))
   for idx := range grid {
       clone[idx] = make([]rune, len(grid[idx]))
       copy(clone[idx], grid[idx])
   }
   return clone
}

func main() {
   file, err := os.Open("./input/14.txt")
   if err != nil { panic(err) }

   scanner := bufio.NewScanner(file)

   orig_grid := make([][]rune, 0)

   for scanner.Scan() {
       line := scanner.Text()
       orig_grid = append(orig_grid, []rune(line))
   }
   if err := scanner.Err(); err != nil {
       panic(err)
   }

   grid := copyGrid(orig_grid)

   // Perform physics
   for y := 1; y < len(grid); y++ {
       for x := 0; x < len(grid[0]); x++ {
           if grid[y][x] == 'O' {
               grid[y][x] = '.'
               for scan_y := y-1; scan_y >= -1; scan_y-- {
                   if scan_y == -1 || grid[scan_y][x] != '.' {
                       grid[scan_y+1][x] = 'O'
                       break
                   }
               }
           }
       }
   }

   p1_score := 0

   // Calculate score
   for row_idx, row := range grid {
       row_score := len(grid) - row_idx
       for _, char := range row {
           if char == 'O' {
               p1_score += row_score
           }
       }
   }

   fmt.Printf("Part 1: %v\n", p1_score)
   //fmt.Printf("Part 2: %v\n", p2_score)
}
