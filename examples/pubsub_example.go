package main

import (
    "fmt"
    "github.com/devmangel/Yasmine/internal/core/pubsub"
)

func main() {
    // Configurar el gestor de Pub/Sub
    pubSubManager := pubsub.NewPubSubManager()

    // Crear un nuevo canal
    channel := pubSubManager.CreateChannel("example_channel")

    // Crear un suscriptor con un callback
    subscriber := pubsub.NewSubscriber("sub1", func(message interface{}) {
        fmt.Printf("Received message: %v\n", message)
    })

    // Suscribir al suscriptor al canal
    pubSubManager.Subscribe("example_channel", subscriber)

    // Publicar un mensaje en el canal
    pubSubManager.Publish("example_channel", "Hello, Pub/Sub!")

    // Output: Received message: Hello, Pub/Sub!
}
