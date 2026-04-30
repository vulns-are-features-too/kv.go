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

func TestGetKeys(t *testing.T) {
	t.Parallel()

	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	require.NoError(t, db.Set("k1", "v"))
	require.NoError(t, db.Set("k2", "v"))
	keys, err := db.GetKeys()
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"k1", "k2"}, keys)
}

func TestCopyString(t *testing.T) {
	t.Parallel()

	// ARRANGE
	k1 := "k1"
	k2 := "k2"
	val := "value"
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	require.NoError(t, db.Set(k1, val))

	// ACT
	require.NoError(t, db.Copy(k1, k2))

	// ASSERT
	keys, err := db.GetKeys()
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{k1, k2}, keys)

	result, err := db.Get(k2)
	require.NoError(t, err)
	assert.Equal(t, val, result)
}

func TestCopyNonExistentKeyReturnsError(t *testing.T) {
	t.Parallel()

	// ARRANGE
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	// ACT
	err := db.Copy("k1", "k2")

	// ASSERT
	require.Error(t, err)
	assert.Equal(t, "key not found: k1", err.Error())
}

func TestCopySrcEqualsDst(t *testing.T) {
	t.Parallel()

	// ARRANGE
	key := "k"
	val := "value"
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	require.NoError(t, db.Set(key, val))

	// ACT
	require.NoError(t, db.Copy(key, key))

	// ASSERT
	keys, err := db.GetKeys()
	require.NoError(t, err)
	assert.Equal(t, []string{key}, keys)

	result, err := db.Get(key)
	require.NoError(t, err)
	assert.Equal(t, val, result)
}
