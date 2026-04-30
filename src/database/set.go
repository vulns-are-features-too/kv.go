package database

import "sync"

type setCommand struct {
	key   string
	value string
	hook  func(err error)
}

func (cmd setCommand) exec(db *kvDatabase) {
	db.data[cmd.key] = cmd.value
	cmd.hook(nil)
}

func (db *kvDatabase) Set(key string, value string) error {
	var err error = nil
	var wg sync.WaitGroup
	wg.Add(1)
	db.command_queue <- setCommand{
		key:   key,
		value: value,
		hook: func(e error) {
			err = e
			wg.Done()
		},
	}
	wg.Wait()
	return err
}
