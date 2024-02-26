package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
	// for i := 0; i < len(d); i++ {
	// 	fmt.Println(i+1, d[i])
	// }
}

func newDeck() deck {
	cardElementals := []string{"Water", "Fire", "Lightning", "Wind", "Stone"}
	cardCharacters := []string{"Phoenix", "Fox", "Koi", "Dragon", "Tiger", "Rabbit"}
	cards := deck{}

	for _, elemental := range cardElementals {
		for _, char := range cardCharacters {
			cards = append(cards, elemental+" "+char)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil || len(bs) == 0 {
		fmt.Println("Error:", err)
		return newDeck()
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
