package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "container/list"
)

const (
    PASSTHROUGH = iota
    FLIPFLOP
    CONJUNCTION
)

type Module struct {
    variant int
    outputs []string
    active_inputs map[string]bool
    total_inputs int
    high bool
}

type Pulse struct {
    origin string
    destination string
    high bool
}

func main() {
    file, err := os.Open("./input/20.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    modules := make(map[string]Module)

    for scanner.Scan() {
        line := scanner.Text()

        split := strings.Split(line, " -> ")
        name := split[0]
        outputs := split[1]


        var stripped_name string

        if name[0] == '%' || name[0] == '&' {
            stripped_name = name[1:]
        } else {
            stripped_name = name
        }

        // If key doesn't exist, go returns default-initialized object
        // Module will either be default here, or will have some parameters set
        // from previously parsed output steps
        module := modules[ stripped_name ]

        if name[0] == '%' {
            module.variant = FLIPFLOP
        } else if name[0] == '&' {
            module.variant = CONJUNCTION
            module.high = true
        }

        module.outputs = strings.Split(outputs, ", ")
        module.active_inputs = make(map[string]bool)
        
        for _, output := range module.outputs {
            output_module := modules[output]
            output_module.total_inputs++
            modules[output] = output_module
        }

        modules[ stripped_name ] = module
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    pulse_queue := list.New()

    low_count := 0
    high_count := 0

    for count := 0; count < 1000; count++ {
        pulse_queue.PushBack( Pulse{ origin: "button", destination: "broadcaster", high: false } )

        for pulse_queue.Front() != nil {
            front := pulse_queue.Front()
            pulse_queue.Remove(front)
            pulse := front.Value.(Pulse)

            if pulse.high {
                high_count++
            } else {
                low_count++
            }

            dest_module := modules[pulse.destination]
            update_outputs := false

            if dest_module.variant == PASSTHROUGH {
                update_outputs = true
            } else if dest_module.variant == FLIPFLOP {
                if pulse.high == false {
                    dest_module.high = !dest_module.high
                    update_outputs = true
                }
            } else if dest_module.variant == CONJUNCTION {
                update_outputs = true
                dest_module.active_inputs[pulse.origin] = pulse.high
                if len(dest_module.active_inputs) == dest_module.total_inputs {
                    all_inputs_high := true
                    for _, high := range dest_module.active_inputs {
                        if !high { all_inputs_high = false }
                    }
                    dest_module.high = !all_inputs_high
                }
            }

            modules[pulse.destination] = dest_module

            if update_outputs {
                for _, output := range dest_module.outputs {
                    pulse_queue.PushBack( Pulse{ origin: pulse.destination, destination: output, high: dest_module.high } )
                }
            }

        }
    }

    p1_score := high_count * low_count
    fmt.Printf("Part 1: %v\n", p1_score)
}
