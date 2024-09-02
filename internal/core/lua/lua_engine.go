package lua

import (
    "github.com/yuin/gopher-lua"
    "log"
)

// LuaEngine es responsable de ejecutar scripts Lua.
type LuaEngine struct {
    state *lua.LState
}

// NewLuaEngine crea una nueva instancia de LuaEngine.
func NewLuaEngine() *LuaEngine {
    return &LuaEngine{
        state: lua.NewState(),
    }
}

// ExecuteScript ejecuta un script Lua desde una cadena de texto.
func (le *LuaEngine) ExecuteScript(script string) error {
    if err := le.state.DoString(script); err != nil {
        log.Printf("Error ejecutando script Lua: %v", err)
        return err
    }
    return nil
}

// Close cierra el estado Lua.
func (le *LuaEngine) Close() {
    le.state.Close()
}
