package todo

import todov1 "github.com/pivaldi/mmw-contracts/gen/go/todo/v1"

// ErrorCode constants for the todo module.
// These are proto-generated values — the TypeScript client uses the same
// TodoErrorCode enum from gen/ts/todo/v1/todo_pb.ts.
// Cast to platform.ErrorCode at the application layer call site.
const (
	ErrorCodeInvalidTitle            = todov1.TodoErrorCode_TODO_ERROR_CODE_INVALID_TITLE
	ErrorCodeInvalidDueDate          = todov1.TodoErrorCode_TODO_ERROR_CODE_INVALID_DUE_DATE
	ErrorCodeInvalidID               = todov1.TodoErrorCode_TODO_ERROR_CODE_INVALID_ID
	ErrorCodeNotFound                = todov1.TodoErrorCode_TODO_ERROR_CODE_NOT_FOUND
	ErrorCodeAlreadyExists           = todov1.TodoErrorCode_TODO_ERROR_CODE_ALREADY_EXISTS
	ErrorCodeCannotCompleteCancelled = todov1.TodoErrorCode_TODO_ERROR_CODE_CANNOT_COMPLETE_CANCELLED
	ErrorCodeCannotModifyCompleted   = todov1.TodoErrorCode_TODO_ERROR_CODE_CANNOT_MODIFY_COMPLETED
	ErrorCodeInvalidStatusTransition = todov1.TodoErrorCode_TODO_ERROR_CODE_INVALID_STATUS_TRANSITION
)
