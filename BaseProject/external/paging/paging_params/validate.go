package paging_params

type EntityFilter interface {
	WithFields() []string
	SearchFields() []string
	SortFields() []string
}

func ValidateFilter(fields FilterList, filter EntityFilter) (inValidKey string, value string) {

	withFields := filter.WithFields()
	inValidKey, value = validateFilter(fields.GetFields(), withFields...)
	if len(inValidKey) > 0 {
		return inValidKey, value
	}

	searchFields := filter.SearchFields()
	inValidKey, value = validateFilter(fields.GetSearch(), searchFields...)
	if len(inValidKey) > 0 {
		return inValidKey, value
	}

	sortFields := filter.SortFields()
	inValidKey, value = validateFilter(fields.GetSort(), sortFields...)
	if len(inValidKey) > 0 {
		return inValidKey, value
	}

	return "", ""
}

// TODO: Remove this function in some day
func validateFilter(filter *map[string]string, fields ...string) (inValidKey string, value string) {
	if filter == nil {
		return "", ""
	}

	inValid := true
	for key, val := range *filter {

		for _, v := range fields {
			if v == key {
				inValid = false
				break
			}
			inValidKey = key
			value = val
		}

		if inValid {
			return inValidKey, value
		}
	}

	return "", ""
}
