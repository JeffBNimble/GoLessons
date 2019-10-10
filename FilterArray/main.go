package main

import (
	"fmt"
	"log"
)

func main() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	fmt.Printf("There are %d total champions: %v\n\n", len(champions), champions)

	brawlers := filter(champions, func(champ champion) bool {
		return champ.hasClass("Brawler") && champ.Cost >= 3
	})

	fmt.Printf("Found %d brawlers, %v\n\n", len(brawlers), brawlers)
	fmt.Printf("There are %d total champions: %v\n\n", len(champions), champions)
}

func filter(champs []champion, f filterFunc) []champion {
	count := 0
	for _, champion := range champs {
		if f(champion) {
			champs[count] = champion
			count++
		}
	}

	return champs[:count]
}

type filterFunc func(champion) bool
