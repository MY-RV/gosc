package map_helper

func GetOrCreate[K comparable, V any](_map map[K]V, key K, callback func() V) V {
	if value, exists := _map[key]; exists {
		return value
	}

	value := callback()
	_map[key] = value
	return value
}