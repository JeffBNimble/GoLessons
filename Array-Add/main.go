package main

import (
	"log"
)

func main() {
	akali, ev, kat, elise, gnar := createChampions()

	// Adding to an array, you can't add, set allocated elements
	champs := [3]champion{}
	champs[0] = akali
	champs[1] = ev
	champs[2] = kat
	//champs[3] = elise

	log.Printf("There are %d champs, %v\n\n", len(champs), champs)

	assassins := []champion{}
	log.Printf("assassins length=%d, capacity=%d\n\n", len(assassins), cap(assassins))
	assassins = append(assassins, akali, ev, kat)
	log.Printf("assassins length=%d, capacity=%d, contents=%v\n\n", len(assassins), cap(assassins), assassins)

	shapeshifters := []champion{elise, gnar}
	log.Printf(
		"shapeshifters length=%d, capacity=%d, contents=%v\n\n",
		len(shapeshifters),
		cap(shapeshifters),
		shapeshifters,
	)

	allChamps := []champion{}
	allChamps = append(allChamps, assassins...)
	allChamps = append(allChamps, shapeshifters...)
	log.Printf("allChamps length=%d, capacity=%d, contents=%v\n\n", len(allChamps), cap(allChamps), allChamps)
}

func createChampions() (champion, champion, champion, champion, champion) {
	akali := champion{
		Name:    "Akali",
		Classes: []string{"Assassin"},
		Origins: []string{"Ninja"},
		Cost:    4,
	}

	ev := champion{
		Name:    "Evelynn",
		Classes: []string{"Assassin"},
		Origins: []string{"Demon"},
		Cost:    3,
	}

	kat := champion{
		Name:    "Katarina",
		Classes: []string{"Assassin"},
		Origins: []string{"Imperial"},
		Cost:    3,
	}

	elise := champion{
		Name:    "Elise",
		Classes: []string{"Shapeshifter"},
		Origins: []string{"Demon"},
		Cost:    1,
	}

	gnar := champion{
		Name:    "Gnar",
		Classes: []string{"Shapeshifter"},
		Origins: []string{"Wild", "Yordle"},
		Cost:    4,
	}

	return akali, ev, kat, elise, gnar
}
