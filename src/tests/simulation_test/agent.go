package simulation_test

import (
	"math/rand"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testAgent struct {
	id        int
	ctx       testContext
	knownData map[string]any
}

func makeAgent(id int, ctx testContext) testAgent {
	ta := testAgent{
		id:        id,
		ctx:       ctx,
		knownData: make(map[string]any),
	}

	return ta
}

func (ta testAgent) getRandKnownData() (string, any) {
	//nolint:gosec
	randIdx := rand.Intn(len(ta.knownData))

	i := 0
	for k, v := range ta.knownData {
		if i == randIdx {
			return k, v
		}

		i++
	}

	panic("testAgent.getRandKnownData out of range")
}

func (ta testAgent) run(iterations int64) {
	// seed init data for other actions
	setAction{}.run(ta, 0)

	for i := range iterations {
		ta.ctx.getRandAction().run(ta, i)
	}
}

func (ta testAgent) assertKnownData() {
	for k, v := range ta.knownData {
		val, err := ta.ctx.db.Get(k)
		require.NoError(ta.ctx.t, err)
		assert.Equal(ta.ctx.t, v, val)
	}
}
