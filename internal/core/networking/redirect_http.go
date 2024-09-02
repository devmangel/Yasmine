package networking

import (
    "log"
    "net/http"
)

// RedirectToHTTPS redirige todas las solicitudes HTTP a HTTPS.
func RedirectToHTTPS() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
    })

    log.Println("Starting HTTP to HTTPS redirect server on :80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatalf("HTTP redirect server failed: %v", err)
    }
}
