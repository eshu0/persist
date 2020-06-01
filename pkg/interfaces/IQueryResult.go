package persist

type IQueryResult interface {
	QuerySucceeded() bool
	Error() error
}