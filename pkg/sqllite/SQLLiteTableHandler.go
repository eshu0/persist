package SQLL

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	sli "github.com/eshu0/simplelogger/interfaces"
	per "github.com/eshu0/persist/pkg/interfaces"
)   

type SQLLiteTableHandler struct {
	per.IStorageHandler
	Parent SQLLiteDatastore
}

func NewSQLLiteTableHandler(datastore SQLLiteDatastore) *SQLLiteTableHandler {
	ds := SQLLiteTableHandler{}
	ds.SetPersistantStorage(datastore)
	return &ds
}

func (handler *SQLLiteTableHandler) GetPersistantStorage() IPersistantStorage {
	return handler.Parent
}

func (handler *SQLLiteTableHandler) SetPersistantStorage(persistant IPersistantStorage){
	handler.Parent = persistant
}

func (handler *SQLLiteTableHandler) CreateStrutures() bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) Wipe() bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) ReadAll()  []IDataItem {
	// this needs to be implemented
	return nil
}

func (handler *SQLLiteTableHandler) Wipe() bool {
	// this needs to be implemented
	return false
}


func (handler *SQLLiteTableHandler) Create(data IDataItem) bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) Read(data IDataItem) IDataItem {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) Update(data IDataItem) bool {
	// this needs to be implemented
	return false
}


func (handler *SQLLiteTableHandler) Delete(data IDataItem) bool {
	// this needs to be implemented
	return false
}

// This function ProjectsDBStruct removes all data for the table
func (handler *SQLLiteTableHandler) ExecuteQuery(query string) int64 {
	statement, _ := handler.Parent.GetDatabase().Prepare(query)
	res, err := statement.Exec()
	if err ==  nil {
		rowsaff, rerr := res.RowsAffected()
		if rerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteQuery - RowsAffected Error",rerr)
			return -1
		}
		return rowsaff
	}else{
		handler.Parent.GetLog().LogErrorE("ExecuteQuery",err)
		return -1
	}
}

// This adds ProjectsDBStruct to the database 
func (handler *SQLLiteTableHandler) ExecuteInsertQuery(string query,params ...interface{}) int64 {
	statement, _ := handler.Parent.GetDatabase().Prepare(query)
	res, err := statement.Exec(params...)
	if err ==  nil {
		lastid, lerr := res.LastInsertId()
		if lerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery - LastInsertId",lerr)
			return -1
		}
		return lastid
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery",err)
		return -1
	}
}

func (handler *SQLLiteTableHandler) ExecuteResult(query string) []*IDataItem {
	statement, _ := handler.Parent.GetDatabase().Prepare(query)
	rows, err := statement.Query()
	if err ==  nil {
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResult",err)
		empty := []*IDataItem{}
		return empty
	}
}


func (handler *SQLLiteTableHandler) ExecuteResultWithData(query string, params ...interface{}) []*IDataItem {
	statement, _ := handler.Parent.GetDatabase().Prepare(query)
	rows, err := statement.Query(params...)
	if err ==  nil {
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData",err)
		empty := []*IDataItem{}
		return empty
	}
}


func (handler *SQLLiteTableHandler) ParseRows(rows *sql.Rows) []*IDataItem {
	return []*IDataItem{}
}
