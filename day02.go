package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RoundColours struct {
	red   int
	green int
	blue  int
}

func main() {
	file, err := os.Open("./input/02.txt")
	if err != nil { panic(err) }

	scanner := bufio.NewScanner(file)

	all_games := make([][]RoundColours, 0)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")
		all_rounds := split[1]

		game := make([]RoundColours, 0)

		for _, round_string := range strings.Split(all_rounds, "; ") {
			var round RoundColours
			for _, cube := range strings.Split(round_string, ", ") {
				split := strings.Split(cube, " ")
				number, err := strconv.Atoi(split[0])
				if err != nil {
					panic(err)
				}
				colour := split[1]

				switch colour {
				case "red": round.red = number
				case "green": round.green = number
				case "blue": round.blue = number
				}
			}
			game = append(game, round)
		}
		all_games = append(all_games, game)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	minimum_cubes_list := make([]RoundColours, 0)

	for _, game := range all_games {
		minimum := RoundColours{
			red:   0,
			green: 0,
			blue:  0,
		}
		for _, round := range game {
			if round.red > minimum.red { minimum.red = round.red }
			if round.green > minimum.green { minimum.green = round.green }
			if round.blue > minimum.blue { minimum.blue = round.blue }
		}
		minimum_cubes_list = append(minimum_cubes_list, minimum)
	}

	var P1_LIMITS = RoundColours{
		red:   12,
		green: 13,
		blue:  14,
	}
	p1_score := 0
	p2_score := 0

	for idx, min_cubes := range minimum_cubes_list {
		if min_cubes.red <= P1_LIMITS.red && min_cubes.green <= P1_LIMITS.green && min_cubes.blue <= P1_LIMITS.blue {
			p1_score += idx + 1
		}
		p2_score += min_cubes.red * min_cubes.green * min_cubes.blue
	}

	fmt.Printf("Part 1: %v\n", p1_score)
	fmt.Printf("Part 2: %v\n", p2_score)
}
