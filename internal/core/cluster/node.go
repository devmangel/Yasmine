package cluster

import "time"

// Node representa un nodo en el clúster.
type Node struct {
    ID        string
    Address   string
    LastCheck time.Time
    IsAlive   bool
}

// NewNode crea una nueva instancia de Node.
func NewNode(id, address string) *Node {
    return &Node{
        ID:        id,
        Address:   address,
        LastCheck: time.Now(),
        IsAlive:   true,
    }
}

// CheckHealth actualiza el estado de salud del nodo.
func (n *Node) CheckHealth() {
    // Lógica para verificar si el nodo está activo
    n.LastCheck = time.Now()
    // Aquí se debería implementar la lógica de verificación, por ahora asumimos que está vivo.
    n.IsAlive = true
}

// MarkDead marca un nodo como inactivo.
func (n *Node) MarkDead() {
    n.IsAlive = false
}
