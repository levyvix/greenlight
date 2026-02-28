package data

import (
	"fmt"

	"github.com/levyvix/greenlight-api/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func Validatefilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be a positive integer value")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be a positive integer value")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", fmt.Sprintf("must be one of the following values: %v", f.SortSafeList))
}
