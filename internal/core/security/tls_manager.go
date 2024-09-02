package security

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
)

// TLSManager gestiona la configuración y operación de TLS/SSL.
type TLSManager struct {
    config *tls.Config
}

// NewTLSManager crea una nueva instancia de TLSManager.
func NewTLSManager(certFile, keyFile, caFile string) (*TLSManager, error) {
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return nil, err
    }

    caCert, err := ioutil.ReadFile(caFile)
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

    return &TLSManager{config: tlsConfig}, nil
}

// GetTLSConfig devuelve la configuración TLS actual.
func (tm *TLSManager) GetTLSConfig() *tls.Config {
    return tm.config
}

// RotateCertificates permite la rotación de certificados TLS/SSL.
func (tm *TLSManager) RotateCertificates(certFile, keyFile, caFile string) error {
    log.Println("Rotating TLS certificates...")

    newCert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return err
    }

    caCert, err := ioutil.ReadFile(caFile)
    if err != nil {
        return err
    }

    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tm.config.Certificates = []tls.Certificate{newCert}
    tm.config.ClientCAs = caCertPool

    log.Println("TLS certificates rotated successfully.")
    return nil
}
