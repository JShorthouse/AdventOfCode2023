package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

type Card struct {
    id int
    winners []int
    numbers []int
}

func main() {
    file, err := os.Open("./input/04.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    cards := make([]Card, 0)

    idx := 0

    for scanner.Scan() {
        idx += 1
        card := Card{ id: idx }
        line := scanner.Text()

        game_split := strings.Split(line, ": ")
        number_split := strings.Split(game_split[1], "|")

        winners_split := strings.Fields(number_split[0])
        numbers_split := strings.Fields(number_split[1])

        for _, str := range winners_split {
            num, err := strconv.Atoi(str)
            if err != nil { panic(err) }

            card.winners = append(card.winners, num)
        }

        for _, str := range numbers_split {
            num, err := strconv.Atoi(str)
            if err != nil { panic(err) }

            card.numbers = append(card.numbers, num)
        }

        cards = append(cards, card)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    p1_score := 0

    for _, card := range cards {
        card_score := 0

        for _, num := range card.numbers {
            if slices.Contains(card.winners, num) {
                if card_score == 0 {
                    card_score = 1
                } else {
                    card_score *= 2
                }
            }
        }

        p1_score += card_score
    }

    card_collection := make([]Card, 0)
    card_collection = append(card_collection, cards...)

    idx = 0
    for idx < len(card_collection) {
        card := card_collection[idx]

        num_winners := 0
        for _, num := range card.numbers {
            if slices.Contains(card.winners, num) {
                num_winners += 1
            }
        }

        card_id := card.id + 1
        for num_winners > 0 && card_id-1 < len(cards) {
            card_collection = append(card_collection, cards[card_id - 1])
            card_id += 1
            num_winners -= 1
        }
        idx += 1
    }

    fmt.Printf("Part 1: %v\n", p1_score)
    fmt.Printf("Part 2: %v\n", len(card_collection))
}
