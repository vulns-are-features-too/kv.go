package database

import (
	"fmt"
	"sync"
)

type KvDatabase interface {
	Start()
	Stop()
	Get(key string) (string, error)
	GetKeys() ([]string, error)
	Set(key, value string) error
}

type command interface {
	exec(db *kvDatabase)
}

type kvDatabase struct {
	data          map[string]any
	command_queue chan command
	running       bool
	wg            sync.WaitGroup
}

func MakeDatabase() KvDatabase {
	return &kvDatabase{
		data:          make(map[string]any),
		command_queue: make(chan command),
		running:       false,
	}
}

func (db *kvDatabase) Start() {
	db.wg.Add(1)
	db.running = true

	go func() {
		for db.running {
			cmd := <-db.command_queue
			cmd.exec(db)
		}
	}()
}

func (db *kvDatabase) Stop() {
	db.running = false
	db.wg.Done()
	db.wg.Wait()
}

func NotFoundError(key string) error {
	return fmt.Errorf("key not found: %s", key)
}

func IncompatibleTypeError(expected, actual string) error {
	return fmt.Errorf("Incompatible types: expected %s, got %s", expected, actual)
}
