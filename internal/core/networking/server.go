package networking

import (
    "log"
    "net/http"
)

// StartServer inicia el servidor HTTPS con la configuración dada.
func StartServer(tlsConfig *tls.Config, handler http.Handler) {
    server := &http.Server{
        Addr:      ":443",
        Handler:   handler,
        TLSConfig: tlsConfig,
    }

    log.Println("Starting HTTPS server on :443")
    if err := server.ListenAndServeTLS("", ""); err != nil {
        log.Fatalf("HTTPS server failed: %v", err)
    }
}
