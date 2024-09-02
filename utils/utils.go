package utils

import "strings"

// IsEmpty verifica si una cadena está vacía.
func IsEmpty(str string) bool {
    return len(strings.TrimSpace(str)) == 0
}

// Contains verifica si un slice contiene un elemento específico.
func Contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// MaxInt devuelve el mayor de dos enteros.
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}
