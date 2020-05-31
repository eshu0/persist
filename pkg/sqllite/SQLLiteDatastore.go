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

	StorageHandlers map[string] IStorageHandler
}


func  RemoveStorageHandler(name string)cv{
	sqlds.StorageHandlers[name] = nil
	return true
}

func  GetStorageHandler(name string) per.IStorageHandlerv{
	return sqlds.StorageHandlers[name]
}

func  SetStorageHander(name string, store IStorageHandler) {
	sqlds.StorageHandlers[name] = store
}

func (sqlds *SQLLiteDatastore) GetAllStorageHandlers() map[string]IStorageHandler{
	return sqlds.StorageHandlers
}
	
// Get/Set the logging for the interface
func (sqlds *SQLLiteDatastore) GetLog() sli.ISimpleLogger{
	return sqlds.Log
}

func (sqlds *SQLLiteDatastore) SetLog(logger sli.ISimpleLogger){
	sqlds.Log = logger
}

func (sqlds *SQLLiteDatastore) CreateHandlers() bool {

return true
}

func (sqlds *SQLLiteDatastore) 	CreateStrutures() bool {
	return true
}

func (sqlds *SQLLiteDatastore) Wipe() bool{
	return true
}
