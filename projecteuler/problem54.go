/*
https://projecteuler.net/problem=54
Poker Hands

In the card game poker, a hand consists of five cards and are ranked, from lowest to highest, in the following way:

High Card: Highest value card.
One Pair: Two cards of the same value.
Two Pairs: Two different pairs.
Three of a Kind: Three cards of the same value.
Straight: All cards are consecutive values.
Flush: All cards of the same suit.
Full House: Three of a kind and a pair.
Four of a Kind: Four cards of the same value.
Straight Flush: All cards are consecutive values of same suit.
Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
The cards are valued in the order:
2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.

If two players have the same ranked hands then the rank made up of the highest value wins; for example, a pair of eights beats a pair of fives (see example 1 below). But if two ranks tie, for example, both players have a pair of queens, then highest cards in each hand are compared (see example 4 below); if the highest cards tie then the next highest cards are compared, and so on.

The file, poker.txt (https://projecteuler.net/project/resources/p054_poker.txt), contains one-thousand random hands dealt to two players. Each line of the file contains ten cards (separated by a single space): the first five are Player 1's cards and the last five are Player 2's cards. You can assume that all hands are valid (no invalid characters or repeated cards), each player's hand is in no specific order, and in each hand there is a clear winner.

How many hands does Player 1 win?
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("problem 54")
	rounds := getRounds("./p054_poker.txt")
	p1WinCount := 0
	for _, round := range rounds {
		h1 := round.player1
		s1 := h1.score()
		h2 := round.player2
		s2 := h2.score()
		if s1 > s2 {
			p1WinCount++
		}
	}
	fmt.Println("Player 1 win count: ", p1WinCount)
}

type round struct {
	number  int
	player1 hand
	player2 hand
	winner  int
}

type hand struct {
	cards [5]card
}

type card struct {
	rank int
	suit int
}

func (h *hand) score() int {
	//2-14 x 10^0: High Card: Highest value card.
	//2-14 X 10^1: One Pair: Two cards of the same value.
	//2-14 X 10^3: Two Pairs: Two different pairs.
	//2-14 X 10^5: Three of a Kind: Three cards of the same value.
	//2-14 X 10^7: Straight: All cards are consecutive values.
	//2-14 X 10^9: Flush: All cards of the same suit.
	//2-14 X 10^11: Full House: Three of a kind and a pair.
	//2-14 X 10^13: Four of a Kind: Four cards of the same value.
	//2-14 X 10^15: Straight Flush: All cards are consecutive values of same suit.
	//2-14 X 10^17: Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.

	//FLUSH(s)
	hasFlush := true
	tmpSuit := h.cards[0].suit
	for _, c := range h.cards {
		if tmpSuit == c.suit {
			tmpSuit = c.suit
		} else {
			hasFlush = false
			break
		}
	}

	//STRAIGHT
	sorted := []int{h.cards[0].rank, h.cards[1].rank, h.cards[2].rank, h.cards[3].rank, h.cards[4].rank}
	sort.Ints(sorted)
	topCard := sorted[4]
	tmpRank := topCard
	hasStraight := true
	for i := 3; i >= 0; i-- {
		diff := tmpRank - sorted[i]
		if diff == 1 {
			tmpRank = sorted[i]
		} else {
			hasStraight = false
			break
		}
	}

	if hasStraight && hasFlush {
		if topCard == 14 { //ACE
			fmt.Println("ROYAL FLUSH")
			return int(math.Pow10(17))
		} else {
			fmt.Println("STRAIGHT FLUSH!!")
			return int(math.Pow10(15)) * topCard
		}
	}

	//build a map of cards (rank=>count)
	m := make(map[int]int)
	for _, x := range sorted {
		if val, ok := m[x]; ok {
			m[x] = val + 1
		} else {
			m[x] = 1
		}
	}
	//fmt.Println(sorted)
	//fmt.Println(m)

	//4 of a kind
	for val, count := range m {
		//4 of a kind
		if count == 4 {
			fmt.Println("4 of a kind!!")
			return int(math.Pow10(13)) * val
		}
	}

	hasPair := false
	singlePairVal := -1
	hasTwoPair := false
	twoPairVal := -1
	hasSet := false
	setVal := -1
	for val, count := range m {
		if count == 3 {
			hasSet = true
			setVal = val
		} else if count == 2 {
			if hasPair {
				hasTwoPair = true
				twoPairVal = val
			} else {
				hasPair = true
				singlePairVal = val
			}
		}
	}

	//Full House
	if hasSet && hasPair {
		//fmt.Println("Full House!!")
		//fmt.Println(h.cards)
		return int(math.Pow10(11)) * setVal
	}

	//Flush
	if hasFlush {
		return int(math.Pow10(9)) * topCard
	}
	//Straight
	if hasStraight {
		return int(math.Pow10(7)) * topCard
	}

	//Three of a kind
	if hasSet {
		return int(math.Pow10(5)) * setVal
	}

	//Two Pair
	if hasTwoPair {
		//use the higher value of the 2 pairs as the multiplier, and add the other as a tie break
		if twoPairVal > singlePairVal {
			return (int(math.Pow10(3)) * twoPairVal) + singlePairVal
		} else {
			return (int(math.Pow10(3)) * singlePairVal) + twoPairVal
		}
	}

	//Single Pair
	if hasPair {
		return int(math.Pow10(1)) * singlePairVal
	}

	//High Card (default)
	return topCard
}

func newCards(strs []string) [5]card {
	cards := [5]card{}
	for i, s := range strs {
		cards[i] = newCard(s)
	}
	return cards
}

func newCard(s string) card {
	return card{rank: getRank(s[0]), suit: getSuit(s[1])}
}

func getRank(s byte) int {
	if s == 'A' {
		return 14
	} else if s == 'K' {
		return 13
	} else if s == 'Q' {
		return 12
	} else if s == 'J' {
		return 11
	} else if s == 'T' {
		return 10
	} else if s == '9' {
		return 9
	} else if s == '8' {
		return 8
	} else if s == '7' {
		return 7
	} else if s == '6' {
		return 6
	} else if s == '5' {
		return 5
	} else if s == '4' {
		return 4
	} else if s == '3' {
		return 3
	} else if s == '2' {
		return 2
	}

	fmt.Println("ERROR")
	return 0
}

func getSuit(s byte) int {
	if s == 'S' {
		return 4
	} else if s == 'D' {
		return 3
	} else if s == 'C' {
		return 2
	} else if s == 'H' {
		return 1
	}

	fmt.Println("ERROR")
	return 0
}

func getRounds(path string) []round {
	rounds := make([]round, 0)
	lines, _ := getLines(path)
	for num, v := range lines {
		cards := strings.Split(v, " ")
		player1 := hand{cards: newCards(cards[0:5])}
		player2 := hand{cards: newCards(cards[5:10])}
		round := round{number: num, player1: player1, player2: player2}
		rounds = append(rounds, round)
	}
	return rounds
}

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
