package persist

type IStorageHandler interface {

	// Make sure that the Peristant storage is available
	// Thsiw ill have access to files, database, logging etc
	GetPersistantStorage() IPersistantStorage
	SetPersistantStorage(persistant IPersistantStorage) //IStorageHandler

	// This function creates all the structures that are needed for storage
	// this could be files, tables etc
	CreateStructures() IQueryResult

	// This might be over kill
	// Maybe all reading/writing should be handled by implementation not here
	
	// Wipe all data
	//Wipe() IQueryResult
	//ReadAll() IQueryResult

	// CRUD operations
	//Create(data IDataItem) IQueryResult
	//Read(data IDataItem)   IQueryResult
	//Update(data IDataItem) IQueryResult
	//Delete(data IDataItem) IQueryResult

}
