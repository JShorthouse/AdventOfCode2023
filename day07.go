package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "slices"
    "strconv"
    "strings"
)

type Card rune

type Hand struct {
    cards []Card
    score int
    rank int
}

var CARD_ORDER = []Card{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2', '1'}

func cardRanking(card Card) int {
    idx := slices.Index(CARD_ORDER, card)
    if idx == -1 { panic("Card not found") }
    return len(CARD_ORDER) - idx
}

func sortCards(cards []Card) {
    sort.Slice(cards, func(i, j int) bool {
        i_rank := cardRanking(cards[i])
        j_rank := cardRanking(cards[j])

        return i_rank > j_rank
    })
}

const FIVE_OF_A_KIND  int = 10
const FOUR_OF_A_KIND  int =  9
const FULL_HOUSE      int =  8
const THREE_OF_A_KIND int =  7
const TWO_PAIR        int =  6
const ONE_PAIR        int =  5
const HIGH_CARD       int =  4

func rankHand(hand *Hand) {
    s_cards := make([]Card, 0)
    s_cards = append(s_cards, hand.cards...)
    sortCards(s_cards)

    last_card := s_cards[0]
    cur_streak := 1
    unique_cards := 1
    card_counts := make([]int, 0)

    for i := 1; i < 5; i++ {
        if s_cards[i] != last_card {
            unique_cards += 1
            card_counts = append(card_counts, cur_streak)
            cur_streak = 0
            last_card = s_cards[i]
        }
        cur_streak += 1
    }

    card_counts = append(card_counts, cur_streak)

    slices.Sort(card_counts)
    slices.Reverse(card_counts)

    if unique_cards == 1 {
        hand.rank = FIVE_OF_A_KIND
    } else if unique_cards == 2 && card_counts[0] == 4 {
        hand.rank = FOUR_OF_A_KIND
    } else if unique_cards == 2 {
        hand.rank = FULL_HOUSE
    } else if card_counts[0] == 3 {
        hand.rank = THREE_OF_A_KIND
    } else if card_counts[0] == 2 && card_counts[1] == 2 {
        hand.rank = TWO_PAIR
    } else if card_counts[0] == 2 {
        hand.rank = ONE_PAIR
    } else {
        hand.rank = HIGH_CARD
    }
}

func main() {
    file, err := os.Open("./input/07.txt")
    if err != nil { panic(err) }

    scanner := bufio.NewScanner(file)

    hands := make([]Hand, 0)

    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, " ")
        cards := []Card(split[0])
        score, err := strconv.Atoi(split[1])
        if err != nil { panic(err) }

        hand := Hand{
            cards: cards,
            score: score,
        }

        rankHand(&hand)

        hands = append(hands, hand)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    sort.Slice(hands, func (i, j int) bool {
        if hands[i].rank != hands[j].rank {
            return hands[i].rank > hands[j].rank
        }

        // Same rank, compare card values
        for idx, i_card := range hands[i].cards {
            j_card := hands[j].cards[idx]
            i_card_rank := cardRanking(i_card)
            j_card_rank := cardRanking(j_card)
            if i_card_rank != j_card_rank {
                return i_card_rank > j_card_rank
            }
        }

        return true
    })

    slices.Reverse(hands)

    p1_score := 0
    for idx, hand := range hands {
        p1_score += (idx+1) * hand.score
    }

    fmt.Printf("Part 1: %v\n", p1_score)
}
