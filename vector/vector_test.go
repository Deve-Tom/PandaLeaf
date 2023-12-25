package vector

import (
	"fmt"
	"testing"
)

func TestVector_Front(t *testing.T) {
	t.Parallel()

	var data Vector[int]
	items, err := data.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(items)
	}
}
