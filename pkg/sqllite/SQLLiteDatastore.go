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


func (sqlds *SQLLiteDatastore) RemoveStorageHandler(name string) bool {
	sqlds.StorageHandlers[name] = nil
	return true
}

func (sqlds *SQLLiteDatastore) GetStorageHandler(name string) per.IStorageHandler{
	return sqlds.StorageHandlers[name]
}

func (sqlds *SQLLiteDatastore) SetStorageHander(name string, store per.IStorageHandler) {
	sqlds.StorageHandlers[name] = store
}

func (sqlds *SQLLiteDatastore) GetAllStorageHandlers() map[string]per.IStorageHandler{
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
