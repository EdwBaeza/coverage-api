package utils

func SumMap[K comparable, V int64 | float64](m []map[K]interface{}, valueFunc func(map[K]interface{}) V) V {
	var sum V
	for _, item := range m {
		sum += valueFunc(item)
	}
	return sum
}

func Map[K, V string](m []map[K]interface{}, valueFunc func(map[K]interface{}) V) []V {
	var result []V
	for _, item := range m {
		result = append(result, valueFunc(item))
	}
	return result
}

func Sum[K comparable, V int64 | float64](m []K, valueFunc func(K) V) V {
	var sum V
	for _, item := range m {
		sum += valueFunc(item)
	}
	return sum
}
