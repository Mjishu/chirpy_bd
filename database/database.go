package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type ChirpData struct {
	Chirps map[string]Chirp `json:"chirps"`
}

type Chirp struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func Update_database() {
	fileContent, err := os.ReadFile("../database.json") //this is still trying to access data/database?
	if err != nil {
		fmt.Println("breakingnews you fucked up p.1")
		panic(err)
	}
	var data ChirpData
	err = json.Unmarshal(fileContent, &data) //error bc fileContent is empty
	if err != nil {
		fmt.Println("Breaking news you fucked up")
		panic(err)
	}
	//fmt.Printf("data is of type %T\n")
}

func SayHello() {
	fmt.Println("hello from database")
}
