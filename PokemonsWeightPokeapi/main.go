package main

import (
	"fmt"
	"os"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
	pokemon := os.Args[1]

	p, err := pokeapi.Pokemon(pokemon)

	if err != nil {
		fmt.Println("Pokemon couldn't be found!")
		os.Exit(0)
	}

	str := fmt.Sprintf("Pokemonumuzun ismi: %s\nAğırlığı: %d\nResminin bağlantısı: %s", p.Name, p.Weight, p.Sprites.FrontDefault)
	fmt.Println(str)
}
