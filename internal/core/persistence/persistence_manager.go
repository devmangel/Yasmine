package persistence

import (
    "github.com/devmangel/Yasmine/internal/core/data_structures"
    "sync"
)

// PersistenceManager gestiona las operaciones de persistencia.
type PersistenceManager struct {
    rdb *RDBPersistence
    aof *AOFPersistence
    lock sync.Mutex
}

// NewPersistenceManager crea una nueva instancia de PersistenceManager.
func NewPersistenceManager(rdb *RDBPersistence, aof *AOFPersistence) *PersistenceManager {
    return &PersistenceManager{
        rdb: rdb,
        aof: aof,
    }
}

// SaveRDB guarda los datos en formato RDB.
func (pm *PersistenceManager) SaveRDB(data *data_structures.HashTable) error {
    pm.lock.Lock()
    defer pm.lock.Unlock()
    return pm.rdb.Save(data)
}

// LoadRDB carga los datos desde un archivo RDB.
func (pm *PersistenceManager) LoadRDB() (*data_structures.HashTable, error) {
    pm.lock.Lock()
    defer pm.lock.Unlock()
    return pm.rdb.Load()
}

// AppendAOF registra una operación en el archivo AOF.
func (pm *PersistenceManager) AppendAOF(command string, args ...interface{}) error {
    pm.lock.Lock()
    defer pm.lock.Unlock()
    return pm.aof.Append(command, args...)
}
