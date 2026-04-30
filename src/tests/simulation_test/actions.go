package simulation_test

import (
	"common"
	"fmt"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type action interface {
	run(ta testAgent, iteration int64)
}

type setAction struct{}

func (a setAction) run(ta testAgent, i int64) {
	key := fmt.Sprintf("a%d:s%d", ta.id, i)
	val := strconv.FormatInt(i, 10)
	ta.knownData[key] = val
	err := ta.ctx.db.Set(key, val)
	require.NoError(ta.ctx.t, err)
}

type getAction struct{}

func (a getAction) run(ta testAgent, _ int64) {
	key, val := ta.getRandKnownData()
	result, err := ta.ctx.db.Get(key)
	require.NoError(ta.ctx.t, err)
	assert.Equal(ta.ctx.t, val, result)
}

type getKeysAction struct{}

func (a getKeysAction) run(ta testAgent, _ int64) {
	allKeys, err := ta.ctx.db.GetKeys()
	require.NoError(ta.ctx.t, err)
	ownKeys := common.GetMapKeys(ta.knownData)
	assert.Subset(ta.ctx.t, allKeys, ownKeys)
}

type copyAction struct{}

func (a copyAction) run(ta testAgent, _ int64) {
	key, val := ta.getRandKnownData()
	newKey := common.RandKey(10)
	ta.knownData[newKey] = val
	err := ta.ctx.db.Copy(key, newKey)
	require.NoError(ta.ctx.t, err)
}
