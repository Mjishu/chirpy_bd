package main

import (
    "net/http"
)

func readinessHandler(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

type apiConfig struct{
    siteHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
        cfg.siteHits ++

        next.ServeHTTP(w,r)
    })
}

func main(){
    serverMux := http.NewServeMux()

    //Access all files in dir index.html is base /
    fileserver := http.FileServer(http.Dir("."))
    fileserver_strip := http.StripPrefix("/app",fileserver)

    apiCfg := &apiConfig{
        siteHits:0,
    }

    //get requests

    serverMux.Handle("/app/*", apiCfg.middlewareMetricsInc(fileserver_strip))

    serverMux.HandleFunc("GET /api/healthz",readinessHandler)

    serverMux.HandleFunc("GET /admin/metrics",apiCfg.metricsHandler)
    serverMux.HandleFunc("/api/reset", apiCfg.resetHandler)
    
    //post requests
    //serverMux.HandleFunc("POST /api/validate_chirp",handlerChirpsValidate)
    serverMux.HandleFunc("POST /api/chirps",handlerChirpsValidate)

    server := &http.Server{
        Addr:"localhost:8080",
        Handler:serverMux,
    }
    if err := server.ListenAndServe(); err!=nil{
        panic(err)
    }
} 
