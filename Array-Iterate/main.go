package main

import (
	"log"
)

func main() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	for index, champ := range champions {
		log.Printf("%s is at index %d\n", champ.Name, index)
	}

	log.Println("---------------------")

	for _, champ := range champions {
		log.Printf("Champion is %s\n", champ.Name)
	}

	log.Println("---------------------")

	for index := range champions {
		log.Printf("There is a champion at index %d\n", index)
	}

	log.Println("---------------------")

	for i := 0; i < len(champions); i++ {
		log.Printf("%s is at index %d\n", champions[i].Name, i)
	}

	log.Println("---------------------")

	for i := len(champions) - 1; i >= 0; i-- {
		log.Printf("%s is at index %d\n", champions[i].Name, i)
	}
}
