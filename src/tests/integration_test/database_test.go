// Package integration_test for Integration tests
package integration_test

import (
	"database"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests on the database's sync mechanisms

func TestSetGetStringLoop(t *testing.T) {
	t.Parallel()

	loops := 1000
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	var wg sync.WaitGroup

	wg.Add(loops)

	for l := range loops {
		go func(i int) {
			key := fmt.Sprintf("k%d", i)
			val := fmt.Sprintf("v%d", i)
			err := db.Set(key, val)
			assert.NoError(t, err)
			result, err := db.Get(key)
			assert.NoError(t, err)
			assert.Equal(t, val, result)
			wg.Done()
		}(l)
	}

	wg.Wait()

	keys, err := db.GetKeys()
	require.NoError(t, err)
	assert.Len(t, keys, loops)

	wg.Add(loops)

	for l := loops - 1; l >= 0; l-- {
		go func(i int) {
			key := fmt.Sprintf("k%d", i)
			val := fmt.Sprintf("v%d", i)
			result, err := db.Get(key)
			assert.NoError(t, err)
			assert.Equal(t, val, result)
			wg.Done()
		}(l)
	}

	wg.Wait()
}

func TestCopyEachKeyManyTimesWith1GoroutinePerBaseKeyValue(t *testing.T) {
	t.Parallel()

	// ARRANGE
	loops := 100
	baseMap := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}
	baseLen := len(baseMap)

	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	for k, v := range baseMap {
		require.NoError(t, db.Set(k, v))
	}

	var wg sync.WaitGroup
	wg.Add(baseLen)

	// ACT
	for k := range baseMap {
		go func(baseKey string) {
			oldKey := baseKey
			newKey := fmt.Sprintf("%s:%d", baseKey, 0)

			for l := range loops {
				assert.NoError(t, db.Copy(oldKey, newKey))
				oldKey = newKey
				newKey = fmt.Sprintf("%s:%d", baseKey, l)
			}

			wg.Done()
		}(k)
	}

	wg.Wait()

	// ASSERT
	for k, v := range baseMap {
		for l := range loops - 1 {
			result, err := db.Get(fmt.Sprintf("%s:%d", k, l))
			require.NoError(t, err)
			assert.Equal(t, v, result)
		}
	}

	keys, err := db.GetKeys()
	require.NoError(t, err)
	assert.Len(t, keys, loops*baseLen)
}
