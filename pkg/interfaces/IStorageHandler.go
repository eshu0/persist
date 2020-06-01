package persist

type IStorageHandler interface {

	GetPersistantStorage() IPersistantStorage
	SetPersistantStorage(persistant IPersistantStorage)

	// This function creates all the structures that are needed for storage
	// this could be files, tables etc
	CreateStructures() IQueryResult

	// Wipe all data
	Wipe() IQueryResult
	ReadAll() IQueryResult

	// CRUD operations
	Create(data IDataItem) IQueryResult
	Read(data IDataItem)   IQueryResult
	Update(data IDataItem) IQueryResult
	Delete(data IDataItem) IQueryResult

}
