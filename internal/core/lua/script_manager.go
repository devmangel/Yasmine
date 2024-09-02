package lua

import (
    "sync"
)

// ScriptManager gestiona la ejecución y almacenamiento de scripts Lua.
type ScriptManager struct {
    engine  *LuaEngine
    scripts map[string]string
    lock    sync.RWMutex
}

// NewScriptManager crea una nueva instancia de ScriptManager.
func NewScriptManager() *ScriptManager {
    return &ScriptManager{
        engine:  NewLuaEngine(),
        scripts: make(map[string]string),
    }
}

// AddScript añade un nuevo script al gestor.
func (sm *ScriptManager) AddScript(name, script string) {
    sm.lock.Lock()
    defer sm.lock.Unlock()
    sm.scripts[name] = script
}

// ExecuteScript ejecuta un script Lua almacenado por su nombre.
func (sm *ScriptManager) ExecuteScript(name string) error {
    sm.lock.RLock()
    script, exists := sm.scripts[name]
    sm.lock.RUnlock()

    if !exists {
        return ErrScriptNotFound
    }

    return sm.engine.ExecuteScript(script)
}

// RemoveScript elimina un script del gestor.
func (sm *ScriptManager) RemoveScript(name string) {
    sm.lock.Lock()
    defer sm.lock.Unlock()
    delete(sm.scripts, name)
}

// Close cierra el ScriptManager y libera recursos.
func (sm *ScriptManager) Close() {
    sm.engine.Close()
}

// ErrScriptNotFound es un error que se devuelve cuando un script no se encuentra.
var ErrScriptNotFound = errors.New("script not found")
