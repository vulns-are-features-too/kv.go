// Package unit_tests Unit tests
package unit_test

import (
	"database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGetString(t *testing.T) {
	t.Parallel()

	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	err := db.Set("s", "string")
	require.NoError(t, err)
	s, err := db.Get("s")
	require.NoError(t, err)
	assert.Equal(t, "string", s)
}
