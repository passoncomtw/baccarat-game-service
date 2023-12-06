package main

import (
	"fmt"
	"math/rand"
	"time"
)

var m_CardNum = 416
var m_CardSize = 52

var m_AllCards []int
var m_CopyCards []int
var m_First_Card int
var m_RedCardPos int
var m_PlayerCard []int
var m_BankerCard []int
var m_PlayerPoint int
var m_BankerPoint int
var m_RedCardFlag bool

var m_kCardPoint = [52]int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0,
	1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0}

var m_kPoint = [19]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8}

var m_kBankThird = [10][10]bool{
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, true, true},
	{true, true, true, true, true, true, true, true, false, true},
	{false, false, true, true, true, true, true, true, false, false},
	{false, false, false, false, false, false, true, true, false, false},
	{false, false, false, false, false, false, true, true, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func GetCard() (card int) {
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(len(m_AllCards))
	m_AllCards = remove(m_AllCards, value)
	card = value % 52
	return card
}

func LoadRemainCards() {
	m_AllCards = m_CopyCards
}

func DoRoundThrowCard() {
	LoadRemainCards()
	card := GetCard()
	m_First_Card = card
	num := GetMin(card%13, 9) + 1
	for i := 0; i < num; i++ {
		GetCard()
	}
}

func BacResultSingleGame() {
	m_RedCardPos = -1
	m_PlayerCard = nil
	m_BankerCard = nil

	for i := 0; i < 2; i++ {
		m_PlayerCard = append(m_PlayerCard, GetCard())
		m_BankerCard = append(m_BankerCard, GetCard())
	}

	m_PlayerPoint = m_kCardPoint[m_PlayerCard[0]] + m_kCardPoint[m_PlayerCard[1]]
	m_PlayerPoint = m_kPoint[m_PlayerPoint]

	m_BankerPoint = m_kCardPoint[m_BankerCard[0]] + m_kCardPoint[m_BankerCard[1]]
	m_BankerPoint = m_kPoint[m_BankerPoint]

	if !(m_PlayerPoint >= 8 || m_BankerPoint >= 8) {
		if m_PlayerPoint <= 5 {
			m_PlayerCard = append(m_PlayerCard, GetCard())

			m_PlayerPoint += m_kCardPoint[m_PlayerCard[2]]
			m_PlayerPoint = m_kPoint[m_PlayerPoint]

			if m_kBankThird[m_BankerPoint][m_kCardPoint[m_PlayerCard[2]]] {
				m_BankerCard = append(m_BankerCard, GetCard())

				m_BankerPoint += m_kCardPoint[m_BankerCard[2]]
				m_BankerPoint = m_kPoint[m_BankerPoint]
			}
		} else {
			if m_BankerPoint <= 5 {
				m_BankerCard = append(m_BankerCard, GetCard())

				m_BankerPoint += m_kCardPoint[m_BankerCard[2]]
				m_BankerPoint = m_kPoint[m_BankerPoint]
			}
		}
	}
}

func main() {
	for i := 0; i < 416; i++ {
		m_CopyCards = append(m_CopyCards, i)
	}

	LoadRemainCards()
	BacResultSingleGame()

	fmt.Printf("PlayerPoint %d  %d  %d\n", m_PlayerPoint, m_PlayerCard, len(m_AllCards))
	fmt.Printf("BankerPoint %d  %d  %d\n", m_BankerPoint, m_BankerCard, len(m_AllCards))

}
