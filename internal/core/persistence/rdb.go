package persistence

import (
    "encoding/gob"
    "github.com/devmangel/Yasmine/internal/core/data_structures"
    "os"
)

// RDBPersistence maneja la persistencia en formato RDB.
type RDBPersistence struct {
    filePath string
}

// NewRDBPersistence crea una nueva instancia de RDBPersistence.
func NewRDBPersistence(filePath string) *RDBPersistence {
    return &RDBPersistence{filePath: filePath}
}

// Save guarda los datos en un archivo RDB.
func (r *RDBPersistence) Save(data *data_structures.HashTable) error {
    file, err := os.Create(r.filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := gob.NewEncoder(file)
    return encoder.Encode(data)
}

// Load carga los datos desde un archivo RDB.
func (r *RDBPersistence) Load() (*data_structures.HashTable, error) {
    file, err := os.Open(r.filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data := data_structures.NewHashTable()
    decoder := gob.NewDecoder(file)
    if err := decoder.Decode(data); err != nil {
        return nil, err
    }

    return data, nil
}
