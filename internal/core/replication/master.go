package replication

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// MasterNode representa el nodo maestro en el sistema de replicación.
type MasterNode struct {
    data *data_structures.HashTable
}

// NewMasterNode crea una nueva instancia de MasterNode.
func NewMasterNode(data *data_structures.HashTable) *MasterNode {
    return &MasterNode{
        data: data,
    }
}

// GetData devuelve los datos actuales del nodo maestro.
func (m *MasterNode) GetData() *data_structures.HashTable {
    return m.data
}
