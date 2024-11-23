package data

import "greenlight.richie.net/internal/validator"

type Filters struct {
	Page     int
	PageSize int
	Sort 	 string
	SortList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0 && f.Page <= 10_000_000, "page", "must be between 1 and 10 million")
	v.Check(f.PageSize > 0 && f.PageSize <= 100, "page_size", "must be between 1 and 100")
	v.Check(validator.In(f.Sort, f.SortList...), "sort", "invalid sort value")
}