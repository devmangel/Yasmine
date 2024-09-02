package cluster

import (
    "sync"
)

// ClusterManager gestiona la coordinación entre nodos en el clúster.
type ClusterManager struct {
    nodes []*Node
    lock  sync.RWMutex
}

// NewClusterManager crea una nueva instancia de ClusterManager.
func NewClusterManager() *ClusterManager {
    return &ClusterManager{
        nodes: make([]*Node, 0),
    }
}

// AddNode añade un nuevo nodo al clúster.
func (cm *ClusterManager) AddNode(node *Node) {
    cm.lock.Lock()
    defer cm.lock.Unlock()
    cm.nodes = append(cm.nodes, node)
}

// RemoveNode elimina un nodo del clúster.
func (cm *ClusterManager) RemoveNode(nodeID string) {
    cm.lock.Lock()
    defer cm.lock.Unlock()

    for i, node := range cm.nodes {
        if node.ID == nodeID {
            cm.nodes = append(cm.nodes[:i], cm.nodes[i+1:]...)
            break
        }
    }
}

// GetNode devuelve un nodo basado en su ID.
func (cm *ClusterManager) GetNode(nodeID string) *Node {
    cm.lock.RLock()
    defer cm.lock.RUnlock()

    for _, node := range cm.nodes {
        if node.ID == nodeID {
            return node
        }
    }

    return nil
}

// BalanceLoad selecciona un nodo para distribuir la carga.
func (cm *ClusterManager) BalanceLoad() *Node {
    cm.lock.RLock()
    defer cm.lock.RUnlock()

    // Lógica simple de balanceo de carga (por ejemplo, round-robin o el nodo con menos carga)
    if len(cm.nodes) == 0 {
        return nil
    }

    // Por simplicidad, seleccionamos el primer nodo.
    // Puedes implementar un algoritmo más sofisticado.
    return cm.nodes[0]
}
