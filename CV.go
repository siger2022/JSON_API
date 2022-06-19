package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Address struct {
	Country string `json:"Country:"`
	City    string `json:"City:"`
	Street  string `json:"Street:"`
}

type user struct {
	First     string `json:"First name:"`
	Last      string `json:"Last name:"`
	Age       int    `json:"Age:"`
	Language  string `json:"languages:"`
	Address   Address
	Education string     `json:"Education:"`
	CreatedAt time.Time  `json:"Created at:"`
	DeletedAt *time.Time `json:"deleted_at:"`
}

func main() {
	t := time.Now()
	users := []user{
		{First: "Ilari", Last: "Skywalker", Age: 43, Language: "Japanese, English, Finnish and Cantonese", Address: Address{"Finland", "Vanta", "Koskenkorvakatu 90, 22"}, Education: "Energy techology", CreatedAt: time.Now(), DeletedAt: &t},
		{First: "Sabrina", Last: "The builder", Age: 30, Language: "English, Arabic, and Latvian", Address: Address{"Latvia", "Liepaja", "Vilandes iela"}, Education: "Political science", CreatedAt: time.Now(), DeletedAt: &t},
	}

	output, err := json.MarshalIndent(users, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	// for decoding

	content, err := ioutil.ReadFile("CV.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var CV []user

	err2 := json.Unmarshal(content, &CV)
	if err2 != nil {
		fmt.Println("Error JSON UnMarshalling")
		fmt.Println(err2.Error())
	}
	for _, person := range CV {
		fmt.Println(person)
	}
}
