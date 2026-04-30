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

func TestGetKeys(t *testing.T) {
	t.Parallel()

	api := makeAPI()
	key1 := randKey(20)
	key2 := randKey(20)

	require.NoError(t, api.set(key1, ""))
	require.NoError(t, api.set(key2, ""))

	result, err := api.getKeys()
	require.NoError(t, err)
	assert.Subset(t, result, []string{key1, key2})
}

func TestCopy(t *testing.T) {
	t.Parallel()

	api := makeAPI()
	key1 := randKey(10)
	key2 := randKey(20)
	val := "value"

	require.NoError(t, api.set(key1, val))
	require.NoError(t, api.copy(key1, key2))

	keys, err := api.getKeys()
	require.NoError(t, err)
	assert.Subset(t, keys, []string{key1, key2})

	newVal, err := api.get(key2)
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("\"%s\"", val), newVal)
}
