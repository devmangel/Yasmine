     # Documentación de la API de Yasmine

     ## Comandos

     ### SET
     Establece un valor en la tabla hash.
     ```go
     setCmd := commands.NewSetStringCommand(hashTable)
     result, err := setCmd.Execute("clave", "valor")
     ```

     ### GET
     Recupera un valor de la tabla hash.
     ```go
     getCmd := commands.NewGetStringCommand(hashTable)
     value, err := getCmd.Execute("clave")
     ```

     ## Persistencia

     ### Guardar en RDB
     Guarda el estado actual de Yasmine en un archivo RDB.
     ```go
     rdbPersistence := persistence.NewRDBPersistence("data/yasmine.rdb")
     rdbPersistence.Save(hashTable)
     ```

     ### Cargar desde RDB
     Carga el estado de Yasmine desde un archivo RDB.
     ```go
     loadedHT, err := rdbPersistence.Load()
     ```

     ## Replicación

     ### Sincronizar Datos
     Sincroniza los datos desde el maestro a los esclavos.
     ```go
     replicationManager.SyncToSlaves(masterHT)
     ```

     ## Pub/Sub

     ### Publicar un Mensaje
     Publica un mensaje en un canal.
     ```go
     pubSubManager.Publish("canal", "mensaje")
     ```

     ### Suscribirse a un Canal
     Suscribe un cliente a un canal.
     ```go
     pubSubManager.Subscribe("canal", subscriber)
     ```

     ## Lua Scripting

     ### Ejecutar un Script Lua
     Ejecuta un script Lua dentro de Yasmine.
     ```go
     luaEngine.ExecuteScript(`print("Hola desde Lua")`)
     ```
     ```
