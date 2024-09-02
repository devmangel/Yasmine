# Arquitectura de Yasmine

## Visión General

Yasmine es una base de datos en memoria distribuida diseñada para ser rápida y escalable. Se basa en una arquitectura modular que permite una fácil extensión y mantenimiento.

## Componentes Principales

- **Comandos**: Implementan la lógica para manipular estructuras de datos.
- **Persistencia**: Maneja la durabilidad de los datos mediante RDB y AOF.
- **Replicación**: Asegura la alta disponibilidad de los datos replicando entre nodos.
- **Networking**: Proporciona soporte HTTP/HTTPS para comunicación segura.
- **Pub/Sub**: Implementa un sistema de mensajería en tiempo real.
- **Lua Scripting**: Permite la ejecución de scripts Lua para extender la funcionalidad.

## Flujo de Datos

1. **Comando**: Un cliente envía un comando (por ejemplo, `SET`, `GET`).
2. **Procesamiento**: El comando se procesa utilizando las estructuras de datos internas.
3. **Persistencia**: El estado se guarda en disco utilizando RDB o AOF.
4. **Replicación**: Si es un nodo maestro, los datos se replican a los esclavos.
5. **Respuesta**: El resultado se devuelve al cliente.

