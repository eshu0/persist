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

func (handler *SQLLiteTableHandler) CreateStructures() per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Wipe() per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) ReadAll() per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Create(data per.IDataItem) per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Read(data per.IDataItem) per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult() 
}

func (handler *SQLLiteTableHandler) Update(data per.IDataItem) per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}


func (handler *SQLLiteTableHandler) Delete(data per.IDataItem) per.IQueryResult {
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}


func (handler *SQLLiteTableHandler) ConvertResult(data per.IQueryResult) SQLLiteQueryResult {
	// this needs to be implemented
	return  ResultToSQLLiteQueryResult(data)
}

// Conversion of results
func (handler *SQLLiteTableHandler) WipeC() SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Wipe())
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadAllC() SQLLiteQueryResult {
	return  handler.ConvertResult(handler.ReadAll())
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadC(data per.IDataItem) SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Read(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) CreateC(data per.IDataItem) SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Create(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadC(data per.IDataItem) SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Create(data))
}
// Conversion of results
func (handler *SQLLiteTableHandler) UpdateC(data per.IDataItem) SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Create(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) DeleteC(data per.IDataItem) SQLLiteQueryResult {
	return  handler.ConvertResult(handler.Create(data))
}

// End IStorage Handler 

// SQL LIte Execution Functions

// This is to be overwritten
func (handler *SQLLiteTableHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	return NewDataQueryResult(false,[]per.IDataItem{})
}

// These can be used as is

func (handler *SQLLiteTableHandler) ExecuteQuery(query string,params ...interface{}) per.IQueryResult  {
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteQueryWithDatay - Prepare",perr)
		return NewRowsAffectedQueryResult(-1)
	}
	res, err := statement.Exec(params...)
	if err ==  nil {
		rowsaff, rerr := res.RowsAffected()
		if rerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteQuery - RowsAffected Error",rerr)
			return NewRowsAffectedQueryResult(-1)
		}
		return NewRowsAffectedQueryResult(rowsaff)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteQueryWithDatay",err)
		return NewRowsAffectedQueryResult(-1)
	}
}


func (handler *SQLLiteTableHandler) ExecuteInsertQuery(query string,params ...interface{}) per.IQueryResult  {
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery - Prepare",perr)
		return NewRowsAffectedQueryResult(-1) 
	}
	res, err := statement.Exec(params...)
	if err ==  nil {
		lastid, lerr := res.LastInsertId()
		if lerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery - LastInsertId",lerr)
			return NewRowsAffectedQueryResult(-1) 
		}
		return NewRowsAffectedQueryResult(lastid) 
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery",err)
		return NewRowsAffectedQueryResult(-1) 
	}
}

func (handler *SQLLiteTableHandler) ExecuteResult(query string, params ...interface{}) per.IQueryResult  {
	empty := []per.IDataItem{}
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData - Prepare",perr)
		return NewDataQueryResult(false,empty)
	}
	rows, err := statement.Query(params...)
	if err ==  nil {
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData",err)
		return NewDataQueryResult(false,empty)
	}
}
