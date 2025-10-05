package infra

import "errors"

var (
	ErrUnauthorized              = errors.New("unauthorized")
	ErrInvalidEmailOrPassword    = errors.New("invalid email or password")
	ErrInvalidPhoneOrPassword    = errors.New("invalid phone number or password")
	ErrInvalidNicknameOrPassword = errors.New("invalid nickname or password")

	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrEmailRequired       = errors.New("email field is required")
	ErrPhoneNumberRequired = errors.New("phone number field is required")
	ErrNicknameRequired    = errors.New("nickname field is required")
	ErrEmailInUse          = errors.New("this email address is already in use")
	ErrPhoneNumberInUse    = errors.New("this phone number is already in use")
	ErrNicknameInUse       = errors.New("this nickname is already in use")
	ErrUserNotExists       = errors.New("user does not exist")

	ErrTransaction   = errors.New("transaction could not start")
	ErrRollback      = errors.New("rollback failed")
	ErrCommit        = errors.New("commit failed")
	ErrJsonStringify = errors.New("failed to stringify data to JSON")
	ErrModelCreate   = errors.New("failed to create model")
	ErrQueryError    = errors.New("query execution failed")

	ErrDBConnection = errors.New("failed to establish database connection")
	ErrDBPing       = errors.New("failed to ping database")
)
