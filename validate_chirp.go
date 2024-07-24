package main

import (
    "net/http"
    "encoding/json"
    "log"
    "os/exec"
    "strings"
    "fmt"
    chirpyDb "github.com/mjishu/chirpy/database"
)

var chirps []returnVals
var id int = 1

type parameters struct {
    Body string `json:"body"`
}
type returnVals struct {
    //Valid bool `json:"valid"`
    Id int `json:"id"`
    Body string `json:"body"`
}
func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
        return
    }

    bad_words := []string{"kerfuffle","sharbert","fornax"}
    //uuid := generateId();

    const maxChirpLength = 140
    if len(params.Body) > maxChirpLength {
        respondWithError(w, http.StatusBadRequest, "Chirp is too long")
        return
    }

    for _,word := range bad_words{
        split := strings.Fields(params.Body)
        for i,split_word := range split{
            if strings.ToLower(split_word) == word{
                split[i] = "****"
            }
        }
        params.Body = strings.Join(split," ")
    }

    newChirp := returnVals{
        Id: id,
        Body: params.Body,
    }

    respondWithJSON(w, http.StatusOK, newChirp)
    chirps = append(chirps,newChirp)
    id += 1
    fmt.Println("chirps made so far")
    for chirp := range chirps{
        fmt.Printf("Id: %v\nBody: %s\n",chirps[chirp].Id,chirps[chirp].Body)
    }
    chirpyDb.Update_database()
    WriteJson()
    chirpyDb.SayHello()
}

func generateId()string{
    newUUID, err := exec.Command("uuidgen").Output()
    if err != nil{
        fmt.Println(err)
        return ""
    }
    return string(newUUID)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
    if code > 499 {
        log.Printf("Responding with 5XX error: %s", msg)
    }
    type errorResponse struct {
        Error string `json:"error"`
    }
    respondWithJSON(w, code, errorResponse{
        Error: msg,
    })
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    dat, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling JSON: %s", err)
        w.WriteHeader(500)
        return
    }
    w.WriteHeader(code)
    w.Write(dat)
}

