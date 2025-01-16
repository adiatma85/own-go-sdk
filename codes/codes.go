package codes

import (
	"math"

	"github.com/adiatma85/own-go-sdk/language"
	"github.com/adiatma85/own-go-sdk/operator"
)

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

const (
	// third party/client errors
	CodeClient = Code(iota + 1500)
	CodeClientMarshal
	CodeClientUnmarshal
	CodeClientErrorOnRequest
	CodeClientErrorOnReadBody
)

const (
	// general file I/O errors
	CodeFile = Code(iota + 1600)
	CodeFilePathOpenFailed
)

const (
	// auth errors
	CodeAuth = Code(iota + 1700)
	CodeAuthRefreshTokenExpired
	CodeAuthAccessTokenExpired
	CodeAuthFailure
	CodeAuthInvalidToken
	CodeForbidden
	CodeAuthRevokeRefreshTokenFailed
	CodeInvalidAccessToken
	CodeInvalidRefreshToken
)

const (
	// JSON encoding errors
	CodeJSONSchema = Code(iota + 1900)
	CodeJSONSchemaInvalid
	CodeJSONSchemaNotFound
	CodeJSONStructInvalid
	CodeJSONRawInvalid
	CodeJSONValidationError
	CodeJSONMarshalError
	CodeJSONUnmarshalError
)

const (
	// Reset Password Error
	CodePasswordDoesNotMatch = Code(iota + 3800)
	CodeFailedResetPassword
	CodeResetPasswordTokenExpired
	CodeEmptyEmail
	CodeInvalidEmail
	CodeSameCurrentPassword
	CodePasswordIsNotFilled
	CodeResetPasswordTokenInvalid
)

const (
	// Redis Cache Error
	CodeRedisGet = Code(iota + 3900)
	CodeRedisSetex
	CodeRedisIncrement
	CodeFailedLock
	CodeFailedReleaseLock
	CodeLockExist
	CodeCacheMarshal
	CodeCacheUnmarshal
	CodeCacheGetSimpleKey
	CodeCacheSetSimpleKey
	CodeCacheDeleteSimpleKey
	CodeCacheGetHashKey
	CodeCacheSetHashKey
	CodeCacheDeleteHashKey
	CodeCacheSetExpiration
	CodeCacheDecode
	CodeCacheLockNotAcquired
	CodeCacheInvalidCastType
	CodeCacheNotFound
	CodeRedisIncr
	CodeRedisDecr
	CodeRedisHSet
	CodeRedisHGet
	CodeRedisScan
	CodeRedisDelete
)

const (
	// Scheduler Running Error
	CodeSchedulerRunningError = Code(iota + 4000)
)

