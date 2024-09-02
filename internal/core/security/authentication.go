package security

import "errors"

var (
    // ErrInvalidCredentials se devuelve cuando las credenciales son inválidas.
    ErrInvalidCredentials = errors.New("invalid credentials")
)

// User representa un usuario en el sistema.
type User struct {
    Username string
    Password string // En un sistema real, esto estaría hasheado.
}

// Authenticator maneja la autenticación de usuarios.
type Authenticator struct {
    users map[string]*User
}

// NewAuthenticator crea una nueva instancia de Authenticator.
func NewAuthenticator() *Authenticator {
    return &Authenticator{
        users: make(map[string]*User),
    }
}

// AddUser añade un nuevo usuario al sistema.
func (a *Authenticator) AddUser(username, password string) {
    a.users[username] = &User{
        Username: username,
        Password: password,
    }
}

// Authenticate verifica si las credenciales son correctas.
func (a *Authenticator) Authenticate(username, password string) (*User, error) {
    user, exists := a.users[username]
    if !exists || user.Password != password {
        return nil, ErrInvalidCredentials
    }
    return user, nil
}
