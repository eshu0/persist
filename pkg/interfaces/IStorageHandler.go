package persist

import (
	sli "github.com/eshu0/simplelogger/interfaces"
)

type IStorageHandler interface {

	GetPersistantStorage() IPersistantStorage
	SetPersistantStorage(persistant IPersistantStorage)

	// This function creates all the structures that are needed for storage
	// this could be files, tables etc
	CreateStrutures() bool

	// Wipe all data
	Wipe() bool
	ReadAll() []IDataItem

	// CRUD operations
	Create(data IDataItem) bool
	Read(data IDataItem)   IDataItem
	Update(data IDataItem) bool
	Delete(data IDataItem) bool

}
