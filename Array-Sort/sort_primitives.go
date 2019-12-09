package main

import (
	"log"
	"sort"
)

func sortPrimitives() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	var names []string
	for _, champ := range champions {
		names = append(names, champ.Name)
	}

	names = names[:10]

	log.Printf(
		"The first %d names are %v, sorted = %t\n\n",
		len(names),
		names,
		sort.StringsAreSorted(names),
	)

	sort.Strings(names)
	log.Printf(
		"The sorted names (ascending) are %v, sorted = %t\n\n",
		names,
		sort.StringsAreSorted(names),
	)

	var gold []int
	for _, champ := range champions {
		gold = append(gold, champ.Cost)
	}

	gold = gold[:10]

	log.Printf(
		"The first %d gold costs are %v, sorted = %t\n\n",
		len(gold),
		gold,
		sort.IntsAreSorted(gold),
	)

	sort.Ints(gold)
	log.Printf(
		"The sorted gold are %v, sorted = %t\n\n",
		gold,
		sort.IntsAreSorted(gold),
	)
}
