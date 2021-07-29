package main

import (
	"fmt"
	pokedex "golangpokemons/pokemons"
)

func main() {
	service := pokedex.NewPokeApi()
	service.LoadJSON()

	language := pokedex.GetLanguage()
	fmt.Println("\n" + language + ":")
	pokedex.GetMainMenu(service.Pokemons, language)
}
