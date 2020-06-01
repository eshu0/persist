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

func (res *SQLLiteQueryResult) QuerySucceeded() bool{
	return res.Succeeded
}

func (res *SQLLiteQueryResult) Error() error {
	return res.ErrorData

}
