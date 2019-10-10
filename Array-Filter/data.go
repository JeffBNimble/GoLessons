package main

import (
	"encoding/json"
	"io"
	"os"
)

type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}

func (c champion) hasAllClasses(classes ...string) bool {
	classCount := len(classes)
	foundCount := 0
	for _, class := range classes {
		if found := c.hasClass(class); found {
			foundCount++
		}
	}

	return foundCount == classCount
}

func (c champion) hasSomeClasses(classes ...string) bool {
	for _, class := range classes {
		if found := c.hasClass(class); found {
			return true
		}
	}

	return false
}

func (c champion) hasClass(class string) bool {
	for _, champClass := range c.Classes {
		if champClass == class {
			return true
		}
	}

	return false
}

func (c champion) hasAllOrigins(origins ...string) bool {
	originCount := len(origins)
	foundCount := 0
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			foundCount++
		}
	}

	return foundCount == originCount
}

func (c champion) hasSomeOrigins(origins ...string) bool {
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			return true
		}
	}

	return false
}

func (c champion) hasOrigin(origin string) bool {
	for _, champOrigin := range c.Origins {
		if champOrigin == origin {
			return true
		}
	}

	return false
}

func loadChampions() ([]champion, error) {
	file, err := os.Open("tft_champions.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var champions []champion
	for {
		if err := json.NewDecoder(file).Decode(&champions); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return champions, nil
}
