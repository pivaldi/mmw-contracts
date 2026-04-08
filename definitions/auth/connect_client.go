package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	authv1 "github.com/pivaldi/mmw-contracts/gen/go/auth/v1"
	"github.com/pivaldi/mmw-contracts/gen/go/auth/v1/authv1connect"
)

// HTTPClient adapts authv1connect.AuthServiceClient to AuthService.
// Use it when the auth service runs in a separate process and must be reached
// over HTTP (Connect protocol) rather than via an in-process call.
type HTTPClient struct {
	client authv1connect.AuthServiceClient
}

// compile-time assertion
var _ AuthService = (*HTTPClient)(nil)

// NewHTTPClient wraps an authv1connect.AuthServiceClient so it satisfies
// AuthService. Construct the underlying client with:
//
//	authv1connect.NewAuthServiceClient(&http.Client{}, "http://localhost:8091")
func NewHTTPClient(client authv1connect.AuthServiceClient) *HTTPClient {
	return &HTTPClient{client: client}
}

func (c *HTTPClient) Register(
	ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	resp, err := c.client.Register(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("register: %w", err)
	}

	return resp.Msg, nil
}

func (c *HTTPClient) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	resp, err := c.client.Login(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	return resp.Msg, nil
}

// ValidateToken calls the auth service's ValidateToken RPC.
// Returns an error if the token is invalid/expired or if the call fails.
func (c *HTTPClient) ValidateToken(
	ctx context.Context, req *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	resp, err := c.client.ValidateToken(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("validate token: %w", err)
	}
	if !resp.Msg.GetIsValid() {
		return nil, fmt.Errorf("token is invalid or expired")
	}

	return resp.Msg, nil
}

func (c *HTTPClient) ChangePassword(
	ctx context.Context, req *authv1.ChangePasswordRequest) (*authv1.ChangePasswordResponse, error) {
	resp, err := c.client.ChangePassword(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("change password: %w", err)
	}

	return resp.Msg, nil
}

func (c *HTTPClient) DeleteUser(ctx context.Context, req *authv1.DeleteUserRequest) (*authv1.DeleteUserResponse, error) {
	resp, err := c.client.DeleteUser(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("delete user: %w", err)
	}

	return resp.Msg, nil
}
