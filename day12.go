package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
    "slices"
)

type SpringField struct {
    springs []bool
    insertion_points []int
    goal_groups []int
}

type BoolSequenceGenerator struct {
    length uint
    count uint
    max_val uint
    cur_num uint
    sequence []bool
}

func newGenerator(length uint, count uint) BoolSequenceGenerator {
    return BoolSequenceGenerator{
        length: length,
        count: count,
        max_val: uint(math.Pow(2, float64(length))) - 1,
    }
}

func (self *BoolSequenceGenerator) next() bool {
    for {
        if self.cur_num > self.max_val {
            return false
        }

        self.sequence = make([]bool, 0, self.length)

        num_generated := uint(0)

        mask := uint(0x1)
        for idx := uint(0); idx < self.length; idx++ {
            value := (self.cur_num & mask) > 0
            self.sequence = append(self.sequence, value)
            mask = mask << 1
            if value { num_generated += 1 }
        }

        self.cur_num++

        if num_generated == self.count {
            return true
        }
    }
}

func countGroups(list []bool) []int {
    group_list := make([]int, 0)

    true_count := 0

    for _, elem := range list {
        if elem == false && true_count > 0 {
            group_list = append(group_list, true_count)
            true_count = 0
        }
        if elem == true {
            true_count += 1
        }
    }

    if true_count > 0 {
        group_list = append(group_list, true_count)
    }

    return group_list
}

func main() {
    file, err := os.Open("./input/12.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    fields := make([]SpringField, 0)

    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, " ")

        spring_map := split[0]
        groups := split[1]

        var field SpringField

        for idx, char := range spring_map {
            spring_present := false
            if char == '#' {
                spring_present = true
            } else if char == '?' {
                field.insertion_points = append(field.insertion_points, idx)
            }

            field.springs = append(field.springs, spring_present)
        }

        for _, str_num := range strings.Split(groups, ",") {
            num, err := strconv.Atoi(str_num)
            if err != nil {
                panic(err)
            }
            field.goal_groups = append(field.goal_groups, num)
        }

        fields = append(fields, field)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    total_valid_arrangements := 0

    for _, field := range fields {
        total_springs := 0
        for _, num := range field.goal_groups {
            total_springs += num
        }
        springs_present := 0
        for _, present := range field.springs {
            if present { springs_present += 1 }
        }
        springs_needed := total_springs - springs_present

        generator := newGenerator(uint(len(field.insertion_points)), uint(springs_needed))

        valid_arrangements := 0
        for generator.next() {
            springs_copy := make([]bool, len(field.springs))
            copy(springs_copy, field.springs)
            seq := generator.sequence

            for idx, insertion_pos := range field.insertion_points {
                if seq[idx] == true {
                    springs_copy[insertion_pos] = true
                }
            }

            new_groups := countGroups(springs_copy)
            if slices.Equal(new_groups, field.goal_groups) {
                valid_arrangements += 1
            }
        }

        total_valid_arrangements += valid_arrangements
    }

    fmt.Printf("Part 1: %v\n", total_valid_arrangements)
}
