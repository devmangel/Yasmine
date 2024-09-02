package data_structures

import (
    "sort"
    "sync"
)

// SortedSet representa un conjunto de elementos únicos que están ordenados.
type SortedSet struct {
    elements []interface{}
    comparator func(a, b interface{}) bool
    lock sync.RWMutex
}

// NewSortedSet crea una nueva instancia de SortedSet.
// El comparator es una función que define el orden de los elementos.
func NewSortedSet(comparator func(a, b interface{}) bool) *SortedSet {
    return &SortedSet{
        elements: []interface{}{},
        comparator: comparator,
    }
}

// Add añade un elemento al conjunto ordenado.
// La función inserta el elemento en la posición correcta para mantener el orden.
func (s *SortedSet) Add(item interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()

    idx := sort.Search(len(s.elements), func(i int) bool {
        return s.comparator(item, s.elements[i])
    })

    if idx < len(s.elements) && s.elements[idx] == item {
        return // El elemento ya está presente
    }

    s.elements = append(s.elements[:idx], append([]interface{}{item}, s.elements[idx:]...)...)
}

// Remove elimina un elemento del conjunto ordenado.
func (s *SortedSet) Remove(item interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()

    idx := sort.Search(len(s.elements), func(i int) bool {
        return s.elements[i] == item
    })

    if idx < len(s.elements) && s.elements[idx] == item {
        s.elements = append(s.elements[:idx], s.elements[idx+1:]...)
    }
}

// Contains verifica si un elemento está en el conjunto ordenado.
func (s *SortedSet) Contains(item interface{}) bool {
    s.lock.RLock()
    defer s.lock.RUnlock()

    idx := sort.Search(len(s.elements), func(i int) bool {
        return s.elements[i] == item
    })

    return idx < len(s.elements) && s.elements[idx] == item
}

// Len devuelve la cantidad de elementos en el conjunto ordenado.
func (s *SortedSet) Len() int {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return len(s.elements)
}

// Items devuelve todos los elementos en el conjunto ordenado.
func (s *SortedSet) Items() []interface{} {
    s.lock.RLock()
    defer s.lock.RUnlock()

    items := make([]interface{}, len(s.elements))
    copy(items, s.elements)
    return items
}
