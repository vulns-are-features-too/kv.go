package simulation_test

import (
	"database"
	"math/rand"
	"testing"
)

type testContext struct {
	t       *testing.T
	db      database.KvDatabase
	actions []action
}

func makeContext(t *testing.T, db database.KvDatabase) testContext {
	t.Helper()

	return testContext{
		t:  t,
		db: db,
		actions: []action{
			setAction{},
			getAction{},
			getKeysAction{},
			copyAction{},
		},
	}
}

//nolint:ireturn
func (ctx testContext) getRandAction() action {
	//nolint:gosec
	return ctx.actions[rand.Intn(len(ctx.actions))]
}
