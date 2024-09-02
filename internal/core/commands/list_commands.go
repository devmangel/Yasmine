package commands

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// LPushCommand añade un valor al inicio de una lista.
type LPushCommand struct {
    BaseCommand
    list *data_structures.List
}

// NewLPushCommand crea una nueva instancia de LPushCommand.
func NewLPushCommand(list *data_structures.List) *LPushCommand {
    return &LPushCommand{
        BaseCommand: BaseCommand{commandName: "LPUSH"},
        list:        list,
    }
}

// Execute ejecuta el comando LPUSH.
func (cmd *LPushCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    cmd.list.Prepend(args[0])
    return "OK", nil
}

// RPushCommand añade un valor al final de una lista.
type RPushCommand struct {
    BaseCommand
    list *data_structures.List
}

// NewRPushCommand crea una nueva instancia de RPushCommand.
func NewRPushCommand(list *data_structures.List) *RPushCommand {
    return &RPushCommand{
        BaseCommand: BaseCommand{commandName: "RPUSH"},
        list:        list,
    }
}

// Execute ejecuta el comando RPUSH.
func (cmd *RPushCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    cmd.list.Append(args[0])
    return "OK", nil
}

// LPopCommand elimina y devuelve el primer valor de una lista.
type LPopCommand struct {
    BaseCommand
    list *data_structures.List
}

// NewLPopCommand crea una nueva instancia de LPopCommand.
func NewLPopCommand(list *data_structures.List) *LPopCommand {
    return &LPopCommand{
        BaseCommand: BaseCommand{commandName: "LPOP"},
        list:        list,
    }
}

// Execute ejecuta el comando LPOP.
func (cmd *LPopCommand) Execute(args ...interface{}) (interface{}, error) {
    if cmd.list.Len() == 0 {
        return nil, nil
    }

    firstNode := cmd.list.head
    cmd.list.Remove(firstNode)
    return firstNode.Value, nil
}

// RPopCommand elimina y devuelve el último valor de una lista.
type RPopCommand struct {
    BaseCommand
    list *data_structures.List
}

// NewRPopCommand crea una nueva instancia de RPopCommand.
func NewRPopCommand(list *data_structures.List) *RPopCommand {
    return &RPopCommand{
        BaseCommand: BaseCommand{commandName: "RPOP"},
        list:        list,
    }
}

// Execute ejecuta el comando RPOP.
func (cmd *RPopCommand) Execute(args ...interface{}) (interface{}, error) {
    if cmd.list.Len() == 0 {
        return nil, nil
    }

    lastNode := cmd.list.tail
    cmd.list.Remove(lastNode)
    return lastNode.Value, nil
}
