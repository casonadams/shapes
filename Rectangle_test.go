package shape_test

import (
	"testing"

	shape "github.com/casonadams/shapes"
	assert "github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	r := shape.NewRectangle(9, 6)
	actual_result := r.Area()
	expected_result := float64(54)
	assert.Equal(t, actual_result, expected_result)
}

func BenchmarkRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shape.NewRectangle(9, 6)
	}
}

func BenchmarkRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shape.NewRectangle(9, 6).Area()
	}
}
