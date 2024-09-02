package main

import (
    "log"
    "net/http"

    "github.com/devmangel/Yasmine/internal/core/cluster"
    "github.com/devmangel/Yasmine/internal/core/commands"
    "github.com/devmangel/Yasmine/internal/core/data_structures"
    "github.com/devmangel/Yasmine/internal/core/lua"
    "github.com/devmangel/Yasmine/internal/core/networking"
    "github.com/devmangel/Yasmine/internal/core/persistence"
    "github.com/devmangel/Yasmine/internal/core/pubsub"
    "github.com/devmangel/Yasmine/internal/core/replication"
    "github.com/devmangel/Yasmine/internal/core/security"
)

func main() {
    // Configuración básica de estructuras de datos
    hashTable := data_structures.NewHashTable()
    list := data_structures.NewList()
    set := data_structures.NewSet()
    sortedSet := data_structures.NewSortedSet(func(a, b interface{}) bool {
        return a.(float64) < b.(float64)
    })

    // Configuración de seguridad
    authenticator := security.NewAuthenticator()
    authorizer := security.NewAuthorizer()
    tlsManager, err := security.NewTLSManager("certs/server.crt", "certs/server.key", "certs/ca.crt")
    if err != nil {
        log.Fatalf("Error configurando TLS: %v", err)
    }

    // Configuración de persistencia
    rdbPersistence := persistence.NewRDBPersistence("data/yasmine.rdb")
    aofPersistence := persistence.NewAOFPersistence("data/yasmine.aof")
    persistenceManager := persistence.NewPersistenceManager(rdbPersistence, aofPersistence)

    // Configuración de replicación
    masterNode := replication.NewMasterNode(hashTable)
    slaveNode := replication.NewSlaveNode()
    replicationManager := replication.NewReplicationManager(masterNode, []*replication.SlaveNode{slaveNode})

    // Configuración de Pub/Sub
    pubSubManager := pubsub.NewPubSubManager()

    // Configuración del clúster
    clusterManager := cluster.NewClusterManager()
    node := cluster.NewNode("node-1", "localhost")
    clusterManager.AddNode(node)

    // Configuración de Lua scripting
    scriptManager := lua.NewScriptManager()

    // Configuración de comandos
    setStringCommand := commands.NewSetStringCommand(hashTable)
    getStringCommand := commands.NewGetStringCommand(hashTable)
    lPushCommand := commands.NewLPushCommand(list)
    rPushCommand := commands.NewRPushCommand(list)
    sAddCommand := commands.NewSAddCommand(set)
    zAddCommand := commands.NewZAddCommand(sortedSet)

    // Configuración de networking
    httpsConfig := tlsManager.GetTLSConfig()
    handler := http.NewServeMux()
    handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Yasmine!"))
    })

    go networking.RedirectToHTTPS()
    networking.StartServer(httpsConfig, handler)

    // Iniciar el servidor y otros componentes
    log.Println("Yasmine server started successfully.")
}
