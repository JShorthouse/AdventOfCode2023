package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
    "sort"
)

type Vector3 struct {
    x int
    y int
    z int
}

type Brick struct {
    min Vector3
    max Vector3
}

func parseCoordinates(str string) Vector3 {
    split := strings.Split(str, ",")
    x, x_err := strconv.Atoi(split[0])
    y, y_err := strconv.Atoi(split[1])
    z, z_err := strconv.Atoi(split[2])
    if x_err != nil || y_err != nil || z_err != nil {
        panic("parseCoordinates error")
    }
    return Vector3{ x, y, z }
}

func checkIntersection(first Brick, second Brick) bool {
    return first.min.z <= second.max.z && first.max.z >= second.min.z &&
           first.min.y <= second.max.y && first.max.y >= second.min.y &&
           first.min.x <= second.max.x && first.max.x >= second.min.x
}

var physics_called = 0

// Return count of bricks moved, and new list of brick positions
func runPhysicsStep(bricks []Brick) (int, []Brick) {
    new_positions := make([]Brick, len(bricks))
    copy(new_positions, bricks)

    move_count := 0
    for idx, brick := range new_positions {
        last_valid_pos := brick

        fallen_brick := brick
        fallen_valid := true

        for fallen_valid {
            fallen_brick.min.z -= 1
            fallen_brick.max.z -= 1

            if fallen_brick.min.z < 1 { // Floor check
                fallen_valid = false
            } else {
                for _, other_brick := range new_positions {
                    if brick == other_brick {
                        continue
                    }
                    if checkIntersection(fallen_brick, other_brick) {
                        fallen_valid = false
                        break
                    }
                }
            }
            if fallen_valid {
                last_valid_pos = fallen_brick
            }
        }

        new_positions[idx] = last_valid_pos
        if last_valid_pos != brick {
            move_count += 1
        }
    }
    
    return move_count, new_positions
}

func main() {
    file, err := os.Open("./input/22.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    bricks := make([]Brick, 0)

    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, "~")

        c1 := parseCoordinates(split[0])
        c2 := parseCoordinates(split[1])

        min := Vector3{ x: min(c1.x, c2.x), y: min(c1.y, c2.y), z: min(c1.z, c2.z) }
        max := Vector3{ x: max(c1.x, c2.x), y: max(c1.y, c2.y), z: max(c1.z, c2.z) }

        bricks = append(bricks, Brick{ min, max })
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // Let all bricks settle
    still_falling := 1
    for still_falling > 0 {
        still_falling, bricks = runPhysicsStep(bricks)
    }

    // Sort list into vertical position order
    sort.Slice(bricks, func(i, j int) bool {
        return bricks[i].max.z < bricks[j].max.z
    })

    valid_removals := 0
    total_moved := 0

    // Test removing every brick
    for _, brick := range bricks {
        test_bricks := make([]Brick, len(bricks))
        copy(test_bricks, bricks)

        test_bricks = slices.DeleteFunc(test_bricks, func(b Brick) bool { return b == brick })

        num_moved, _ := runPhysicsStep(test_bricks)
        if num_moved == 0 {
            valid_removals += 1
        }
        total_moved += num_moved
    }

    fmt.Printf("Part 1: %v\n", valid_removals)
    fmt.Printf("Part 2: %v\n", total_moved)
}