// Error messages only
var ErrorMessages = AppMessage{
	CodeInvalidValue:            ErrMsgBadRequest,
	CodeContextDeadlineExceeded: ErrMsgContextTimeout,
	CodeContextCanceled:         ErrMsgContextTimeout,
	CodeInternalServerError:     ErrMsgInternalServerError,
	CodeServerUnavailable:       ErrMsgServiceUnavailable,
	CodeNotImplemented:          ErrMsgNotImplemented,
	CodeBadRequest:              ErrMsgBadRequest,
	CodeNotFound:                ErrMsgNotFound,
	CodeConflict:                ErrMsgConflict,
	CodeUnauthorized:            ErrMsgUnauthorized,
	CodeTooManyRequest:          ErrMsgTooManyRequest,
	CodeMarshal:                 ErrMsgBadRequest,
	CodeUnmarshal:               ErrMsgBadRequest,
	CodeJSONMarshalError:        ErrMsgBadRequest,
	CodeJSONUnmarshalError:      ErrMsgBadRequest,
	CodeJSONValidationError:     ErrMsgBadRequest,

	CodeSQL:                   ErrMsgInternalServerError,
	CodeSQLInit:               ErrMsgInternalServerError,
	CodeSQLBuilder:            ErrMsgInternalServerError,
	CodeSQLTxBegin:            ErrMsgInternalServerError,
	CodeSQLTxCommit:           ErrMsgInternalServerError,
	CodeSQLTxRollback:         ErrMsgInternalServerError,
	CodeSQLTxExec:             ErrMsgInternalServerError,
	CodeSQLPrepareStmt:        ErrMsgInternalServerError,
	CodeSQLRead:               ErrMsgInternalServerError,
	CodeSQLRowScan:            ErrMsgInternalServerError,
	CodeSQLRecordDoesNotExist: ErrMsgNotFound,
	CodeSQLUniqueConstraint:   ErrMsgConflict,
	CodeSQLConflict:           ErrMsgConflict,
	CodeSQLNoRowsAffected:     ErrMsgInternalServerError,

	CodeClient:                ErrMsgInternalServerError,
	CodeClientMarshal:         ErrMsgInternalServerError,
	CodeClientUnmarshal:       ErrMsgInternalServerError,
	CodeClientErrorOnRequest:  ErrMsgBadRequest,
	CodeClientErrorOnReadBody: ErrMsgBadRequest,

	CodeAuth:                         ErrMsgUnauthorized,
	CodeAuthRefreshTokenExpired:      ErrMsgRefreshTokenExpired,
	CodeAuthAccessTokenExpired:       ErrMsgAccessTokenExpired,
	CodeAuthFailure:                  ErrMsgUnauthorized,
	CodeAuthInvalidToken:             ErrMsgInvalidToken,
	CodeForbidden:                    ErrMsgForbidden,
	CodeAuthRevokeRefreshTokenFailed: ErrMsgRevokeRefreshTokenFailed,

	// File I/O error
	CodeFile:               ErrMsgInternalServerError,
	CodeFilePathOpenFailed: ErrMsgInternalServerError,

	CodeFailedResetPassword:       ErrMsgResetPassword,
	CodePasswordDoesNotMatch:      ErrMsgPasswordDoesNotMatch,
	CodeResetPasswordTokenExpired: ErrMsgResetTokenExpired,
	CodeEmptyEmail:                ErrMsgEmptyEmail,
	CodeInvalidEmail:              ErrMsgInvalidEmail,
	CodeSameCurrentPassword:       ErrMsgSameCurrentPassword,
	CodePasswordIsNotFilled:       ErrMsgPasswordIsNotFilled,
	CodeResetPasswordTokenInvalid: ErrMsgResetTokenInvalid,

	CodeLockExist:            ErrMsgLockExist,
	CodeRedisGet:             ErrMsgInternalServerError,
	CodeRedisSetex:           ErrMsgInternalServerError,
	CodeFailedLock:           ErrMsgInternalServerError,
	CodeFailedReleaseLock:    ErrMsgInternalServerError,
	CodeCacheMarshal:         ErrMsgInternalServerError,
	CodeCacheUnmarshal:       ErrMsgInternalServerError,
	CodeCacheGetSimpleKey:    ErrMsgInternalServerError,
	CodeCacheSetSimpleKey:    ErrMsgInternalServerError,
	CodeCacheDeleteSimpleKey: ErrMsgInternalServerError,
	CodeCacheGetHashKey:      ErrMsgInternalServerError,
	CodeCacheSetHashKey:      ErrMsgInternalServerError,
	CodeCacheDeleteHashKey:   ErrMsgInternalServerError,
	CodeCacheSetExpiration:   ErrMsgInternalServerError,
	CodeCacheDecode:          ErrMsgInternalServerError,
	CodeCacheLockNotAcquired: ErrMsgInternalServerError,
	CodeCacheInvalidCastType: ErrMsgInternalServerError,
	CodeCacheNotFound:        ErrMsgInternalServerError,

	CodeSchedulerRunningError: ErrMsgInternalServerError,
}

var ApplicationMessages = AppMessage{
	// Other
	CodeAccepted: SuccessAccepted,
}

func Compile(code Code, lang string) DisplayMessage {
	if appMsg, ok := ApplicationMessages[code]; ok {
		return DisplayMessage{
			StatusCode: appMsg.StatusCode,
			Title:      operator.Ternary(lang == language.Indonesian, appMsg.TitleID, appMsg.TitleEN),
			Body:       operator.Ternary(lang == language.Indonesian, appMsg.BodyID, appMsg.BodyEN),
		}
	}

	return DisplayMessage{
		StatusCode: SuccessDefault.StatusCode,
		Title:      operator.Ternary(lang == language.Indonesian, SuccessDefault.TitleID, SuccessDefault.TitleEN),
		Body:       operator.Ternary(lang == language.Indonesian, SuccessDefault.BodyID, SuccessDefault.BodyEN),
	}
}
