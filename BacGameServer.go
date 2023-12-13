package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func removeCard(cardSource []Card, index int) []Card {
	cardSource[index] = cardSource[len(cardSource)-1]
	return cardSource[:len(cardSource)-1]
}

func drawCard(cardSource []Card) (Card, []Card) {
	index := rand.Intn(len(cardSource))
	card := cardSource[index]
	nextCardSource := removeCard(cardSource, index)
	return card, nextCardSource
}

func dealBothCards(cardSource []Card, bankerResult []Card, playerResult []Card) ([]Card, []Card, []Card) {
	playerCard, afterDrawPlayerCards := drawCard(cardSource)
	bankerCard, afterDrawBankerCards := drawCard(afterDrawPlayerCards)
	nextBankerResult := append(bankerResult, bankerCard)
	nextPlayerResult := append(playerResult, playerCard)

	return afterDrawBankerCards, nextBankerResult, nextPlayerResult
}

func dealBankerCards(cardSource []Card, bankerResult []Card) ([]Card, []Card) {
	bankerCard, afterDrawBankerCards := drawCard(cardSource)
	nextBankerResult := append(bankerResult, bankerCard)

	return afterDrawBankerCards, nextBankerResult
}

func dealPlayerCards(cardSource []Card, playerResult []Card) ([]Card, []Card) {
	bankerCard, afterDrawBankerCards := drawCard(cardSource)
	nextPlayerResult := append(playerResult, bankerCard)

	return afterDrawBankerCards, nextPlayerResult
}

func initialBothCards(cardSource []Card, bankerResult []Card, playerResult []Card) ([]Card, []Card, []Card) {
	after1RoundCards, after1RoundBankerResult, after1RoundPlayerResult := dealBothCards(cardSource, bankerResult, playerResult)
	return dealBothCards(after1RoundCards, after1RoundBankerResult, after1RoundPlayerResult)
}

func finalDealPlayerCard(cardSource []Card, playerResult []Card) ([]Card, []Card) {
	playerCount := playerResult[0].count + playerResult[1].count

	if playerCount < 6 {
		return dealPlayerCards(cardSource, playerResult)
	}
	return cardSource, playerResult
}

func finalDealBankerCard(cardSource []Card, bankerResult []Card, playerResult []Card) ([]Card, []Card) {
	if len(playerResult) == 2 {
		return cardSource, bankerResult
	}

	bankerCount := bankerResult[0].count + bankerResult[1].count
	playerThirdCard := playerResult[2]

	firstCondition := []int{0, 1, 2}
	if slices.Contains(firstCondition, bankerCount) {
		if playerThirdCard.count == 8 {
			return cardSource, bankerResult
		}
		return dealBankerCards(cardSource, bankerResult)
	}

	secondCondition := []int{3, 5}
	if slices.Contains(secondCondition, bankerCount) {
		playerThirdCondition := []int{0, 1, 8, 9}
		if slices.Contains(playerThirdCondition, playerThirdCard.count) {
			return cardSource, bankerResult
		}
		return dealBankerCards(cardSource, bankerResult)
	}

	if bankerCount == 4 {
		playerThirdCondition := []int{0, 1, 2, 3, 8, 9}
		if slices.Contains(playerThirdCondition, playerThirdCard.count) {
			return cardSource, bankerResult
		}
		return dealBankerCards(cardSource, bankerResult)
	}

	if bankerCount == 6 {
		return dealBankerCards(cardSource, bankerResult)
	}

	return cardSource, bankerResult
}

type SUIT string

const (
	Spade   SUIT = "spade"
	Heart   SUIT = "heart"
	Diamond SUIT = "diamond"
	Club    SUIT = "club"
)

type Card struct {
	suit  SUIT
	value int
	count int
}

