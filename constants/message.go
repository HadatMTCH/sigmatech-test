package constants

import "errors"

const (
	PGDuplicateConstraint = "pq: duplicate key value violates unique constraint "
	ClientErrTimeout      = " (Client.Timeout exceeded while awaiting headers)"
	RedisNilValue         = "redis: nil"
	PGNoRows              = "sql: no rows in result set"
)

var (
	ErrDuplicate           = "Duplicate for data %s already in database"
	ErrHttpClient          = "Error when trying to access %s"
	ErrQueryParamsRequired = "Query Params required: %s"
	ErrIncompleteProfile   = "Please complete your profile: %s"
)

var (
	ErrNameIsRequired           = errors.New("name is required")
	ErrEmailIsRequired          = errors.New("email is required")
	ErrEmailIsNotValid          = errors.New("email is not valid")
	ErrPasswordIsRequired       = errors.New("password is required")
	ErrOldPasswordIsRequired    = errors.New("old Password is required")
	ErrNewPasswordIsRequired    = errors.New("new Password is required")
	ErrPasswordAlreadyTaken     = errors.New("new Password and Old Password should not be same")
	ErrEmailAndPasswordNotMatch = errors.New("email or Password Not Match")
	ErrKeyIsNotInvalidType      = errors.New("key is of invalid type")
	ErrTokenIsRequired          = errors.New("token is required")
	ErrTokenInvalid             = errors.New("token is invalid")
	ErrEligbleAccess            = errors.New("no right to access the API")
	ErrPasswordNotMatch         = errors.New("password doesn't match")
	ErrSaveTokenToRedis         = errors.New("error Save Token to redis")
	ErrGetTokenFromRedis        = errors.New("error Get Token from redis")
	ErrSaveOTPToRedis           = errors.New("error Save OTP to redis")
	ErrEmailNotFound            = errors.New("email not found")
	ErrTokenAlreadyExpired      = errors.New("token Already Expired")
	ErrTokenReplaced            = errors.New("please re login for next process")
	ErrBeginTransaction         = errors.New("dB Transaction Begin Failed")
	ErrDataNotFound             = errors.New("no Data")
	ErrIneligible               = errors.New("user is ineligible to Access This Feature")
)
