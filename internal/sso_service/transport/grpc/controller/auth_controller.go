package grpccontroller

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	ssoentity "social-network/internal/sso_service/entity"
	ssoauthpb "social-network/protos/gen/sso_service"
)

type SSOAuthUseCase interface {
	SignUp(ctx context.Context, user *ssoentity.User) (userID int, err error)
	SignIn(ctx context.Context, user *ssoentity.User, ss *ssoentity.Session, token *ssoentity.JWTToken) error
	SignOut(ctx context.Context, userID int, ss *ssoentity.Session, token *ssoentity.JWTToken) error
	RefreshToken(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error
}

// SSOAuthService controller
type SSOAuthController struct {
	usecase SSOAuthUseCase
	ssoauthpb.UnimplementedSSOAuthServiceServer
}

func RegisterSSOAuth(gRPCServer *grpc.Server, authsrv SSOAuthUseCase) {
	ssoauthpb.RegisterSSOAuthServiceServer(gRPCServer, &SSOAuthController{usecase: authsrv})
}

// New create a new SSOAuthService, return *SSOAuthService
func NewSSOAuthController(aus SSOAuthUseCase) *SSOAuthController {
	return &SSOAuthController{usecase: aus}
}

// SignUp
func (a *SSOAuthController) SignUp(clstream grpc.ClientStreamingServer[ssoauthpb.SignUpRequest, ssoauthpb.SignUpResponse]) error {
	slog.Error("SignUp: Implement me")

	return nil
}

// SignIn
func (a *SSOAuthController) SignIn(
	ctx context.Context,
	in *ssoauthpb.SignInRequest,
) (*ssoauthpb.SignInResponse, error) {

	slog.Error("SignIn: Implement me")

	return nil, nil
}

// SignOut
func (a *SSOAuthController) SignOut(
	ctx context.Context,
	in *ssoauthpb.SignOutRequest,
) (*ssoauthpb.SignOutResponse, error) {

	slog.Error("SignOut: Implement me")

	return nil, nil
}

// RefreshToken
func (a *SSOAuthController) RefreshToken(
	ctx context.Context,
	in *ssoauthpb.RefreshTokenRequest,
) (*ssoauthpb.RefreshTokenResponse, error) {

	slog.Error("SignOut: Implement me")

	return nil, nil
}
