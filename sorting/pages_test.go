package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetItems(t *testing.T) {
	{
		sortParam := 0 // 0 1 2
		sortOrder := 0 // 0 1
		items := map[string][2]int{
			"item1": {2, 10},
			"item2": {10, 5},
			"item3": {1, 15},
			"item4": {1, 7},
			"item5": {1, 1},
			"item6": {1, 20},
			"item7": {1, 4},
		}
		perPage := 3
		pageNumber := 1
		result := getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item4", "item5", "item6"}, result)

		sortParam = 1
		result = getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item6", "item7", "item1"}, result)

		sortParam = 2
		result = getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item4", "item1", "item3"}, result)

		sortOrder = 1
		result = getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item4", "item2", "item7"}, result)

		pageNumber = 2
		result = getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item5"}, result)

		perPage = 2
		result = getItems(items, sortParam, sortOrder, perPage, pageNumber)
		require.Equal(t, []string{"item2", "item7"}, result)
	}
}
