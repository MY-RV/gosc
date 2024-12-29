package map_helper

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
    result := make(map[K]V, len(m))
    for k, v := range m {
        result[k] = v
    }
    return result
}
