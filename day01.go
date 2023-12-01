package main

import (
    "fmt"
    "bufio"
    "os"
)

var WORD_NAMES = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func calculate_answer(numbers [][]int) int {
    total := 0
    for _, nums := range numbers {
        total += nums[0]*10 + nums[len(nums)-1]
    }
    return total
}

func main() {
    file, err := os.Open("./input/01.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    all_nums := make([][]int, 0)
    all_nums_and_words := make([][]int, 0)

    for scanner.Scan() {
        line_nums := make([]int, 0)
        line_nums_and_words := make([]int, 0)

        line := scanner.Text()
        for idx, char := range line {
            if char >= '0' && char <= '9' {
                value := int(char - '0')
                line_nums = append(line_nums, value)
                line_nums_and_words = append(line_nums_and_words, value)
            } else {
                for word_idx, word := range WORD_NAMES {
                    if idx + len(word) - 1 < len(line) {
                        slice_str := string(line[idx : idx+len(word)])
                        if slice_str == word {
                            line_nums_and_words = append(line_nums_and_words, word_idx + 1)
                            break
                        }
                    }
                }
            }
        }
        all_nums = append(all_nums, line_nums)
        all_nums_and_words = append(all_nums_and_words, line_nums_and_words)
    }
    if err := scanner.Err(); err != nil { panic(err) }

    fmt.Printf("Part 1: %v\n", calculate_answer(all_nums))
    fmt.Printf("Part 2: %v\n", calculate_answer(all_nums_and_words))
}
