package commands

import "errors"

// Command es una interfaz que todos los comandos deben implementar.
type Command interface {
    Execute(args ...interface{}) (interface{}, error)
    Name() string
}

// BaseCommand es una estructura que puede ser embebida en comandos específicos.
type BaseCommand struct {
    commandName string
}

// Name devuelve el nombre del comando.
func (bc *BaseCommand) Name() string {
    return bc.commandName
}

// Execute es un método que debe ser implementado por los comandos específicos.
func (bc *BaseCommand) Execute(args ...interface{}) (interface{}, error) {
    return nil, nil
}

// ErrInvalidArguments se devuelve cuando los argumentos pasados a un comando no son válidos.
var ErrInvalidArguments = errors.New("invalid arguments provided")