func main() {
	cardSource := []Card{
		{suit: Spade, value: 1, count: 1}, {suit: Spade, value: 2, count: 2}, {suit: Spade, value: 3, count: 3}, {suit: Spade, value: 4, count: 4}, {suit: Spade, value: 5, count: 5}, {suit: Spade, value: 6, count: 6}, {suit: Spade, value: 7, count: 7}, {suit: Spade, value: 8, count: 8}, {suit: Spade, value: 9, count: 9}, {suit: Spade, value: 10, count: 0}, {suit: Spade, value: 11, count: 0}, {suit: Spade, value: 12, count: 0}, {suit: Spade, value: 13, count: 0},
		{suit: Heart, value: 1, count: 1}, {suit: Heart, value: 2, count: 2}, {suit: Heart, value: 3, count: 3}, {suit: Heart, value: 4, count: 4}, {suit: Heart, value: 5, count: 5}, {suit: Heart, value: 6, count: 6}, {suit: Heart, value: 7, count: 7}, {suit: Heart, value: 8, count: 8}, {suit: Heart, value: 9, count: 9}, {suit: Heart, value: 10, count: 0}, {suit: Heart, value: 11, count: 0}, {suit: Heart, value: 12, count: 0}, {suit: Heart, value: 13, count: 0},
		{suit: Diamond, value: 1, count: 1}, {suit: Diamond, value: 2, count: 2}, {suit: Diamond, value: 3, count: 3}, {suit: Diamond, value: 4, count: 4}, {suit: Diamond, value: 5, count: 5}, {suit: Diamond, value: 6, count: 6}, {suit: Diamond, value: 7, count: 7}, {suit: Diamond, value: 8, count: 8}, {suit: Diamond, value: 9, count: 9}, {suit: Diamond, value: 10, count: 0}, {suit: Diamond, value: 11, count: 0}, {suit: Diamond, value: 12, count: 0}, {suit: Diamond, value: 13, count: 0},
		{suit: Club, value: 1, count: 1}, {suit: Club, value: 2, count: 2}, {suit: Club, value: 3, count: 3}, {suit: Club, value: 4, count: 4}, {suit: Club, value: 5, count: 5}, {suit: Club, value: 6, count: 6}, {suit: Club, value: 7, count: 7}, {suit: Club, value: 8, count: 8}, {suit: Club, value: 9, count: 9}, {suit: Club, value: 10, count: 0}, {suit: Club, value: 11, count: 0}, {suit: Club, value: 12, count: 0}, {suit: Club, value: 13, count: 0},

		{suit: Spade, value: 1, count: 1}, {suit: Spade, value: 2, count: 2}, {suit: Spade, value: 3, count: 3}, {suit: Spade, value: 4, count: 4}, {suit: Spade, value: 5, count: 5}, {suit: Spade, value: 6, count: 6}, {suit: Spade, value: 7, count: 7}, {suit: Spade, value: 8, count: 8}, {suit: Spade, value: 9, count: 9}, {suit: Spade, value: 10, count: 0}, {suit: Spade, value: 11, count: 0}, {suit: Spade, value: 12, count: 0}, {suit: Spade, value: 13, count: 0},
		{suit: Heart, value: 1, count: 1}, {suit: Heart, value: 2, count: 2}, {suit: Heart, value: 3, count: 3}, {suit: Heart, value: 4, count: 4}, {suit: Heart, value: 5, count: 5}, {suit: Heart, value: 6, count: 6}, {suit: Heart, value: 7, count: 7}, {suit: Heart, value: 8, count: 8}, {suit: Heart, value: 9, count: 9}, {suit: Heart, value: 10, count: 0}, {suit: Heart, value: 11, count: 0}, {suit: Heart, value: 12, count: 0}, {suit: Heart, value: 13, count: 0},
		{suit: Diamond, value: 1, count: 1}, {suit: Diamond, value: 2, count: 2}, {suit: Diamond, value: 3, count: 3}, {suit: Diamond, value: 4, count: 4}, {suit: Diamond, value: 5, count: 5}, {suit: Diamond, value: 6, count: 6}, {suit: Diamond, value: 7, count: 7}, {suit: Diamond, value: 8, count: 8}, {suit: Diamond, value: 9, count: 9}, {suit: Diamond, value: 10, count: 0}, {suit: Diamond, value: 11, count: 0}, {suit: Diamond, value: 12, count: 0}, {suit: Diamond, value: 13, count: 0},
		{suit: Club, value: 1, count: 1}, {suit: Club, value: 2, count: 2}, {suit: Club, value: 3, count: 3}, {suit: Club, value: 4, count: 4}, {suit: Club, value: 5, count: 5}, {suit: Club, value: 6, count: 6}, {suit: Club, value: 7, count: 7}, {suit: Club, value: 8, count: 8}, {suit: Club, value: 9, count: 9}, {suit: Club, value: 10, count: 0}, {suit: Club, value: 11, count: 0}, {suit: Club, value: 12, count: 0}, {suit: Club, value: 13, count: 0},

		{suit: Spade, value: 1, count: 1}, {suit: Spade, value: 2, count: 2}, {suit: Spade, value: 3, count: 3}, {suit: Spade, value: 4, count: 4}, {suit: Spade, value: 5, count: 5}, {suit: Spade, value: 6, count: 6}, {suit: Spade, value: 7, count: 7}, {suit: Spade, value: 8, count: 8}, {suit: Spade, value: 9, count: 9}, {suit: Spade, value: 10, count: 0}, {suit: Spade, value: 11, count: 0}, {suit: Spade, value: 12, count: 0}, {suit: Spade, value: 13, count: 0},
		{suit: Heart, value: 1, count: 1}, {suit: Heart, value: 2, count: 2}, {suit: Heart, value: 3, count: 3}, {suit: Heart, value: 4, count: 4}, {suit: Heart, value: 5, count: 5}, {suit: Heart, value: 6, count: 6}, {suit: Heart, value: 7, count: 7}, {suit: Heart, value: 8, count: 8}, {suit: Heart, value: 9, count: 9}, {suit: Heart, value: 10, count: 0}, {suit: Heart, value: 11, count: 0}, {suit: Heart, value: 12, count: 0}, {suit: Heart, value: 13, count: 0},
		{suit: Diamond, value: 1, count: 1}, {suit: Diamond, value: 2, count: 2}, {suit: Diamond, value: 3, count: 3}, {suit: Diamond, value: 4, count: 4}, {suit: Diamond, value: 5, count: 5}, {suit: Diamond, value: 6, count: 6}, {suit: Diamond, value: 7, count: 7}, {suit: Diamond, value: 8, count: 8}, {suit: Diamond, value: 9, count: 9}, {suit: Diamond, value: 10, count: 0}, {suit: Diamond, value: 11, count: 0}, {suit: Diamond, value: 12, count: 0}, {suit: Diamond, value: 13, count: 0},
		{suit: Club, value: 1, count: 1}, {suit: Club, value: 2, count: 2}, {suit: Club, value: 3, count: 3}, {suit: Club, value: 4, count: 4}, {suit: Club, value: 5, count: 5}, {suit: Club, value: 6, count: 6}, {suit: Club, value: 7, count: 7}, {suit: Club, value: 8, count: 8}, {suit: Club, value: 9, count: 9}, {suit: Club, value: 10, count: 0}, {suit: Club, value: 11, count: 0}, {suit: Club, value: 12, count: 0}, {suit: Club, value: 13, count: 0},

		{suit: Spade, value: 1, count: 1}, {suit: Spade, value: 2, count: 2}, {suit: Spade, value: 3, count: 3}, {suit: Spade, value: 4, count: 4}, {suit: Spade, value: 5, count: 5}, {suit: Spade, value: 6, count: 6}, {suit: Spade, value: 7, count: 7}, {suit: Spade, value: 8, count: 8}, {suit: Spade, value: 9, count: 9}, {suit: Spade, value: 10, count: 0}, {suit: Spade, value: 11, count: 0}, {suit: Spade, value: 12, count: 0}, {suit: Spade, value: 13, count: 0},
		{suit: Heart, value: 1, count: 1}, {suit: Heart, value: 2, count: 2}, {suit: Heart, value: 3, count: 3}, {suit: Heart, value: 4, count: 4}, {suit: Heart, value: 5, count: 5}, {suit: Heart, value: 6, count: 6}, {suit: Heart, value: 7, count: 7}, {suit: Heart, value: 8, count: 8}, {suit: Heart, value: 9, count: 9}, {suit: Heart, value: 10, count: 0}, {suit: Heart, value: 11, count: 0}, {suit: Heart, value: 12, count: 0}, {suit: Heart, value: 13, count: 0},
		{suit: Diamond, value: 1, count: 1}, {suit: Diamond, value: 2, count: 2}, {suit: Diamond, value: 3, count: 3}, {suit: Diamond, value: 4, count: 4}, {suit: Diamond, value: 5, count: 5}, {suit: Diamond, value: 6, count: 6}, {suit: Diamond, value: 7, count: 7}, {suit: Diamond, value: 8, count: 8}, {suit: Diamond, value: 9, count: 9}, {suit: Diamond, value: 10, count: 0}, {suit: Diamond, value: 11, count: 0}, {suit: Diamond, value: 12, count: 0}, {suit: Diamond, value: 13, count: 0},
		{suit: Club, value: 1, count: 1}, {suit: Club, value: 2, count: 2}, {suit: Club, value: 3, count: 3}, {suit: Club, value: 4, count: 4}, {suit: Club, value: 5, count: 5}, {suit: Club, value: 6, count: 6}, {suit: Club, value: 7, count: 7}, {suit: Club, value: 8, count: 8}, {suit: Club, value: 9, count: 9}, {suit: Club, value: 10, count: 0}, {suit: Club, value: 11, count: 0}, {suit: Club, value: 12, count: 0}, {suit: Club, value: 13, count: 0},
	}

	var bankerResult []Card
	var playerResult []Card

	afterInitialCardSource, afterInitialBankerResult, afterInitialPlayerResult := initialBothCards(cardSource, bankerResult, playerResult)

	afterFinalPlayerCardSource, finalPlayerResult := finalDealPlayerCard(afterInitialCardSource, afterInitialPlayerResult)
	finalSourceCard, finalBankerResult := finalDealBankerCard(afterFinalPlayerCardSource, afterInitialBankerResult, finalPlayerResult)

	fmt.Printf("len(finalSourceCard): %d\n", len(finalSourceCard))
	fmt.Printf("finalBankerResult: %v\n", finalBankerResult)
	fmt.Printf("finalPlayerResult: %v\n", finalPlayerResult)
}
