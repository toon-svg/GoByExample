package main

import "fmt"

type player struct {
	name      string
	weapon    string
	health    int
	armorType string
}

func main() {
	player1 := player{
		name:      "Josh",
		weapon:    "Steel sword",
		health:    100,
		armorType: "Light",
	}
	fmt.Println(player1)
}
