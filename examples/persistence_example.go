package main

import (
    "fmt"

    "github.com/devmangel/Yasmine/internal/core/data_structures"
    "github.com/devmangel/Yasmine/internal/core/persistence"
)

func main() {
    // Crear una nueva tabla hash y establecer un valor
    ht := data_structures.NewHashTable()
    ht.Set("key1", "value1")

    // Configurar la persistencia RDB
    rdbPersistence := persistence.NewRDBPersistence("data/yasmine_example.rdb")

    // Guardar los datos en formato RDB
    rdbPersistence.Save(ht)

    // Cargar los datos desde el archivo RDB
    loadedHT, _ := rdbPersistence.Load()

    // Mostrar el valor recuperado
    value, _ := loadedHT.Get("key1")
    fmt.Printf("Loaded key1 from RDB: %s\n", value)

    // Output: Loaded key1 from RDB: value1
}
