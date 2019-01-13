package shape_test

import (
	"testing"

	shape "github.com/casonadams/shapes"
	assert "github.com/stretchr/testify/assert"
)

func TestTriangleArea(t *testing.T) {
	tri := shape.NewTriangle(3, 6)
	actual_result := tri.Area()
	expected_result := float64(9)
	assert.Equal(t, actual_result, expected_result)
}

func BenchmarkTriangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shape.NewTriangle(3, 6)
	}
}
