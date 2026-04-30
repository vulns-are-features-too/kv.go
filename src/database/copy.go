package database

import "sync"

type copyCommand struct {
	src  string
	dst  string
	hook func(err error)
}

func (cmd copyCommand) exec(db *kvDatabase) {
	val, ok := db.data[cmd.src]
	if !ok {
		cmd.hook(notFoundError(cmd.src))
		return
	}

	db.data[cmd.dst] = val
	cmd.hook(nil)
}

func (db *kvDatabase) Copy(srcKey string, dstKey string) error {
	if srcKey == dstKey {
		return nil
	}

	var err error = nil
	var wg sync.WaitGroup
	wg.Add(1)
	db.command_queue <- copyCommand{
		src: srcKey,
		dst: dstKey,
		hook: func(e error) {
			err = e
			wg.Done()
		},
	}
	wg.Wait()

	return err
}
