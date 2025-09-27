package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertAssign(t *testing.T) {
	{
		var dst int64
		err := Assign(&dst, 10)
		assert.NoError(t, err)
		assert.Equal(t, int64(10), dst)
	}
	{
		var dst string
		err := Assign(&dst, 10)
		assert.NoError(t, err)
		assert.Equal(t, "10", dst)
	}
	{
		type color int
		var dst color
		err := Assign(&dst, 10)
		assert.NoError(t, err)
		assert.Equal(t, color(10), dst)
	}
	{
		type color int
		var dst int
		err := Assign(&dst, color(10))
		assert.NoError(t, err)
		assert.Equal(t, 10, dst)
	}
}
