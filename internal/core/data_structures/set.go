package data_structures

import "sync"

// Set representa un conjunto de elementos únicos, seguro para concurrencia.
type Set struct {
    data map[interface{}]struct{}
    lock sync.RWMutex
}

// NewSet crea una nueva instancia de Set.
func NewSet() *Set {
    return &Set{
        data: make(map[interface{}]struct{}),
    }
}

// Add añade un elemento al conjunto.
func (s *Set) Add(item interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data[item] = struct{}{}
}

// Remove elimina un elemento del conjunto.
func (s *Set) Remove(item interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()
    delete(s.data, item)
}

// Contains verifica si un elemento está en el conjunto.
func (s *Set) Contains(item interface{}) bool {
    s.lock.RLock()
    defer s.lock.RUnlock()
    _, exists := s.data[item]
    return exists
}

// Len devuelve la cantidad de elementos en el conjunto.
func (s *Set) Len() int {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return len(s.data)
}

// Clear elimina todos los elementos del conjunto.
func (s *Set) Clear() {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data = make(map[interface{}]struct{})
}

// Items devuelve una lista de todos los elementos en el conjunto.
func (s *Set) Items() []interface{} {
    s.lock.RLock()
    defer s.lock.RUnlock()
    items := make([]interface{}, 0, len(s.data))
    for item := range s.data {
        items = append(items, item)
    }
    return items
}
