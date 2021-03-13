package sorting

import (
	"sort"
)

type Info []struct {
	name  string
	value int
}

func (in Info) Len() int {
	return len(in)
}

func (in Info) Less(i, j int) bool {
	if in[i].value < in[j].value {
		return true
	}
	if in[i].value == in[j].value && in[i].name < in[j].name {
		return true
	}
	return false
}

func (in Info) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

// sortParam = by name, popularity, price
// sortOrder = asc, desc
// O(nlogn)
func getItems(items map[string][2]int, sortParam, sortOrder, itemsPerPage, pageNumber int) []string {
	offset := pageNumber * itemsPerPage
	limit := itemsPerPage + offset
	if limit >= len(items) {
		limit = 0
	}

	// name
	if sortParam == 0 {
		sortedNames := make([]string, 0, len(items))

		for name := range items {
			sortedNames = append(sortedNames, name)
		}
		// O(nlogn) for sorting
		if sortOrder == 0 {
			sort.Strings(sortedNames)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(sortedNames)))
		}
		if limit != 0 {
			return sortedNames[offset:limit]
		}
		return sortedNames[offset:]
	}

	var sortedItems Info

	for name, value := range items {
		sortedItems = append(sortedItems, struct {
			name  string
			value int
		}{name: name, value: value[sortParam-1]})
	}

	// O(n*logn) for sorting
	if sortOrder == 0 {
		sort.Sort(sortedItems)
	} else {
		sort.Sort(sort.Reverse(sortedItems))
	}

	if limit != 0 {
		sortedItems = sortedItems[offset:limit]
	} else {
		sortedItems = sortedItems[offset:]
	}

	var result []string
	for _, el := range sortedItems {
		result = append(result, el.name)
	}
	return result
}
