package main

import (
    "fmt"

    "github.com/devmangel/Yasmine/internal/core/data_structures"
    "github.com/devmangel/Yasmine/internal/core/replication"
)

func main() {
    // Crear la tabla hash para el nodo maestro
    masterHT := data_structures.NewHashTable()
    masterNode := replication.NewMasterNode(masterHT)

    // Crear un nodo esclavo
    slaveNode := replication.NewSlaveNode()

    // Configurar el gestor de replicación
    replicationManager := replication.NewReplicationManager(masterNode, []*replication.SlaveNode{slaveNode})

    // Establecer un valor en el nodo maestro
    masterHT.Set("key1", "value1")

    // Sincronizar los datos del maestro al esclavo
    replicationManager.SyncToSlaves(masterHT)

    // Recuperar el valor del nodo esclavo
    value, _ := slaveNode.GetData().Get("key1")
    fmt.Printf("Replicated key1: %s\n", value)

    // Output: Replicated key1: value1
}
