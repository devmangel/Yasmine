package commands

import "github.com/devmangel/Yasmine/internal/core/data_structures"

// ZAddCommand añade un elemento a un conjunto ordenado con una puntuación.
type ZAddCommand struct {
    BaseCommand
    sortedSet *data_structures.SortedSet
}

// NewZAddCommand crea una nueva instancia de ZAddCommand.
func NewZAddCommand(sortedSet *data_structures.SortedSet) *ZAddCommand {
    return &ZAddCommand{
        BaseCommand: BaseCommand{commandName: "ZADD"},
        sortedSet:   sortedSet,
    }
}

// Execute ejecuta el comando ZADD.
func (cmd *ZAddCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 2 {
        return nil, ErrInvalidArguments
    }

    score, ok := args[0].(float64)
    if !ok {
        return nil, ErrInvalidArguments
    }

    member := args[1]
    cmd.sortedSet.Add(member, score)

    return "OK", nil
}

// ZRemCommand elimina un elemento de un conjunto ordenado.
type ZRemCommand struct {
    BaseCommand
    sortedSet *data_structures.SortedSet
}

// NewZRemCommand crea una nueva instancia de ZRemCommand.
func NewZRemCommand(sortedSet *data_structures.SortedSet) *ZRemCommand {
    return &ZRemCommand{
        BaseCommand: BaseCommand{commandName: "ZREM"},
        sortedSet:   sortedSet,
    }
}

// Execute ejecuta el comando ZREM.
func (cmd *ZRemCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    member := args[0]
    cmd.sortedSet.Remove(member)

    return "OK", nil
}

// ZRangeCommand devuelve un rango de elementos de un conjunto ordenado.
type ZRangeCommand struct {
    BaseCommand
    sortedSet *data_structures.SortedSet
}

// NewZRangeCommand crea una nueva instancia de ZRangeCommand.
func NewZRangeCommand(sortedSet *data_structures.SortedSet) *ZRangeCommand {
    return &ZRangeCommand{
        BaseCommand: BaseCommand{commandName: "ZRANGE"},
        sortedSet:   sortedSet,
    }
}

// Execute ejecuta el comando ZRANGE.
func (cmd *ZRangeCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 2 {
        return nil, ErrInvalidArguments
    }

    start, ok := args[0].(int)
    if !ok {
        return nil, ErrInvalidArguments
    }

    end, ok := args[1].(int)
    if !ok {
        return nil, ErrInvalidArguments
    }

    return cmd.sortedSet.Range(start, end), nil
}

// ZScoreCommand devuelve la puntuación de un elemento en un conjunto ordenado.
type ZScoreCommand struct {
    BaseCommand
    sortedSet *data_structures.SortedSet
}

// NewZScoreCommand crea una nueva instancia de ZScoreCommand.
func NewZScoreCommand(sortedSet *data_structures.SortedSet) *ZScoreCommand {
    return &ZScoreCommand{
        BaseCommand: BaseCommand{commandName: "ZSCORE"},
        sortedSet:   sortedSet,
    }
}

// Execute ejecuta el comando ZSCORE.
func (cmd *ZScoreCommand) Execute(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, ErrInvalidArguments
    }

    member := args[0]
    score, exists := cmd.sortedSet.Score(member)
    if !exists {
        return nil, nil
    }

    return score, nil
}
