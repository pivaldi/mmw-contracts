package auth

import authv1 "github.com/pivaldi/mmw-contracts/gen/go/auth/v1"

// User is a convenience alias for the proto-generated User message,
// re-exported so callers can use auth.User without importing authv1 directly.
type User = authv1.User
