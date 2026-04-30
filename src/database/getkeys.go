package database

import (
	"sync"
)

type getKeysCommand struct {
	hook func(result []string, err error)
}

func (cmd getKeysCommand) exec(db *kvDatabase) {
	keys := getMapKeys(db.data)
	cmd.hook(keys, nil)
}

func (db *kvDatabase) GetKeys() ([]string, error) {
	var err error = nil
	var val []string = nil
	var wg sync.WaitGroup
	wg.Add(1)
	db.command_queue <- getKeysCommand{
		hook: func(result []string, e error) {
			val = result
			err = e
			wg.Done()
		},
	}
	wg.Wait()
	return val, err
}
