package main

import (
	"log"
	"sort"
)

func sortSlices() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	champs := champions[:10]

	log.Printf("The first %d champs are %v\n\n", len(champs), champs)

	// Sort champions in ascending alphabetical order by Name
	sort.Slice(champs, func(i, j int) bool {
		return champs[i].Name < champs[j].Name
	})
	log.Printf("The sorted champions (ascending by Name) are %v\n\n", champs)
}
