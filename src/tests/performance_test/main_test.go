package performance_test

import (
	"common"
	"database"
	"math/rand"
	"sync"
	"testing"
)

const (
	agentsCount   = 20
	keysCount     = 100
	itersPerAgent = 1_000_000
)

func BenchmarkDatabase(b *testing.B) {
	db := database.MakeDatabase()

	db.Start()
	defer db.Stop()

	keys := makeKeys(keysCount)
	for _, k := range keys {
		_ = db.Set(k, "")
	}

	b.ResetTimer()

	for range b.N {
		benchmarkOnce(db, keys)
	}

	b.StopTimer()
}

func benchmarkOnce(db database.KvDatabase, keys []string) {
	var wg sync.WaitGroup
	wg.Add(agentsCount)

	for i := range agentsCount {
		go func() {
			runAgent(db, keys, int64(i))
			wg.Done()
		}()
	}

	wg.Wait()
}

func makeKeys(keysCount uint) []string {
	keys := make([]string, keysCount)
	for i := range keysCount {
		keys[i] = common.RandKey(10)
	}

	return keys
}

func runAgent(db database.KvDatabase, keys []string, id int64) {
	//nolint:gosec
	rng := rand.New(rand.NewSource(id))
	for range itersPerAgent {
		key := keys[rng.Intn(len(keys))]

		r := rng.Intn(100)
		switch {
		case r <= 70:
			_, _ = db.Get(key)
		case r <= 95:
			_ = db.Set(key, common.RandKey(20))
		case r <= 99:
			_ = db.Copy(key, keys[rng.Intn(len(keys))])
		default:
			_, _ = db.GetKeys()
		}
	}
}
