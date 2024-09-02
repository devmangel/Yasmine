package security

import "errors"

var (
    // ErrUnauthorized se devuelve cuando el usuario no tiene permisos.
    ErrUnauthorized = errors.New("unauthorized access")
)

// Role representa un rol en el sistema.
type Role struct {
    Name        string
    Permissions []string
}

// Authorizer maneja la autorización de usuarios.
type Authorizer struct {
    roles map[string]*Role
}

// NewAuthorizer crea una nueva instancia de Authorizer.
func NewAuthorizer() *Authorizer {
    return &Authorizer{
        roles: make(map[string]*Role),
    }
}

// AddRole añade un nuevo rol al sistema.
func (az *Authorizer) AddRole(role *Role) {
    az.roles[role.Name] = role
}

// Authorize verifica si un usuario tiene permiso para realizar una acción.
func (az *Authorizer) Authorize(roleName, permission string) error {
    role, exists := az.roles[roleName]
    if !exists {
        return ErrUnauthorized
    }

    for _, perm := range role.Permissions {
        if perm == permission {
            return nil
        }
    }

    return ErrUnauthorized
}
