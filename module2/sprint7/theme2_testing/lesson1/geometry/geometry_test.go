package geometry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHypotenuse(t *testing.T) {
	h := Hypotenuse(3, 4)

	// if h != 5 {
	//     t.Errorf("expected %0.2f, got %0.2f", 5.0, h)
	// }
	assert.Equal(t, 5.0, h)
}

func TestHypotenuseTask(t *testing.T) {
	table := []struct {
		a, b   float64
		result string
	}{
		{0, 0, "0.00"},
		{0, 9, "9.00"},
		{3, 4, "5.00"},
		{10, 21, "23.26"},
		{56, 37, "67.12"},
		{102, 67, "122.04"},
		{34, 17, "38.01"},
		{3456, 1089, "3623.51"},
		{478, 201, "518.54"},
	}

	for _, entry := range table {
		h := Hypotenuse(entry.a, entry.b)
		result := fmt.Sprintf("%.2f", h)
		assert.Equal(t, entry.result, result)
	}

}
