package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "slices"
)

func hashString(input string) int {
    hash := 0
    for _, char := range input {
        hash += int(char)
        hash *= 17
        hash %= 256
    }
    return hash
}

type Lens struct {
    label string
    power int
}

func main() {
    file, err := os.Open("./input/15.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    scanner.Scan() 
    input := scanner.Text()

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    instructions := strings.Split(input, ",")

    p1_score := 0
    for _, instruction := range instructions {
        p1_score += hashString(instruction)
    }

    boxes := make([][]Lens, 256)

    for _, instruction := range instructions {
        var label string
        var operator string
        var num int
        if instruction[len(instruction)-1] == '-' {
            operator = "-"
            label = instruction[:len(instruction)-1]
        } else {
            operator = "="
            num = int(instruction[len(instruction)-1]) - int('0')
            label = instruction[:len(instruction)-2]
        }

        box_id := hashString(label)
        
        if operator == "-" {
            boxes[box_id] = slices.DeleteFunc(boxes[box_id], func(l Lens) bool { return l.label == label })
        }
        if operator == "=" {
            lens := Lens{ label: label, power: num }
            index := slices.IndexFunc(boxes[box_id], func(l Lens) bool { return l.label == label})
            if index != -1 {
                boxes[box_id][index] = lens
            } else {
                boxes[box_id] = append(boxes[box_id], lens)
            }
        }
    }

    p2_score := 0
    for idx, box := range boxes {
        for lens_idx, lens := range box {
            p2_score += (idx+1) * (lens_idx+1) * lens.power
        }
    }

    fmt.Printf("Part 1: %v\n", p1_score)
    fmt.Printf("Part 2: %v\n", p2_score)
}
