package database

import (
    "os"
    "fmt"
    "encoding/json"
)

type ChirpData struct{
    Chirps map[string]Chirp `json:"chirps"`
}

type Chirp struct{
    ID int `json:"id"`
    Body string `json:"body"`
}

func Update_database(){
    fileContent,err := os.ReadFile("~/coding/bootDev/go/chirpy/data/database.json")
    if err!=nil{
        fmt.Println("Error reading file", err)
        return
    }
    var data ChirpData
    err = json.Unmarshal(fileContent, &data)
    if err!= nil{
        fmt.Println("error trying to unmarshal json: ",err)
        return
    }
    fmt.Printf("data is of type %T\n")
}

func SayHello(){
    fmt.Println("hello from database")
}
