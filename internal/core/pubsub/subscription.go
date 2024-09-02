package pubsub

// Subscriber representa un suscriptor en el sistema Pub/Sub.
type Subscriber struct {
    ID       string
    callback func(message interface{})
}

// NewSubscriber crea una nueva instancia de Subscriber.
func NewSubscriber(id string, callback func(message interface{})) *Subscriber {
    return &Subscriber{
        ID:       id,
        callback: callback,
    }
}

// Notify envía un mensaje al suscriptor.
func (s *Subscriber) Notify(message interface{}) {
    s.callback(message)
}

