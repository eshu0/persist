package SQLL

import (
	"database/sql"

	per "github.com/eshu0/persist/pkg/interfaces"
	sl "github.com/eshu0/simpleslogger"
	_ "github.com/mattn/go-sqlite3"
)

type SQLLiteDatastore struct {
	sl.AppLogger

	per.IPersistantStorage

	database *sql.DB

	Filename string

	StorageHandlers map[string]per.IStorageHandler
}

func CreateSQLLiteDatastore(filename string) *SQLLiteDatastore {
	sqlds := SQLLiteDatastore{}
	sqlds.Filename = filename
	StorageHandlers := make(map[string]per.IStorageHandler)
	sqlds.StorageHandlers = StorageHandlers
	return &sqlds
}

func CreateOpenSQLLiteDatastore(filename string) *SQLLiteDatastore {
	sqlds := CreateSQLLiteDatastore(filename)
	sqlds.Open()
	return sqlds
}

func (sqlds *SQLLiteDatastore) Open() {
	// further checks to be added here like checking filepath is correct etc
	dp, err := sql.Open("sqlite3", sqlds.Filename)
	if err != nil {
		panic(err)
	}
	sqlds.database = dp
}

// Storage Handlers
// in this implementation that is Tables

func (sqlds *SQLLiteDatastore) GetStorageHandler(name string) (per.IStorageHandler, bool) {
	res, ok := sqlds.StorageHandlers[name]
	return res, ok
}

func (sqlds *SQLLiteDatastore) SetStorageHander(name string, handler per.IStorageHandler) {
	sqlds.LogDebugf("SetStorageHander", "Setting %s, %v", name, handler)
	sqlds.StorageHandlers[name] = handler
	sqlds.LogDebugf("SetStorageHander", "Handlers = %v", name, sqlds.StorageHandlers)
}

func (sqlds *SQLLiteDatastore) RemoveStorageHandler(name string) bool {
	_, ok := sqlds.StorageHandlers[name]
	if ok {
		delete(sqlds.StorageHandlers, name)
	}
	return ok
}

func (sqlds *SQLLiteDatastore) GetAllStorageHandlers() map[string]per.IStorageHandler {
	return sqlds.StorageHandlers
}

func (sqlds *SQLLiteDatastore) GetDatabase() *sql.DB {
	return sqlds.database
}

func (sqlds *SQLLiteDatastore) CreateStructures() per.IQueryResult {

	success := NewEmptySucceedSQLLiteQueryResult()
	for key, element := range sqlds.StorageHandlers {
		sqlds.LogDebugf("CreateStructures", "Handling %s, %v", key, element)
		res := element.CreateStructures()
		if !res.QuerySucceeded() {
			success = NewEmptyFailedSQLLiteQueryResult()
		}
	}
	return success
}
