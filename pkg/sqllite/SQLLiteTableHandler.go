package SQLL

import (
	"database/sql"

	per "github.com/eshu0/persist/pkg/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SQLLiteTableHandler struct {
	per.IStorageHandler
	Parent   *SQLLiteDatastore
	Executor *SQLLightQueryExecutor
}

func NewSQLLiteTableHandler(datastore *SQLLiteDatastore) *SQLLiteTableHandler {
	ds := SQLLiteTableHandler{}
	ds.Parent = datastore
	ds.Executor = NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler
func (handler *SQLLiteTableHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *SQLLiteTableHandler) SetPersistantStorage(persistant per.IPersistantStorage) {
	res := persistant.(*SQLLiteDatastore)
	handler.Parent = res
}

func (handler *SQLLiteTableHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Wipe() per.IQueryResult {
	handler.Parent.LogDebug("Wipe", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) ReadAll() per.IQueryResult {
	handler.Parent.LogDebug("ReadAll", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Create(data per.IDataItem) per.IQueryResult {
	handler.Parent.LogDebug("Create", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Read(data per.IDataItem) per.IQueryResult {
	handler.Parent.LogDebug("Read", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Update(data per.IDataItem) per.IQueryResult {

	handler.Parent.LogDebug("Update", "Returning empty failed SQL Query Result")
	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) Delete(data per.IDataItem) per.IQueryResult {
	handler.Parent.LogDebug("Delete", "Returning empty failed SQL Query Result")

	// this needs to be implemented
	return NewEmptyFailedSQLLiteQueryResult()
}

func (handler *SQLLiteTableHandler) ConvertResult(data per.IQueryResult) SQLLiteQueryResult {
	// this needs to be implemented
	return ResultToSQLLiteQueryResult(data)
}

// Conversion of results
func (handler *SQLLiteTableHandler) WipeC() SQLLiteQueryResult {
	handler.Parent.LogDebug("WipeC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.Wipe())
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadAllC() SQLLiteQueryResult {
	handler.Parent.LogDebug("ReadAllC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.ReadAll())
}

// Conversion of results
func (handler *SQLLiteTableHandler) CreateC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.LogDebug("CreateC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.Create(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) ReadC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.LogDebug("ReadC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.Read(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) UpdateC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.LogDebug("UpdateC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.Update(data))
}

// Conversion of results
func (handler *SQLLiteTableHandler) DeleteC(data per.IDataItem) SQLLiteQueryResult {
	handler.Parent.LogDebug("DeleteC", "Converting IQuery Result to SQLLiteQueryResult")
	return handler.ConvertResult(handler.Delete(data))
}

// End IStorage Handler

// empty Parserows example
func (handler *SQLLiteTableHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	handler.Parent.LogDebug("ParseRows", "Returing empty results - was this function replaced")
	return NewDataQueryResult(false, []per.IDataItem{})
}
