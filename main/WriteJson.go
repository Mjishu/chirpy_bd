package main

import (
    "os"
    //"encoding/json"
    "fmt"
)

func WriteJson(){ //this doesnt even need to read the file here? it should just add to the file?

    file,err := os.ReadFile("data/database.json")
    if err !=nil{
        panic(err)
    }
    
    fmt.Println("file is opened and ready",file)
}
