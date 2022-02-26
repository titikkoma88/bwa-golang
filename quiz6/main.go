package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

//AddGame
func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main()  {
	gamer := Gamer{Name: "Fifa 2020"}

	gamer.AddGame("The Sandbox")
	gamer.AddGame("Roblox")
	
	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
