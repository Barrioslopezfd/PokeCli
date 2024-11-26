package main

import (
	"fmt"
)

func inspect(pokemon string, caught *map[string]pokemonResponse) error {
	if poke, ok := (*caught)[pokemon]; ok{
		fmt.Printf("Name: %s \n",poke.Name)
		fmt.Printf("Height: %d\n",poke.Height)
		fmt.Printf("Weight: %d\n",poke.Weight)
		fmt.Print("Stats:\n")
		for i:= range poke.Stats {
			fmt.Printf("  -%s: %d\n",poke.Stats[i].Stat.Name, poke.Stats[i].BaseStat)
		}
		fmt.Print("Types:\n")
		for i := range poke.Types {
			fmt.Printf("  - %s: %d\n", poke.Types[i].Type.Name, poke.Stats[i].BaseStat)
		}
	} else {
		fmt.Println("Pokemon not caught")
	}
	return nil
}
