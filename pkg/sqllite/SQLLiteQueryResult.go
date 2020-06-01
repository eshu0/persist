package SQLL

import (
	//"database/sql"

	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)   

type SQLLiteQueryResult struct {
	per.IQueryResult

	Results []per.IDataItem 
	Result per.IDataItem
	
	RowsAffected int64
	LastInsertId int64
	
	Succeeded bool
	ErrorData error
}


func ResultToSQLLiteQueryResult(data IQuery) SQLLiteQueryResult {
	res := data.(SQLLiteQueryResult)
	return res
}


func NewDataQueryResult(succes bool, items []per.IDataItem) SQLLiteQueryResult {
	res := SQLLiteQueryResult{}
	res.Succeeded = succes
	res.Results = items
	return res
}

func NewRowsAffectedQueryResult(val int64) SQLLiteQueryResult {
	res := SQLLiteQueryResult{}
	res.Succeeded = val >= 0
	res.RowsAffected = val
	return res
}

func NewInsertRowsQueryResult(val int64) SQLLiteQueryResult {
	res := SQLLiteQueryResult{}
	res.Succeeded = val >= 0
	res.LastInsertId = val
	return res
}

func NewEmptyFailedSQLLiteQueryResult() SQLLiteQueryResult {
	res := SQLLiteQueryResult{}
	res.Succeeded =false
	return res
}

func NewEmptySucceedSQLLiteQueryResult() SQLLiteQueryResult {
	res := SQLLiteQueryResult{}
	res.Succeeded =true
	return res
}

func (res SQLLiteQueryResult) QuerySucceeded() bool{
	return res.Succeeded
}

func (res SQLLiteQueryResult) Error() error {
	return res.ErrorData

}
