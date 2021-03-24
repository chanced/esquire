package jsonutil_test

import (
	"testing"

	"github.com/chanced/picker/internal/jsonutil"
	"github.com/stretchr/testify/require"
)

func TestEscapeKey(t *testing.T) {
	assert := require.New(t)

	assert.Equal(`path\.with\.period`, jsonutil.EscapeKey("path.with.period"))
	assert.Equal(`path\:with\:colon`, jsonutil.EscapeKey("path:with:colon"))
	assert.Equal(`path\.with\:both`, jsonutil.EscapeKey("path.with:both"))

}
