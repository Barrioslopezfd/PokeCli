package main

import ("fmt")

func pokedex(caught *map[string]pokemonResponse) error{
	fmt.Println("Your Pokedex: ")
	for pokemon := range *caught{
		fmt.Println("  - "+pokemon)
	}
	return nil
}
