package main

import (
    "fmt"
    "net/http"
)

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    template := `
    <html>
        <body>
            <h1>Welcome, Chirpy Admin</h1>
            <p>Chirpy has been visited %d times!</p>
        </body>
    </html>
    `

    responseHTML := fmt.Sprintf(template,cfg.siteHits)

    fmt.Fprint(w,responseHTML)
}
