package main

import "fmt"

func main(){
	vg := [7]string{
		"Mario Kart",
		"Yoshi Island",
		"Mario vs Donkey Kong",
		"Pokemon Diamond",
		"Pokemon Pearl",
		"Pokemon Platinum",
		"Super Smash",
	}

	vgOwned := [2]string{
		"Mario vs Donkey Kong 2",
		"Pokemon Platinum",
	}

	for _, game := range vg{
		for _, ownedGame := range vgOwned{
			if game == ownedGame{
				fmt.Printf("You own %v \n",game)
			}
		}
	}
}