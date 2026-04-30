// Package integration_test for Integration tests
package integration_test

import (
	"database"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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
