package helper

import (
	"math"
	"strconv"
)

// Pagination calculates the limit and offset for pagination
func Pagination(page string, size string) (int, int, int) {
	intPage, err := strconv.Atoi(page)
	if err != nil || intPage < 1 {
		intPage = 1 // Default to page 1 if parsing fails or page is less than 1
	}

	limit := 10 // Default limit
	if size != "" {
		parsedSize, err := strconv.Atoi(size)
		if err == nil && parsedSize > 0 {
			limit = parsedSize
		}
	}

	offset := 0
	if intPage > 1 {
		offset = (intPage - 1) * limit
	}

	return limit, offset, intPage
}

// PaginateData formats the paginated data
func PaginateData(datas []interface{}, totalItems int, page int, limit int) map[string]interface{} {
	currentPage := page
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	pageSize := limit

	return map[string]interface{}{
		"totalItems":  totalItems,
		"pageSize":    pageSize,
		"totalPages":  totalPages,
		"currentPage": currentPage,
		"data":        datas,
	}
}
