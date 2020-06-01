package SQLL

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	sli "github.com/eshu0/simplelogger/interfaces"
	per "github.com/eshu0/persist/pkg/interfaces"
)   

type SQLLiteDatastore struct {
	per.IPersistantStorage

	database *sql.DB

	Filename string

	Log sli.ISimpleLogger

	StorageHandlers map[string] per.IStorageHandler
}

func CreateSQLLiteDatastore(log sli.ISimpleLogger, filename string) *SQLLiteDatastore {
	sqlds := SQLLiteDatastore{}
	sqlds.Filename =filename
	StorageHandlers := make(map[string]per.IStorageHandler)
	sqlds.StorageHandlers = StorageHandlers
	sqlds.SetLog(log)
	return &sqlds
}

func CreateOpenSQLLiteDatastore(log sli.ISimpleLogger, filename string) *SQLLiteDatastore {
	sqlds := CreateSQLLiteDatastore(log, filename)
	sqlds.Open()
	return sqlds
}


func (sqlds *SQLLiteDatastore) Open() {
	// further checks to be added here like checking filepath is correct etc
	sqlds.database, _ = sql.Open("sqlite3", sqlds.Filename)
}

// Storage Handlers 
// in this implementation that is Tables

func (sqlds *SQLLiteDatastore) RemoveStorageHandler(name string) bool {
	_, ok := sqlds.StorageHandlers[name]
    if ok {
        delete(sqlds.StorageHandlers, name)
    }
	return ok
}

func (sqlds *SQLLiteDatastore) GetStorageHandler(name string) (per.IStorageHandler, bool) {
	res, ok :=  sqlds.StorageHandlers[name]
	return res,ok
}

func (sqlds *SQLLiteDatastore) SetStorageHander(name string, store per.IStorageHandler) {
	store.SetPersistantStorage(sqlds)
	sqlds.StorageHandlers[name] = store
}

func (sqlds *SQLLiteDatastore) GetAllStorageHandlers() map[string]per.IStorageHandler{
	return sqlds.StorageHandlers
}
	
// Get/Set the logging for the interface
func (sqlds *SQLLiteDatastore) GetLog() sli.ISimpleLogger {
	return sqlds.Log
}

func (sqlds *SQLLiteDatastore) SetLog(logger sli.ISimpleLogger){
	sqlds.Log = logger
}

func (sqlds *SQLLiteDatastore) GetDatabase() *sql.DB {
	return sqlds.database
}

func (sqlds *SQLLiteDatastore) CreateStructures() per.IQueryResult {
	
	success := NewEmptySucceedSQLLiteQueryResult()
	for _, element := range sqlds.StorageHandlers {
		res := element.CreateStructures()
		if res.QuerySucceeded() {

		}else {
			success = NewEmptyFailedSQLLiteQueryResult()
		}
	}
	return success
}

func (sqlds *SQLLiteDatastore) Wipe() per.IQueryResult{
	success := NewEmptySucceedSQLLiteQueryResult()
	for _, element := range sqlds.StorageHandlers {
		res :=  element.Wipe() 
		if res.QuerySucceeded() {

		}else {
			success = NewEmptyFailedSQLLiteQueryResult()
		}
	}
	return success
}
