package replication

import (
    "sync"
    "github.com/devmangel/Yasmine/internal/core/data_structures"
)

// ReplicationManager coordina la replicación entre nodos maestro y esclavos.
type ReplicationManager struct {
    master *MasterNode
    slaves []*SlaveNode
    lock   sync.Mutex
}

// NewReplicationManager crea una nueva instancia de ReplicationManager.
func NewReplicationManager(master *MasterNode, slaves []*SlaveNode) *ReplicationManager {
    return &ReplicationManager{
        master: master,
        slaves: slaves,
    }
}

// SyncToSlaves sincroniza los datos del maestro a los esclavos.
func (rm *ReplicationManager) SyncToSlaves(data *data_structures.HashTable) error {
    rm.lock.Lock()
    defer rm.lock.Unlock()

    for _, slave := range rm.slaves {
        if err := slave.Sync(data); err != nil {
            return err
        }
    }

    return nil
}

// HandleCommand distribuye un comando a los esclavos para replicación.
func (rm *ReplicationManager) HandleCommand(command string, args ...interface{}) error {
    rm.lock.Lock()
    defer rm.lock.Unlock()

    for _, slave := range rm.slaves {
        if err := slave.ReplicateCommand(command, args...); err != nil {
            return err
        }
    }

    return nil
}
