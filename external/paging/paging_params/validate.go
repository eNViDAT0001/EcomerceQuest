package paging_params

type EntityFilter interface {
	WithFields() []string
	SearchFields() []string
	SortFields() []string
	CompareFields() []string
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

	compareFields := filter.CompareFields()
	inValidKey, value = validateCompare(fields.GetCompare(), compareFields...)
	if len(inValidKey) > 0 {
		return inValidKey, value
	}

	return "", ""
}

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

func validateCompare(filter *map[string]string, fields ...string) (inValidKey string, value string) {
	if filter == nil {
		return "", ""
	}

	inValid := true
	for key, _ := range *filter {
		for _, v := range fields {
			if key == v {
				inValid = false
				break
			}

			column, condition := GetColumnAndCondition(key)

			inValidKey = column
			value = condition
		}

		if inValid {
			return inValidKey, value
		}
	}

	return "", ""
}

func GetColumnAndCondition(v string) (string, string) {
	var column string
	var condition string

	index := len(v) - 1
	for index >= 0 {
		if v[index] == '_' {
			break
		}
		index--
	}

	if index == -1 {
		return "", ""
	}

	column = v[:index]
	condition = v[index+1:]

	return column, condition
}
