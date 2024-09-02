package data_structures

import "sync"

// List representa una lista enlazada doblemente segura para concurrencia.
type List struct {
    head *Node
    tail *Node
    size int
    lock sync.RWMutex
}

// Node representa un nodo en la lista enlazada.
type Node struct {
    value interface{}
    prev  *Node
    next  *Node
}

// NewList crea una nueva instancia de List.
func NewList() *List {
    return &List{}
}

// Append añade un valor al final de la lista.
func (l *List) Append(value interface{}) {
    l.lock.Lock()
    defer l.lock.Unlock()

    newNode := &Node{
        value: value,
    }

    if l.tail == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        newNode.prev = l.tail
        l.tail.next = newNode
        l.tail = newNode
    }
    l.size++
}

// Prepend añade un valor al inicio de la lista.
func (l *List) Prepend(value interface{}) {
    l.lock.Lock()
    defer l.lock.Unlock()

    newNode := &Node{
        value: value,
    }

    if l.head == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        newNode.next = l.head
        l.head.prev = newNode
        l.head = newNode
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
        l.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        l.tail = node.prev
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

    current := l.head
    for current != nil {
        if current.value == value {
            return current
        }
        current = current.next
    }
    return nil
}
