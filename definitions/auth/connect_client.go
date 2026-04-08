package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	authv1 "github.com/pivaldi/mmw-contracts/gen/go/auth/v1"
	"github.com/pivaldi/mmw-contracts/gen/go/auth/v1/authv1connect"
)

// PublicHTTPClient adapts authv1connect.AuthPublicServiceClient to AuthPublicService.
// Use it when the auth service runs in a separate process and must be reached
// over HTTP (Connect protocol) rather than via an in-process call.
type PublicHTTPClient struct {
	client authv1connect.AuthPublicServiceClient
}

var _ AuthPublicService = (*PublicHTTPClient)(nil)

// NewPublicHTTPClient wraps an authv1connect.AuthPublicServiceClient so it satisfies
// AuthPublicService. Construct the underlying client with:
//
//	authv1connect.NewAuthPublicServiceClient(&http.Client{}, "http://localhost:8091")
func NewPublicHTTPClient(client authv1connect.AuthPublicServiceClient) *PublicHTTPClient {
	return &PublicHTTPClient{client: client}
}

func (c *PublicHTTPClient) Register(
	ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	resp, err := c.client.Register(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("register: %w", err)
	}

	return resp.Msg, nil
}

func (c *PublicHTTPClient) Login(
	ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	resp, err := c.client.Login(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	return resp.Msg, nil
}

func (c *PublicHTTPClient) ChangePassword(
	ctx context.Context, req *authv1.ChangePasswordRequest) (*authv1.ChangePasswordResponse, error) {
	resp, err := c.client.ChangePassword(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("change password: %w", err)
	}

	return resp.Msg, nil
}

func (c *PublicHTTPClient) DeleteUser(
	ctx context.Context, req *authv1.DeleteUserRequest) (*authv1.DeleteUserResponse, error) {
	resp, err := c.client.DeleteUser(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, fmt.Errorf("delete user: %w", err)
	}

	return resp.Msg, nil
}

// PrivateHTTPClient adapts authv1connect.AuthPrivateServiceClient to AuthPrivateService.
// Use it when the auth service runs in a separate process.
type PrivateHTTPClient struct {
	client authv1connect.AuthPrivateServiceClient
}

var _ AuthPrivateService = (*PrivateHTTPClient)(nil)

// NewPrivateHTTPClient wraps an authv1connect.AuthPrivateServiceClient so it satisfies
// AuthPrivateService. Construct the underlying client with:
//
//	authv1connect.NewAuthPrivateServiceClient(&http.Client{}, "http://localhost:8091")
func NewPrivateHTTPClient(client authv1connect.AuthPrivateServiceClient) *PrivateHTTPClient {
	return &PrivateHTTPClient{client: client}
}

// ValidateToken calls the auth service's ValidateToken RPC.
// Returns an error if the token is invalid/expired or if the call fails.
func (c *PrivateHTTPClient) ValidateToken(
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
