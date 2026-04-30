package simulation_test

import (
	"database"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testAgent struct {
	id        int
	t         *testing.T
	db        database.KvDatabase
	actions   []action
	knownData map[string]any
}

func makeAgent(t *testing.T, id int, db database.KvDatabase) testAgent {
	t.Helper()
	ta := testAgent{
		id:        id,
		t:         t,
		db:        db,
		knownData: make(map[string]any),
		actions:   nil,
	}
	ta.actions = []action{
		setAction{ta: ta},
		getAction{ta: ta},
		getKeysAction{ta: ta},
		copyAction{ta: ta},
	}

	return ta
}

//nolint:ireturn
func (ta testAgent) getRandAction() action {
	//nolint:gosec
	return ta.actions[rand.Intn(len(ta.actions))]
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
	setAction{ta: ta}.run(0)

	for i := range iterations {
		ta.getRandAction().run(i)
	}
}

func (ta testAgent) assertKnownData() {
	for k, v := range ta.knownData {
		val, err := ta.db.Get(k)
		require.NoError(ta.t, err)
		assert.Equal(ta.t, v, val)
	}
}
