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
	res := datastore.(*per.IPersistantStorage)
	ds.SetPersistantStorage(res)
//	ds.SetPersistantStorage(datastore)

//	ds.Parent = datastore
	return &ds
}

// Start IStorage Handler 
func (handler *SQLLiteTableHandler) GetPersistantStorage() *per.IPersistantStorage {
	return handler.Parent
}

func (handler *SQLLiteTableHandler) SetPersistantStorage(persistant *per.IPersistantStorage){
	//res := persistant.(*SQLLiteDatastore)
	handler.Parent = res
}

func (handler *SQLLiteTableHandler) CreateStructures() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("CreateStructures","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Wipe() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("Wipe","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) ReadAll() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("ReadAll","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Create(data per.IDataItem) per.IQueryResult {
	handler.Parent.GetLog().LogDebug("Create","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Read(data per.IDataItem) per.IQueryResult {
	handler.Parent.GetLog().LogDebug("Read","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult() 
}

func (handler *SQLLiteTableHandler) Update(data per.IDataItem) per.IQueryResult {

	handler.Parent.GetLog().LogDebug("Update","Returning empty failed SQL Query Result")
	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}


func (handler *SQLLiteTableHandler) Delete(data per.IDataItem) per.IQueryResult {
	handler.Parent.GetLog().LogDebug("Delete","Returning empty failed SQL Query Result")

	// this needs to be implemented
	return  NewEmptyFailedSQLLiteQueryResult()
}


func (handler *SQLLiteTableHandler) ConvertResult(data per.IQueryResult) SQLLiteQueryResult {
	// this needs to be implemented
	return  ResultToSQLLiteQueryResult(data)
}

// Conversion of results
func (handler *SQLLiteTableHandler) WipeC() SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("WipeC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.Wipe())
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadAllC() SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("ReadAllC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.ReadAll())
}

// Conversion of results
func (handler *SQLLiteTableHandler) CreateC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("CreateC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.Create(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("ReadC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.Read(data))
}
// Conversion of results
func (handler *SQLLiteTableHandler) UpdateC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("UpdateC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.Update(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) DeleteC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.GetLog().LogDebug("DeleteC","Converting IQuery Result to SQLLiteQueryResult")
	return  handler.ConvertResult(handler.Delete(data))
}

// End IStorage Handler 

// SQL LIte Execution Functions

// This is to be overwritten
func (handler *SQLLiteTableHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	handler.Parent.GetLog().LogDebug("ParseRows","Returing empty results - was this function replaced")
	return NewDataQueryResult(false,[]per.IDataItem{})
}

// These can be used as is

func (handler *SQLLiteTableHandler) ExecuteQuery(query string,params ...interface{}) per.IQueryResult  {
	handler.Parent.GetLog().LogDebug("ExecuteQuery",query)
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteQuery - Prepare",perr)
		return NewRowsAffectedQueryResult(-1)
	}
	res, err := statement.Exec(params...)
	if err ==  nil {
		rowsaff, rerr := res.RowsAffected()
		if rerr !=  nil {
			handler.Parent.GetLog().LogErrorE("ExecuteQuery - RowsAffected Error",rerr)
			return NewRowsAffectedQueryResult(-1)
		}
		handler.Parent.GetLog().LogDebugf("ExecuteQuery","Number of rows affected %d",rowsaff)
		return NewRowsAffectedQueryResult(rowsaff)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteQuery",err)
		return NewRowsAffectedQueryResult(-1)
	}
}


func (handler *SQLLiteTableHandler) ExecuteInsertQuery(query string,params ...interface{}) per.IQueryResult  {
	handler.Parent.GetLog().LogDebug("ExecuteInsertQuery",query)
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
		handler.Parent.GetLog().LogDebugf("ExecuteInsertQuery","Last Insert Id %d",lastid)
		return NewRowsAffectedQueryResult(lastid) 
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteInsertQuery",err)
		return NewRowsAffectedQueryResult(-1) 
	}
}

func (handler *SQLLiteTableHandler) ExecuteResult(query string, params ...interface{}) per.IQueryResult  {
	empty := []per.IDataItem{}
	handler.Parent.GetLog().LogDebug("ExecuteResult",query)
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr !=  nil {
		handler.Parent.GetLog().LogErrorE("ExecuteResult - Prepare",perr)
		return NewDataQueryResult(false,empty)
	}
	rows, err := statement.Query(params...)
	if err ==  nil {
		handler.Parent.GetLog().LogDebug("ExecuteResult","Resulted with rows to be parsed")
		return handler.ParseRows(rows)
	} else {
		handler.Parent.GetLog().LogErrorE("ExecuteResultWithData",err)
		return NewDataQueryResult(false,empty)
	}
}
