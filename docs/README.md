# Yasmine

[![Go Report Card](https://goreportcard.com/badge/github.com/devmangel/Yasmine)](https://goreportcard.com/report/github.com/devmangel/Yasmine)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/devmangel/Yasmine?status.svg)](https://godoc.org/github.com/devmangel/Yasmine)

Yasmine es una base de datos en memoria de alto rendimiento y altamente escalable, inspirada en Redis y escrita en Go. Diseñada para satisfacer las necesidades de aplicaciones modernas, Yasmine ofrece una amplia gama de estructuras de datos, capacidades avanzadas de replicación y persistencia, un sistema de publicación/suscripción robusto y soporte para scripting mediante Lua.

## Características Principales

- **Alto Rendimiento**: Operaciones rápidas y eficientes en memoria.
- **Estructuras de Datos Complejas**: Soporte para listas, conjuntos, conjuntos ordenados, y más.
- **Persistencia Flexible**: Guarda y carga el estado mediante RDB y AOF.
- **Replicación para Alta Disponibilidad**: Replicación entre nodos maestro y esclavo.
- **Pub/Sub**: Sistema de mensajería en tiempo real.
- **Scripting Lua**: Extiende la funcionalidad con scripts personalizados.
- **Seguridad Incorporada**: Autenticación, autorización y soporte HTTPS.

## Tabla de Contenidos

- [Instalación](#instalación)
- [Guía Rápida](#guía-rápida)
- [Documentación](#documentación)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

## Instalación

Yasmine requiere Go 1.16 o superior. Para comenzar, sigue estos pasos:

1. **Clonar el repositorio**:

   ```bash
   git clone https://github.com/devmangel/Yasmine.git
   cd Yasmine
   ```

2. **Instalar dependencias**:

   ```bash
   go mod tidy
   ```

3. **Compilar el proyecto**:

   ```bash
   go build -o yasmine ./cmd/Yasmine/
   ```

4. **Ejecutar Yasmine**:

   ```bash
   ./yasmine
   ```

## Guía Rápida

### Ejecutar un Servidor Yasmine

Para ejecutar un servidor Yasmine localmente:

```bash
go run cmd/Yasmine/main.go
```

### Uso Básico

```go
package main

import (
    "fmt"
    "github.com/devmangel/Yasmine/internal/core/commands"
    "github.com/devmangel/Yasmine/internal/core/data_structures"
)

func main() {
    // Inicializa la tabla hash
    ht := data_structures.NewHashTable()

    // Crea y ejecuta un comando SET
    setCmd := commands.NewSetStringCommand(ht)
    setCmd.Execute("clave", "valor")

    // Crea y ejecuta un comando GET
    getCmd := commands.NewGetStringCommand(ht)
    value, _ := getCmd.Execute("clave")

    fmt.Printf("Valor obtenido: %s\n", value)
}
```

### Persistencia

Yasmine soporta persistencia a través de RDB (snapshotting) y AOF (Append Only File):

```go
// Configura la persistencia RDB
rdbPersistence := persistence.NewRDBPersistence("data/yasmine.rdb")

// Guarda el estado actual
rdbPersistence.Save(hashTable)

// Carga el estado desde un archivo RDB
loadedHT, _ := rdbPersistence.Load()
```

### Replicación

Yasmine facilita la replicación de datos entre nodos:

```go
// Configura el nodo maestro
masterHT := data_structures.NewHashTable()
masterNode := replication.NewMasterNode(masterHT)

// Configura un nodo esclavo
slaveNode := replication.NewSlaveNode()

// Sincroniza datos entre el maestro y el esclavo
replicationManager := replication.NewReplicationManager(masterNode, []*replication.SlaveNode{slaveNode})
replicationManager.SyncToSlaves(masterHT)
```

## Documentación

La documentación completa del proyecto, incluyendo guías de instalación, uso avanzado, y contribución, está disponible en el directorio [docs/](docs/). También puedes consultar la [API de Yasmine](https://godoc.org/github.com/devmangel/Yasmine) para obtener detalles sobre las funciones públicas.

## Contribuciones

Las contribuciones son muy bienvenidas. Por favor, revisa nuestra [Guía para Contribuyentes](docs/CONTRIBUTING.md) antes de enviar un pull request.

### Principios para Contribuir

- **Código de Alta Calidad**: Sigue las mejores prácticas de Go, escribe pruebas para tu código, y documenta tus cambios.
- **Discusión Previa**: Abre un issue antes de empezar a trabajar en cambios importantes para alinearnos en los objetivos.
- **Pruebas Automatizadas**: Asegúrate de que todas las pruebas pasen antes de enviar tu contribución.

## Licencia

Yasmine se distribuye bajo la licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.
