package main

import "net/http"

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request){
    cfg.siteHits = 0
    w.WriteHeader(http.StatusOK)
}
