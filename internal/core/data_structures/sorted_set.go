package data_structures

// SortedSet representa un conjunto ordenado donde cada elemento tiene una puntuación asociada.
type SortedSet struct {
    items      map[interface{}]float64
    comparator func(a, b interface{}) bool
}

// NewSortedSet crea una nueva instancia de SortedSet.
func NewSortedSet(comparator func(a, b interface{}) bool) *SortedSet {
    return &SortedSet{
        items:      make(map[interface{}]float64),
        comparator: comparator,
    }
}

// Add añade un nuevo elemento con su puntuación al conjunto ordenado.
func (ss *SortedSet) Add(value interface{}, score float64) {
    ss.items[value] = score
}

// Remove elimina un elemento del conjunto ordenado.
func (ss *SortedSet) Remove(value interface{}) {
    delete(ss.items, value)
}

// Range devuelve un slice de elementos en un rango específico.
func (ss *SortedSet) Range(start, end int) []interface{} {
    var result []interface{}
    
    for value := range ss.items {
        result = append(result, value)
    }

    // Ordenar los resultados usando el comparador proporcionado
    sortedResult := ss.sort(result)
    
    // Manejo de límites del rango
    if start < 0 {
        start = 0
    }
    if end > len(sortedResult) {
        end = len(sortedResult)
    }

    return sortedResult[start:end]
}

// Score devuelve la puntuación de un elemento en el conjunto ordenado.
func (ss *SortedSet) Score(value interface{}) (float64, bool) {
    score, exists := ss.items[value]
    return score, exists
}

// Método auxiliar para ordenar los resultados.
func (ss *SortedSet) sort(values []interface{}) []interface{} {
    // Convertir el mapa en un slice de pares (valor, puntuación) para ordenarlo
    type pair struct {
        value interface{}
        score float64
    }

    pairs := make([]pair, 0, len(values))
    for _, value := range values {
        score := ss.items[value]
        pairs = append(pairs, pair{value, score})
    }

    // Ordenar los pares según la puntuación utilizando el comparador
    for i := 0; i < len(pairs); i++ {
        for j := i + 1; j < len(pairs); j++ {
            if !ss.comparator(pairs[i].score, pairs[j].score) {
                pairs[i], pairs[j] = pairs[j], pairs[i]
            }
        }
    }

    // Extraer los valores ordenados
    sortedValues := make([]interface{}, len(pairs))
    for i, p := range pairs {
        sortedValues[i] = p.value
    }

    return sortedValues
}
