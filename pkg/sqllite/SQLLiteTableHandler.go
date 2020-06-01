package SQLL

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	//sli "github.com/eshu0/simplelogger/interfaces"
	per "github.com/eshu0/persist/pkg/interfaces"
)   

type SQLLiteTableHandler struct {
	per.IStorageHandler
	Parent *SQLLiteDatastore
}

func NewSQLLiteTableHandler(datastore *SQLLiteDatastore) *SQLLiteTableHandler {
	ds := SQLLiteTableHandler{}
	ds.SetPersistantStorage(datastore)
	return &ds
}

// Start IStorage Handler 

func (handler *SQLLiteTableHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *SQLLiteTableHandler) SetPersistantStorage(persistant per.IPersistantStorage){
	res := persistant.(*SQLLiteDatastore)
	handler.Parent = res
}

func (handler *SQLLiteTableHandler) CreateStructures() bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) Wipe() bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) ReadAll() []per.IDataItem {
	// this needs to be implemented
	return []per.IDataItem{}
}

func (handler *SQLLiteTableHandler) Create(data per.IDataItem) bool {
	// this needs to be implemented
	return false
}

func (handler *SQLLiteTableHandler) Read(data per.IDataItem) per.IDataItem {
	// this needs to be implemented
	return nil //per.IDataItem{}
}

func (handler *SQLLiteTableHandler) Update(data per.IDataItem) bool {
	// this needs to be implemented
	return false
}


func (handler *SQLLiteTableHandler) Delete(data per.IDataItem) bool {
	// this needs to be implemented
	return false
}

// End IStorage Handler 

// SQL LIte Execution Functions

// This is to be overwritten
func (handler *SQLLiteTableHandler) ParseRows(rows *sql.Rows) []per.IDataItem {
	return []per.IDataItem{}
}

// These can be used as is
func (handler *SQLLiteTableHandler) ExecuteQuery(query string) int64 {
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteQuery - Prepare",perr)
		return -1
	}
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

func (handler *SQLLiteTableHandler) ExecuteQueryWithDatay(query string,params ...interface{}) int64 {
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteQueryWithDatay - Prepare",perr)
		return -1
	}
	res, err := statement.Exec(params...)
	if err ==  nil {
		rowsaff, rerr := res.RowsAffected()
		if rerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteQuery - RowsAffected Error",rerr)
			return -1
		}
		return rowsaff
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteQueryWithDatay",err)
		return -1
	}
}


func (handler *SQLLiteTableHandler) ExecuteInsertQuery(query string,params ...interface{}) int64 {
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery - Prepare",perr)
		return -1
	}
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

func (handler *SQLLiteTableHandler) ExecuteResult(query string) []per.IDataItem {
	empty := []per.IDataItem{}
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteResult - Prepare",perr)
		return empty
	}
	rows, err := statement.Query()
	if err ==  nil {
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResult",err)
		return empty
	}
}


func (handler *SQLLiteTableHandler) ExecuteResultWithData(query string, params ...interface{}) []per.IDataItem {
	empty := []per.IDataItem{}
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData - Prepare",perr)
		return empty
	}
	rows, err := statement.Query(params...)
	if err ==  nil {
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData",err)
		return empty
	}
}
