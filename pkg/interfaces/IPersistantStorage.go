package persist

import (
	sli "github.com/eshu0/simplelogger/interfaces"
)

type IPersistantStorage interface {

	// Get/Set the logging for the interface
	GetStorageHandler(name string) IStorageHandler
	SetStorageHander(name string, store IStorageHandler)
	GetAllStorageHandlers() map[string]IStorageHandler
	RemoveStorageHandler(name string) bool

	// Get/Set the logging for the interface
	GetLog() sli.ISimpleLogger
	SetLog(logger sli.ISimpleLogger)

	// This function creates all the handlers
	CreateHandlers() bool

	// This function creates all the structures for all the handlers
	// this could be files, tables etc
	CreateStrutures() bool

	// Wipe all data
	Wipe() bool

}
