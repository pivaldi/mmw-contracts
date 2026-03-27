package defauth

import authv1 "github.com/pivaldi/mmw-contracts/gen/go/auth/v1"

// The Topic Name (Routing Key)
const (
	TopicUserRegistered = "auth.user.registered.v1"
	TopicUserDeleted    = "auth.user.deleted.v1"
	//nolint:gosec // Event type constant, not a credential
	TopicPasswordChanged = "auth.user.password_changed.v1"
	TopicUserLoggedIn    = "auth.user.logged_in.v1"
)

var AllEvents = []string{
	TopicUserRegistered,
	TopicUserDeleted,
	TopicPasswordChanged,
	TopicUserLoggedIn,
}

// UserDeletedEvent is the proto-generated event type for user deletion.
type UserDeletedEvent = authv1.UserDeletedEvent
