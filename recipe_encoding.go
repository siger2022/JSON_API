package main

import (
	"encoding/json"
	"log"
)

//https://www.mrspskitchen.net/spicy-chicken-hotpot-%E9%BA%BB%E8%BE%A3%E9%9B%9E%E7%85%B2/

type Ingredients struct {
	Id       int    //`json:"-"`
	Item     string `json:"Item"`
	Quantity string `json:"Quantity,omitempty"`
}

type Recipe struct {
	Ingredients `json:"Ingredients"`
}

func main() {
	recipe := []Recipe{
		{Ingredients: Ingredients{Id: 1, Item: "Chicken thigh", Quantity: "800g"}},
		{Ingredients: Ingredients{Id: 2, Item: "Onion", Quantity: "1"}},
		{Ingredients: Ingredients{Id: 3, Item: "Spring onion", Quantity: "a few"}},
		{Ingredients: Ingredients{Id: 4, Item: "Garlic", Quantity: "6 gloves"}},
		{Ingredients: Ingredients{Id: 5, Item: "Ginger", Quantity: "4-5 Slices"}},
		{Ingredients: Ingredients{Id: 6, Item: "Chilli", Quantity: "2-3"}},
		{Ingredients: Ingredients{Id: 7, Item: "Chilli pepper", Quantity: "little"}},
		{Ingredients: Ingredients{Id: 8, Item: "Dired chilli ", Quantity: "little"}},
		{Ingredients: Ingredients{Id: 9, Item: "spicy sauce", Quantity: "3 tbsp"}},
		{Ingredients: Ingredients{Id: 10, Item: "soy sauce", Quantity: "1 tbsp"}},
		{Ingredients: Ingredients{Id: 11, Item: "dark soy sauce ", Quantity: "1/2 tbsp"}},
		{Ingredients: Ingredients{Id: 12, Item: "sugar", Quantity: "1 tsp"}},
		{Ingredients: Ingredients{Id: 13, Item: "wine", Quantity: ""}},
	}

	output, err := json.MarshalIndent(recipe, " >> ", "*")
	if err != nil {
		log.Print("There is some error")
		return
	}
	log.Print(string(output))

}

// Now enter to save a json file: go run filename.go > filename.json
