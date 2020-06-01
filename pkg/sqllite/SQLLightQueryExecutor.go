package SQLL

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	//sli "github.com/eshu0/simplelogger/interfaces"
	per "github.com/eshu0/persist/pkg/interfaces"
)   

type SQLLightQueryExecutor struct {
	Parent *SQLLiteDatastore
}

func NewSQLLightQueryExecutor(datastore *SQLLiteDatastore) *SQLLightQueryExecutor {
	ds := SQLLightQueryExecutor{}
	ds.Parent = datastore
	return &ds
}

// These can be used as is

func (handler *SQLLightQueryExecutor) ExecuteQuery(query string,params ...interface{}) per.IQueryResult  {
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


func (handler *SQLLightQueryExecutor) ExecuteInsertQuery(query string,params ...interface{}) per.IQueryResult  {
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

func (handler *SQLLightQueryExecutor) ExecuteResult(query string, params ...interface{}) per.IQueryResult  {
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
