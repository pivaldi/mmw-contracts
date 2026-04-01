package todo

import (
	"context"
	"fmt"

	todov1 "github.com/pivaldi/mmw-contracts/gen/go/todo/v1"
)

// TodoService is the public interface exposed by the todo module to other modules.
// It uses proto-generated request/response types so callers work with the same
// types they already use for HTTP, without depending on the todo module's internals.
type TodoService interface {
	CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error)
	GetTodo(ctx context.Context, req *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error)
	UpdateTodo(ctx context.Context, req *todov1.UpdateTodoRequest) (*todov1.UpdateTodoResponse, error)
	CompleteTodo(ctx context.Context, req *todov1.CompleteTodoRequest) (*todov1.CompleteTodoResponse, error)
	ReopenTodo(ctx context.Context, req *todov1.ReopenTodoRequest) (*todov1.ReopenTodoResponse, error)
	DeleteTodo(ctx context.Context, req *todov1.DeleteTodoRequest) (*todov1.DeleteTodoResponse, error)
	ListTodos(ctx context.Context, req *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error)
}

// compile-time assertion
var _ TodoService = (*InprocClient)(nil)

// InprocClient is a thin wrapper that accepts any TodoService implementation.
// Pass the real service adapter for in-process calls, or a mock for testing.
type InprocClient struct {
	server TodoService
}

func NewInprocClient(server TodoService) *InprocClient {
	return &InprocClient{server: server}
}

func (c *InprocClient) CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	resp, err := c.server.CreateTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) GetTodo(ctx context.Context, req *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	resp, err := c.server.GetTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) UpdateTodo(ctx context.Context, req *todov1.UpdateTodoRequest) (*todov1.UpdateTodoResponse, error) {
	resp, err := c.server.UpdateTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) CompleteTodo(
	ctx context.Context, req *todov1.CompleteTodoRequest,
) (*todov1.CompleteTodoResponse, error) {
	resp, err := c.server.CompleteTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) ReopenTodo(
	ctx context.Context, req *todov1.ReopenTodoRequest,
) (*todov1.ReopenTodoResponse, error) {
	resp, err := c.server.ReopenTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) DeleteTodo(
	ctx context.Context, req *todov1.DeleteTodoRequest,
) (*todov1.DeleteTodoResponse, error) {
	resp, err := c.server.DeleteTodo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

func (c *InprocClient) ListTodos(
	ctx context.Context, req *todov1.ListTodosRequest,
) (*todov1.ListTodosResponse, error) {
	resp, err := c.server.ListTodos(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

// ErrTodoServiceUnavailable is returned by NoopTodoService methods.
var ErrTodoServiceUnavailable = fmt.Errorf("todo service unavailable")

// NoopTodoService is a no-op implementation of TodoService that rejects every
// request. Use it in standalone deployments where the todo service is not wired.
type NoopTodoService struct{}

// compile-time assertion
var _ TodoService = (*NoopTodoService)(nil)

func (NoopTodoService) CreateTodo(_ context.Context, _ *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) GetTodo(_ context.Context, _ *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) UpdateTodo(_ context.Context, _ *todov1.UpdateTodoRequest) (*todov1.UpdateTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) CompleteTodo(
	_ context.Context, _ *todov1.CompleteTodoRequest,
) (*todov1.CompleteTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) ReopenTodo(_ context.Context, _ *todov1.ReopenTodoRequest) (*todov1.ReopenTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) DeleteTodo(_ context.Context, _ *todov1.DeleteTodoRequest) (*todov1.DeleteTodoResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) ListTodos(_ context.Context, _ *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	return nil, ErrTodoServiceUnavailable
}

func (NoopTodoService) Health(_ context.Context) (any, error) {
	return nil, ErrTodoServiceUnavailable
}
