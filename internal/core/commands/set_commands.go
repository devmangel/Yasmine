package commands

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// SAddCommand añade uno o más elementos a un conjunto.
type SAddCommand struct {
    BaseCommand
    set *data_structures.Set
}

// NewSAddCommand crea una nueva instancia de SAddCommand.
func NewSAddCommand(set *data_structures.Set) *SAddCommand {
    return &SAddCommand{
        BaseCommand: BaseCommand{commandName: "SADD"},
        set:         set,
    }
}

// Execute ejecuta el comando SADD.
func (cmd *SAddCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) == 0 {
        return nil, ErrInvalidArguments
    }

    for _, item := range args {
        cmd.set.Add(item)
    }

    return "OK", nil
}

// SRemCommand elimina uno o más elementos de un conjunto.
type SRemCommand struct {
    BaseCommand
    set *data_structures.Set
}

// NewSRemCommand crea una nueva instancia de SRemCommand.
func NewSRemCommand(set *data_structures.Set) *SRemCommand {
    return &SRemCommand{
        BaseCommand: BaseCommand{commandName: "SREM"},
        set:         set,
    }
}

// Execute ejecuta el comando SREM.
func (cmd *SRemCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) == 0 {
        return nil, ErrInvalidArguments
    }

    for _, item := range args {
        cmd.set.Remove(item)
    }

    return "OK", nil
}

// SIsMemberCommand verifica si un elemento es miembro de un conjunto.
type SIsMemberCommand struct {
    BaseCommand
    set *data_structures.Set
}

// NewSIsMemberCommand crea una nueva instancia de SIsMemberCommand.
func NewSIsMemberCommand(set *data_structures.Set) *SIsMemberCommand {
    return &SIsMemberCommand{
        BaseCommand: BaseCommand{commandName: "SISMEMBER"},
        set:         set,
    }
}

// Execute ejecuta el comando SISMEMBER.
func (cmd *SIsMemberCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    item := args[0]
    return cmd.set.Contains(item), nil
}

// SMembersCommand devuelve todos los miembros de un conjunto.
type SMembersCommand struct {
    BaseCommand
    set *data_structures.Set
}

// NewSMembersCommand crea una nueva instancia de SMembersCommand.
func NewSMembersCommand(set *data_structures.Set) *SMembersCommand {
    return &SMembersCommand{
        BaseCommand: BaseCommand{commandName: "SMEMBERS"},
        set:         set,
    }
}

// Execute ejecuta el comando SMEMBERS.
func (cmd *SMembersCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 0 {
        return nil, ErrInvalidArguments
    }

    return cmd.set.Items(), nil
}
