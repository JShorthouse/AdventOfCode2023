package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
)

const (
    ACCEPT = iota
    REJECT
    LESS_THAN
    GREATER_THAN
    JUMP
)

type Instruction struct {
    variant int
    field rune
    value int
    jump_to string
}

type Rock struct {
    x int
    m int
    a int
    s int
}

func runInstructions(rock Rock, instructions map[string][]Instruction) bool {
    cur_ins_list := instructions["in"]
    for {
        ins_loop:
        for _, ins := range cur_ins_list {
            var rock_val int
            if ins.field == 'x' { rock_val = rock.x }
            if ins.field == 'm' { rock_val = rock.m }
            if ins.field == 'a' { rock_val = rock.a }
            if ins.field == 's' { rock_val = rock.s }

            switch ins.variant {
            case ACCEPT:
                return true
            case REJECT:
                return false
            case JUMP:
                cur_ins_list = instructions[ins.jump_to]
                break ins_loop
            case LESS_THAN:
                if rock_val < ins.value {
                    cur_ins_list = instructions[ins.jump_to]
                    break ins_loop
                }
            case GREATER_THAN:
                if rock_val > ins.value {
                    cur_ins_list = instructions[ins.jump_to]
                    break ins_loop
                }
            }
        }
    }
    return true
}

func main() {
    file, err := os.Open("./input/19.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    parsing_rules := true

    rules_regex := regexp.MustCompile(`(\w+){(.*)}`)
    rock_regex := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
    comparison_rule_regex := regexp.MustCompile(`(\w)[><](\d+):(\w+)`)

    instructions := make(map[string][]Instruction)
    instructions["A"] = []Instruction{ Instruction{ variant: ACCEPT } }
    instructions["R"] = []Instruction{ Instruction{ variant: REJECT } }

    rocks := make([]Rock, 0)

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            parsing_rules = false
            continue
        }
        if parsing_rules {
            matches := rules_regex.FindStringSubmatch(line)
            name := matches[1]
            rules := matches[2]

            ins_list := make([]Instruction, 0)
            for _, rule := range strings.Split(rules, ",") {
                var ins Instruction
                if len(rule) > 1 && (rule[1] == '<' || rule[1] == '>') {
                    rule_parts := comparison_rule_regex.FindStringSubmatch(rule)
                    ins.field = rune(rule_parts[1][0])
                    if rule[1] == '>' {
                        ins.variant = GREATER_THAN
                    } else {
                        ins.variant = LESS_THAN
                    }
                    num, err := strconv.Atoi(rule_parts[2])
                    if err != nil { panic(err) }
                    ins.value = num
                    ins.jump_to = rule_parts[3]
                } else {
                    ins.variant = JUMP
                    ins.jump_to = rule
                }
                ins_list = append(ins_list, ins)
            }
            instructions[name] = ins_list
        } else {
            matches := rock_regex.FindStringSubmatch(line)
            x, x_err := strconv.Atoi(matches[1])
            m, m_err := strconv.Atoi(matches[2])
            a, a_err := strconv.Atoi(matches[3])
            s, s_err := strconv.Atoi(matches[4])
            if x_err != nil || m_err != nil || a_err != nil || s_err != nil { panic("Rock number error") }
            
            rocks = append(rocks, Rock{ x, m, a, s })
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    p1_score := 0
    for _, rock := range rocks {
        accepted := runInstructions(rock, instructions)
        if accepted {
            p1_score += rock.x + rock.m + rock.a + rock.s
        }
    }

    fmt.Printf("Part 1: %v\n", p1_score)
}
