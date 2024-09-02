package data_structures

import "sync"

// Node representa un nodo en la lista enlazada.
type Node struct {
    Value interface{}
    prev  *Node
    next  *Node
}

// List representa una lista doblemente enlazada.
type List struct {
    Head *Node  // Exportado
    Tail *Node  // Exportado
    size int
    lock sync.RWMutex
}

// NewList crea una nueva lista vacía.
func NewList() *List {
    return &List{}
}

// Append añade un valor al final de la lista.
func (l *List) Append(value interface{}) {
    l.lock.Lock()
    defer l.lock.Unlock()

    newNode := &Node{
        Value: value,
    }

    if l.Tail == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.prev = l.Tail
        l.Tail.next = newNode
        l.Tail = newNode
    }
    l.size++
}

// Prepend añade un valor al inicio de la lista.
func (l *List) Prepend(value interface{}) {
    l.lock.Lock()
    defer l.lock.Unlock()

    newNode := &Node{
        Value: value,
    }

    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.next = l.Head
        l.Head.prev = newNode
        l.Head = newNode
    }
    l.size++
}

// Remove elimina un nodo específico de la lista.
func (l *List) Remove(node *Node) {
    l.lock.Lock()
    defer l.lock.Unlock()

    if node.prev != nil {
        node.prev.next = node.next
    } else {
        l.Head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        l.Tail = node.prev
    }

    l.size--
}

// Len devuelve la cantidad de elementos en la lista.
func (l *List) Len() int {
    l.lock.RLock()
    defer l.lock.RUnlock()
    return l.size
}

// Find busca un valor en la lista y devuelve el primer nodo encontrado.
func (l *List) Find(value interface{}) *Node {
    l.lock.RLock()
    defer l.lock.RUnlock()

    current := l.Head
    for current != nil {
        if current.Value == value {
            return current
        }
        current = current.next
    }
    return nil
}
