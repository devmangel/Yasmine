package main

import (
    "fmt"

    "github.com/devmangel/Yasmine/internal/core/commands"
    "github.com/devmangel/Yasmine/internal/core/data_structures"
)

func main() {
    // Crear una nueva tabla hash
    ht := data_structures.NewHashTable()

    // Crear los comandos SET y GET
    setCmd := commands.NewSetStringCommand(ht)
    getCmd := commands.NewGetStringCommand(ht)

    // Establecer un valor en la tabla hash
    setCmd.Execute("Miguel", "30")

    // Recuperar el valor de la tabla hash
    value, _ := getCmd.Execute("Miguel")
    fmt.Printf("GET Miguel: %s\n", value)

    // Mostrar el valor esperado
    // Output: GET key1: value1
}
