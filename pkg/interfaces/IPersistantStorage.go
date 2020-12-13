package persist

type IPersistantStorage interface {

	// Get/Set the logging for the interface
	GetStorageHandler(name string) (IStorageHandler, bool)
	SetStorageHander(name string, store IStorageHandler)
	GetAllStorageHandlers() map[string]IStorageHandler
	RemoveStorageHandler(name string) bool

	// This function creates all the structures for all the handlers
	// this could be files, tables etc
	CreateStructures() IQueryResult

	// Wipe all data
	//Wipe() IQueryResult

}
