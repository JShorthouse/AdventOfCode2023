package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

var moves []rune
var move_idx = -1

func getMove() rune {
    move_idx += 1
    if move_idx >= len(moves) {
        move_idx = 0
    }
    return moves[move_idx]
}

func resetMoves() {
    move_idx = -1
}

type Node struct {
    left string
    right string
}

func main() {
    file, err := os.Open("./input/08.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    // Parse moves line
    scanner.Scan()
    moves = []rune(scanner.Text())

    // Skip empty line
    scanner.Scan()

    maze := make(map[string]Node)
    p2_nodes := make([]Node, 0)

    // Parse map lines
    re := regexp.MustCompile(`(\w+) = \((\w+)\, (\w+)\)`)
    for scanner.Scan() {
        split := re.FindStringSubmatch(scanner.Text())

        key := split[1]
        left := split[2]
        right := split[3]

        node := Node{
            left: left,
            right: right,
        }

        maze[key] = node
        if key[2] == 'A' {
            p2_nodes = append(p2_nodes, node)
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    p1_steps := 0
    cur_node := maze["AAA"]
    for {
        p1_steps += 1

        var new_label string
        switch getMove() {
        case 'L':
            new_label = cur_node.left
        case 'R':
            new_label = cur_node.right
        }
        cur_node = maze[new_label]

        if new_label == "ZZZ" {
            break
        }
    }

    resetMoves()

    //p2_steps := 0
    //for {
    //    p2_steps += 1

    //    cur_move := getMove()
    //    num_finished := 0

    //    for idx, node := range p2_nodes {
    //        var new_label string
    //        switch cur_move {
    //        case 'L':
    //            new_label = node.left
    //        case 'R':
    //            new_label = node.right
    //        }
    //        p2_nodes[idx] = maze[new_label]

    //        if new_label[2] == 'Z' {
    //            num_finished += 1
    //        }
    //    }

    //    if num_finished == len(p2_nodes) {
    //        break
    //    }
    //    if num_finished > 2 {
    //        fmt.Println(num_finished)
    //    }
    //}

    fmt.Printf("Part 1: %v\n", p1_steps)
    //fmt.Printf("Part 2: %v\n", p2_steps)
}
