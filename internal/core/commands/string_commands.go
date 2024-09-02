package commands

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// SetStringCommand establece un valor en una tabla hash.
type SetStringCommand struct {
    BaseCommand
    table *data_structures.HashTable
}

// NewSetStringCommand crea una nueva instancia de SetStringCommand.
func NewSetStringCommand(table *data_structures.HashTable) *SetStringCommand {
    return &SetStringCommand{
        BaseCommand: BaseCommand{commandName: "SET"},
        table:       table,
    }
}

// Execute ejecuta el comando SET.
func (cmd *SetStringCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 2 {
        return nil, ErrInvalidArguments
    }

    key, ok := args[0].(string)
    if !ok {
        return nil, ErrInvalidArguments
    }

    value, ok := args[1].(string)
    if !ok {
        return nil, ErrInvalidArguments
    }

    cmd.table.Set(key, value)
    return "OK", nil
}

// GetStringCommand obtiene un valor de una tabla hash.
type GetStringCommand struct {
    BaseCommand
    table *data_structures.HashTable
}

// NewGetStringCommand crea una nueva instancia de GetStringCommand.
func NewGetStringCommand(table *data_structures.HashTable) *GetStringCommand {
    return &GetStringCommand{
        BaseCommand: BaseCommand{commandName: "GET"},
        table:       table,
    }
}

// Execute ejecuta el comando GET.
func (cmd *GetStringCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    key, ok := args[0].(string)
    if !ok {
        return nil, ErrInvalidArguments
    }

    value, exists := cmd.table.Get(key)
    if !exists {
        return nil, nil
    }

    return value, nil
}
