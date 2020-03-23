package features

import "math"

func CalculatePageCount(objectsCount, itemsPerPage int) int {
	floatResult := float64(objectsCount) / float64(itemsPerPage)
	ceil := math.Ceil(floatResult)
	result := int(ceil)
	return result
}
