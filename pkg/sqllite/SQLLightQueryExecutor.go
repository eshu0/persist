package SQLL

import (
	"database/sql"

	per "github.com/eshu0/persist/pkg/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type ParseRows func(*sql.Rows) per.IQueryResult

type SQLLightQueryExecutor struct {
	Parent *SQLLiteDatastore
}

func NewSQLLightQueryExecutor(datastore *SQLLiteDatastore) *SQLLightQueryExecutor {
	ds := SQLLightQueryExecutor{}
	ds.Parent = datastore
	return &ds
}

// These can be used as is

func (handler *SQLLightQueryExecutor) ExecuteQuery(query string, params ...interface{}) per.IQueryResult {
	handler.Parent.LogDebug("ExecuteQuery", query)
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr != nil {
		handler.Parent.LogErrorE("ExecuteQuery - Prepare", perr)
		return NewRowsAffectedQueryResult(-1)
	}
	res, err := statement.Exec(params...)
	if err == nil {
		rowsaff, rerr := res.RowsAffected()
		if rerr != nil {
			handler.Parent.LogErrorE("ExecuteQuery - RowsAffected Error", rerr)
			return NewRowsAffectedQueryResult(-1)
		}
		handler.Parent.LogDebugf("ExecuteQuery", "Number of rows affected %d", rowsaff)
		return NewRowsAffectedQueryResult(rowsaff)
	} else {
		handler.Parent.LogErrorE("ExecuteQuery", err)
		return NewRowsAffectedQueryResult(-1)
	}
}

func (handler *SQLLightQueryExecutor) ExecuteInsertQuery(query string, params ...interface{}) per.IQueryResult {
	handler.Parent.LogDebug("ExecuteInsertQuery", query)
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr != nil {
		handler.Parent.LogErrorE("ExecuteInsertQuery - Prepare", perr)
		return NewInsertRowsQueryResult(-1)
	}
	res, err := statement.Exec(params...)
	if err == nil {
		lastid, lerr := res.LastInsertId()
		if lerr != nil {
			handler.Parent.LogErrorE("ExecuteInsertQuery - LastInsertId", lerr)
			return NewInsertRowsQueryResult(-1)
		}
		handler.Parent.LogDebugf("ExecuteInsertQuery", "Last Insert Id %d", lastid)
		return NewInsertRowsQueryResult(lastid)
	} else {
		handler.Parent.LogErrorE("ExecuteInsertQuery", err)
		return NewInsertRowsQueryResult(-1)
	}
}

func (handler *SQLLightQueryExecutor) ExecuteResult(query string, parser ParseRows, params ...interface{}) per.IQueryResult {
	empty := []per.IDataItem{}
	handler.Parent.LogDebug("ExecuteResult", query)
	statement, perr := handler.Parent.GetDatabase().Prepare(query)
	if perr != nil {
		handler.Parent.LogErrorE("ExecuteResult - Prepare", perr)
		return NewDataQueryResult(false, empty)
	}
	rows, err := statement.Query(params...)
	if err == nil {
		handler.Parent.LogDebug("ExecuteResult", "Resulted with rows to be parsed")
		return parser(rows)
	} else {
		handler.Parent.LogErrorE("ExecuteResultWithData", err)
		return NewDataQueryResult(false, empty)
	}
}
