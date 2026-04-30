package simulation_test

import (
	"fmt"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type action interface {
	run(iteration int64)
}

type setAction struct {
	ta testAgent
}

func (a setAction) run(i int64) {
	key := fmt.Sprintf("a%d:s%d", a.ta.id, i)
	val := strconv.FormatInt(i, 10)
	a.ta.knownData[key] = val
	err := a.ta.db.Set(key, val)
	require.NoError(a.ta.t, err)
}

type getAction struct {
	ta testAgent
}

func (a getAction) run(_ int64) {
	key, val := a.ta.getRandKnownData()
	result, err := a.ta.db.Get(key)
	require.NoError(a.ta.t, err)
	assert.Equal(a.ta.t, val, result)
}

type getKeysAction struct {
	ta testAgent
}

func (a getKeysAction) run(_ int64) {
	allKeys, err := a.ta.db.GetKeys()
	require.NoError(a.ta.t, err)
	ownKeys := getMapKeys(a.ta.knownData)
	assert.Subset(a.ta.t, allKeys, ownKeys)
}
