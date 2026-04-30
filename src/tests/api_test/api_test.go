package api_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGetString(t *testing.T) {
	t.Parallel()

	api := makeAPI()
	key := randKey(10)
	val := randKey(20)

	err := api.set(key, val)
	require.NoError(t, err)

	result, err := api.get(key)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("\"%s\"", val), result)
}
