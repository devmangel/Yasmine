package networking

import (
    "crypto/tls"
    "net/http"
)

// NewHTTPSClient crea un cliente HTTP con soporte para HTTPS.
func NewHTTPSClient(tlsConfig *tls.Config) *http.Client {
    return &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: tlsConfig,
        },
    }
}
d
