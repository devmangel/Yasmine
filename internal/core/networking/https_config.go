package networking

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
)

// HTTPSConfig almacena la configuración SSL/TLS.
type HTTPSConfig struct {
    CertFile string
    KeyFile  string
    CAFile   string
}

// LoadTLSConfig carga los certificados SSL/TLS y retorna una configuración TLS.
func LoadTLSConfig(config *HTTPSConfig) (*tls.Config, error) {
    cert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
    if err != nil {
        return nil, err
    }

    caCert, err := ioutil.ReadFile(config.CAFile)
    if err != nil {
        return nil, err
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientCAs:    caCertPool,
        ClientAuth:   tls.RequireAndVerifyClientCert,
    }

    return tlsConfig, nil
}
