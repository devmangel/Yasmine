package data_structures

import "sync"

// HashTable representa una tabla hash con concurrencia segura.
type HashTable struct {
    data map[string]interface{}
    lock sync.RWMutex
}

// NewHashTable crea una nueva instancia de HashTable.
func NewHashTable() *HashTable {
    return &HashTable{
        data: make(map[string]interface{}),
    }
}

// Set añade o actualiza un valor en la tabla hash.
func (h *HashTable) Set(key string, value interface{}) {
    h.lock.Lock()
    defer h.lock.Unlock()
    h.data[key] = value
}

// Get obtiene un valor de la tabla hash.
func (h *HashTable) Get(key string) (interface{}, bool) {
    h.lock.RLock()
    defer h.lock.RUnlock()
    value, exists := h.data[key]
    return value, exists
}

// Delete elimina un valor de la tabla hash.
func (h *HashTable) Delete(key string) {
    h.lock.Lock()
    defer h.lock.Unlock()
    delete(h.data, key)
}

// Len devuelve la cantidad de elementos en la tabla hash.
func (h *HashTable) Len() int {
    h.lock.RLock()
    defer h.lock.RUnlock()
    return len(h.data)
}
