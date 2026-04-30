package database

import (
	"reflect"
	"sync"
)

type getCommand struct {
	key  string
	hook func(result any, err error)
}

func (cmd getCommand) exec(db *kvDatabase) {
	val, ok := db.data[cmd.key]
	if ok {
		cmd.hook(val, nil)
	} else {
		cmd.hook(nil, NotFoundError(cmd.key))
	}
}

func (db *kvDatabase) Get(key string) (string, error) {
	var err error = nil
	var val any = nil
	var wg sync.WaitGroup
	wg.Add(1)
	db.command_queue <- getCommand{
		key: key,
		hook: func(result any, e error) {
			val = result
			err = e
			wg.Done()
		},
	}
	wg.Wait()

	if err != nil {
		return "", err
	}

	res, ok := val.(string)
	if !ok {
		return "", IncompatibleTypeError("string", reflect.TypeOf(val).Name())
	}

	return res, err
}
