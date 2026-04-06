# protoc-gen-go-contracts

A `buf`/`protoc` plugin that generates the boilerplate definitions layer from proto service definitions.

## What it generates

For each `.proto` file, up to three files are emitted under `definitions/{domain}/`:

| File | Content |
|------|---------|
| `contract.go` | `{Service}` interface, `InprocClient`, `Noop{Service}` |
| `errors.go` | `const ErrorCodeXxx` aliases for every `*ErrorCode` enum value |
| `events.go` | `TopicXxx` constants, `AllEvents` slice, `type XxxEvent` aliases |

The domain name is derived from the proto package: `todo.v1` → `todo`, `auth.v1` → `auth`.

## Build and install

```bash
cd contracts/
go build -o $(go env GOPATH)/bin/protoc-gen-go-contracts ./cmd/protoc-gen-go-contracts
```

## Usage

The plugin is already wired in `buf.gen.yaml`. Run it alongside the other plugins:

```bash
buf generate
```

## Topic routing-key constants

Topic constants are generated from the `(options.v1.topic)` custom option defined in
`proto/options/v1/options.proto`. Annotate each `*Event` message with its routing key:

```proto
import "options/v1/options.proto";

message UserDeletedEvent {
  option (options.v1.topic) = "auth.user.deleted.v1";
  string user_id = 1;
  google.protobuf.Timestamp deleted_at = 2;
}
```

The plugin reads this option via the raw protobuf bytes (no generated code import needed)
and emits:

```go
const TopicUserDeleted = "auth.user.deleted.v1"

var AllEvents = []string{TopicUserDeleted, ...}

type UserDeletedEvent = authv1.UserDeletedEvent
```

Messages without the option still get their type alias but are excluded from `AllEvents`.

## Known limitation: services with custom Go return types (e.g. `auth`)

The plugin maps every RPC method signature directly from the proto definition:

```go
// generated — proto-accurate but not what AuthService needs
ValidateToken(ctx context.Context, req *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error)

// hand-written — domain-idiomatic
ValidateToken(ctx context.Context, token string) (uuid.UUID, error)
```

When the domain contract diverges from the raw proto types (custom return types, unwrapped
primitives, etc.), **keep that service's `contract.go` hand-written** and do not include
its `.proto` file in the plugin's generation scope.

A future improvement would be to annotate RPCs with a custom option
(e.g. `option (options.v1.go_return_type) = "github.com/google/uuid.UUID"`) so the plugin
can honour them. That is not implemented yet.
