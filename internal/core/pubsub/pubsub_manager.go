package pubsub

import (
    "sync"
)

// PubSubManager gestiona los canales y las suscripciones.
type PubSubManager struct {
    channels map[string]*Channel
    lock     sync.RWMutex
}

// NewPubSubManager crea una nueva instancia de PubSubManager.
func NewPubSubManager() *PubSubManager {
    return &PubSubManager{
        channels: make(map[string]*Channel),
    }
}

// CreateChannel crea un nuevo canal o devuelve uno existente.
func (psm *PubSubManager) CreateChannel(name string) *Channel {
    psm.lock.Lock()
    defer psm.lock.Unlock()

    if channel, exists := psm.channels[name]; exists {
        return channel
    }

    channel := NewChannel(name)
    psm.channels[name] = channel
    return channel
}

// Publish publica un mensaje en un canal.
func (psm *PubSubManager) Publish(channelName string, message interface{}) {
    psm.lock.RLock()
    defer psm.lock.RUnlock()

    if channel, exists := psm.channels[channelName]; exists {
        channel.Publish(message)
    }
}

// Subscribe añade una suscripción a un canal.
func (psm *PubSubManager) Subscribe(channelName string, subscriber *Subscriber) {
    psm.lock.RLock()
    defer psm.lock.RUnlock()

    if channel, exists := psm.channels[channelName]; exists {
        channel.Subscribe(subscriber)
    }
}
