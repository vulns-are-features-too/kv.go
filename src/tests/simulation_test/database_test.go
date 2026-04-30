// Package simulation_test Simulation tests on the database
package simulation_test

import (
	"database"
	"sync"
	"testing"
)

func TestSimulation(t *testing.T) {
	t.Parallel()

	// CONFIG
	agentsCount := 10

	// ARRANGE
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	agents := make([]testAgent, agentsCount)
	for i := range agentsCount {
		agents[i] = makeAgent(t, i, db)
	}

	// ACT & ASSERT
	var wg sync.WaitGroup
	wg.Add(agentsCount)

	for i := range agentsCount {
		go func() {
			agents[i].run(100)
			wg.Done()
		}()
	}

	wg.Wait()

	for _, a := range agents {
		a.assertKnownData()
	}
}
