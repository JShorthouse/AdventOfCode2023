package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func generateDifferenceList(start_values []int) [][]int {
    diff_list := make([][]int, 0)
    diff_list = append(diff_list, start_values)

    cur_line := start_values
    cur_diffs := make([]int, 0)

    for {
        all_zeros := true

        for idx := 0; idx < len(cur_line) - 1; idx++ {
            diff := cur_line[idx+1] - cur_line[idx]
            cur_diffs = append(cur_diffs, diff)
            if diff != 0 {
                all_zeros = false
            }
        }
        cur_line = cur_diffs
        diff_list = append(diff_list, cur_diffs)
        cur_diffs = make([]int, 0)

        if all_zeros {
            break
        }
    }

    return diff_list
}

func extrapolateDifferenceList(diff_list [][]int) [][]int {
    ext_list := make([][]int, 0)

    // Pad with zeros on both sides
    for _, line := range diff_list {
        new_line := make([]int, 0)
        new_line = append(new_line, 0)
        new_line = append(new_line, line...)
        new_line = append(new_line, 0)

        ext_list = append(ext_list, new_line)
    }

    for idx := len(ext_list)-2; idx >= 0; idx-- {
        line_length := len(ext_list[idx])-1

        // Extrapolate end value
        first_val := ext_list[idx][line_length-1]
        diff_val := ext_list[idx+1][line_length-1]
        missing_val := first_val + diff_val
        ext_list[idx][line_length] = missing_val

        // Extrapolate start value
        first_val = ext_list[idx][1]
        diff_val = ext_list[idx+1][0]
        missing_val = first_val - diff_val
        ext_list[idx][0] = missing_val
    }

    return ext_list
}

//func printDiffList(diff_list [][]int) {
//    for _, line := range diff_list {
//        fmt.Println(line)
//    }
//}

func main() {
    file, err := os.Open("./input/09.txt")
    if err != nil { panic(err) }

    input_list := make([][]int, 0)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        values := make([]int, 0)
        for _, str_num := range strings.Fields(line) {
            num, err := strconv.Atoi(str_num)
            if err != nil { panic(err) }
            values = append(values, num)
        }
        input_list = append(input_list, values)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    p1_score := 0
    p2_score := 0

    for _, input := range input_list {
        diff_list := generateDifferenceList(input)
        ext_list := extrapolateDifferenceList(diff_list)

        p1_score += ext_list[0][len(ext_list[0])-1]
        p2_score += ext_list[0][0]
    }

    fmt.Printf("Part 1: %v\n", p1_score)
    fmt.Printf("Part 2: %v\n", p2_score)
}
