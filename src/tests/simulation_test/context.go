package simulation_test

import (
	"database"
	"testing"
)

type testContext struct {
	t  *testing.T
	db database.KvDatabase
}

func makeContext(t *testing.T, db database.KvDatabase) testContext {
	t.Helper()

	return testContext{
		t:  t,
		db: db,
	}
}
