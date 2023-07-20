package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)


// A Response struct to map the Entire Response
type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}

func request(site string)(*http.Response, error){
	resp, err := http.Get(site)
   	if err != nil {
      log.Fatalln(err)
	}
	return resp, err
}

func main(){
	resp, err := request("http://pokeapi.co/api/v2/pokedex/kanto/")
	_ = err
	responseData, err := ioutil.ReadAll(resp.Body)
	var responseObject Response
    json.Unmarshal(responseData, &responseObject)
	for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }
}