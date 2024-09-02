package replication

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// SlaveNode representa un nodo esclavo en el sistema de replicación.
type SlaveNode struct {
    data *data_structures.HashTable
}

// NewSlaveNode crea una nueva instancia de SlaveNode.
func NewSlaveNode() *SlaveNode {
    return &SlaveNode{
        data: data_structures.NewHashTable(),
    }
}

// Sync sincroniza los datos desde el nodo maestro al nodo esclavo.
func (s *SlaveNode) Sync(data *data_structures.HashTable) error {
    s.data = data
    return nil
}

// ReplicateCommand ejecuta un comando replicado desde el nodo maestro.
func (s *SlaveNode) ReplicateCommand(command string, args ...interface{}) error {
    // Aquí implementarías la lógica para ejecutar el comando replicado.
    // Por ejemplo, podrías tener un switch para diferentes tipos de comandos.
    switch command {
    case "SET":
        if len(args) != 2 {
            return ErrInvalidArguments
        }
        key, ok := args[0].(string)
        if !ok {
            return ErrInvalidArguments
        }
        value, ok := args[1].(string)
        if !ok {
            return ErrInvalidArguments
        }
        s.data.Set(key, value)
    // Implementarías otros comandos aquí.
    default:
        return ErrUnknownCommand
    }
    return nil
}

// GetData devuelve los datos actuales del nodo esclavo.
func (s *SlaveNode) GetData() *data_structures.HashTable {
    return s.data
}
