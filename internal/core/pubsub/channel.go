package pubsub

import (
    "sync"
)

// Channel representa un canal de Pub/Sub.
type Channel struct {
    name        string
    subscribers []*Subscriber
    lock        sync.RWMutex
}

// NewChannel crea una nueva instancia de Channel.
func NewChannel(name string) *Channel {
    return &Channel{
        name:        name,
        subscribers: make([]*Subscriber, 0),
    }
}

// Publish publica un mensaje a todos los suscriptores del canal.
func (c *Channel) Publish(message interface{}) {
    c.lock.RLock()
    defer c.lock.RUnlock()

    for _, subscriber := range c.subscribers {
        subscriber.Notify(message)
    }
}

// Subscribe añade un suscriptor al canal.
func (c *Channel) Subscribe(subscriber *Subscriber) {
    c.lock.Lock()
    defer c.lock.Unlock()

    c.subscribers = append(c.subscribers, subscriber)
}
