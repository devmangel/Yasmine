package persistence

import (
    "os"
    "fmt"
)

// AOFPersistence maneja la persistencia en formato AOF.
type AOFPersistence struct {
    filePath string
}

// NewAOFPersistence crea una nueva instancia de AOFPersistence.
func NewAOFPersistence(filePath string) *AOFPersistence {
    return &AOFPersistence{filePath: filePath}
}

// Append registra un comando en el archivo AOF.
func (a *AOFPersistence) Append(command string, args ...interface{}) error {
    file, err := os.OpenFile(a.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = fmt.Fprintf(file, "%s %v\n", command, args)
    return err
}
