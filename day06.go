package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func parseNumberFields(line string) []int {
    split := strings.Split(line, ":")
    fields := strings.Fields(split[1])

    output := make([]int, 0)

    for _, str_num := range fields {
        num, err := strconv.Atoi(str_num)
        if err != nil { panic(err) }
        output = append(output, num)
    }

    return output
}

func combineNums(nums []int) int {
    num_str := ""
    for _, num := range nums {
        num_str += strconv.Itoa(num)
    }

    combined_num, err := strconv.Atoi(num_str)
    if err != nil { panic(err) }
    return combined_num
}

func calcDistance(wait_time int, total_time int) int {
    distance := (total_time - wait_time) * wait_time
    return distance
}

func countWinningOptions(max_time int, min_distance int) int {
    winning_count := 0 
    for wait_time := 0; wait_time < max_time; wait_time += 1 {
        dist := calcDistance(wait_time, max_time)
        if dist > min_distance {
            winning_count += 1
        }
    }

    return winning_count
}

func main() {
    file, err := os.Open("./input/06.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    times := parseNumberFields(scanner.Text())

    scanner.Scan()
    min_distances := parseNumberFields(scanner.Text())

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    p1_score := 1

    for round_id := range times {
        p1_score *= countWinningOptions(times[round_id], min_distances[round_id])
    }

    p2_time := combineNums(times)
    p2_distance := combineNums(min_distances)
    p2_score := countWinningOptions(p2_time, p2_distance)

    fmt.Printf("Part 1: %v\n", p1_score)
    fmt.Printf("Part 2: %v\n", p2_score)
}
