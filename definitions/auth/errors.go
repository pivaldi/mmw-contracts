// contracts/definitions/auth/errors.go
package defauth

import authv1 "github.com/pivaldi/mmw-contracts/gen/go/auth/v1"

// ErrorCode constants for the auth module.
// These are proto-generated values — the TypeScript client uses the same
// AuthErrorCode enum from gen/ts/auth/v1/auth_pb.ts.
// Cast to platform.ErrorCode at the application layer call site.
const (
	ErrorCodeInvalidLogin       = authv1.AuthErrorCode_AUTH_ERROR_CODE_INVALID_LOGIN
	ErrorCodeInvalidPassword    = authv1.AuthErrorCode_AUTH_ERROR_CODE_INVALID_PASSWORD
	ErrorCodeInvalidCredentials = authv1.AuthErrorCode_AUTH_ERROR_CODE_INVALID_CREDENTIALS
	ErrorCodeInvalidToken       = authv1.AuthErrorCode_AUTH_ERROR_CODE_INVALID_TOKEN
	ErrorCodeUserNotFound       = authv1.AuthErrorCode_AUTH_ERROR_CODE_USER_NOT_FOUND
	ErrorCodeUserAlreadyExists  = authv1.AuthErrorCode_AUTH_ERROR_CODE_USER_ALREADY_EXISTS
)
