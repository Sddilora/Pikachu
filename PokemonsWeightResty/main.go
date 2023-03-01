package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func main() {
	pokemon := os.Args[1]

	client := resty.New()

	req := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	resp, err := client.R().
		Get(req)

	if err != nil {
		fmt.Println("GET request resulted in an error!")
		os.Exit(0)
	}

	type Poke struct {
		Name    string
		Weight  int
		Sprites map[string]string
	}

	var pokemonData Poke

	json.Unmarshal(resp.Body(), &pokemonData)

	str2 := fmt.Sprintf("Pokemonumuzun ismi: %s\nAğırlığı: %d\nResminin bağlantısı: %s", pokemonData.Name, pokemonData.Weight, pokemonData.Sprites["front_default"])

	fmt.Println(str2)

}
