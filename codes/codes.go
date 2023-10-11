package codes

import "math"

type Code uint32

type AppMessage map[Code]Message

type DisplayMessage struct {
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Body       string `json:"body"`
}

// * Important: For new codes, please add them to the bottom of corresponding list to avoid changing existing codes and potentially breaking existing flow

// Note: Subject to change

const NoCode Code = math.MaxUint32

const (
	// Success code
	CodeSuccess = Code(iota + 10)
	CodeAccepted
)

const (
	// common errors
	CodeInvalidValue = Code(iota + 1000)
	CodeContextDeadlineExceeded
	CodeContextCanceled
	CodeInternalServerError
	CodeServerUnavailable
	CodeNotImplemented
	CodeBadRequest
	CodeNotFound
	CodeConflict
	CodeUnauthorized
	CodeTooManyRequest
	CodeMarshal
	CodeUnmarshal
)

const (
	// SQL errors
	CodeSQL = Code(iota + 1300)
	CodeSQLInit
	CodeSQLBuilder
	CodeSQLTxBegin
	CodeSQLTxCommit
	CodeSQLTxRollback
	CodeSQLTxExec
	CodeSQLPrepareStmt
	CodeSQLRead
	CodeSQLRowScan
	CodeSQLRecordDoesNotExist
	CodeSQLUniqueConstraint
	CodeSQLConflict
	CodeSQLNoRowsAffected
)
