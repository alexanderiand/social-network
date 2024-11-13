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
	SignIn(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error
	SignOut(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error 
	RefreshToken(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error 
}

// SSOAuthService controller
type SSOAuthService struct {
	usecase SSOAuthUseCase
	ssoauthpb.UnimplementedSSOAuthServiceServer
}

// New create a new SSOAuthService instance, return *SSOAuthService
func NewSSOAuthService(uc SSOAuthUseCase) *SSOAuthService {
	return &SSOAuthService{
		usecase: uc,
	}
}

func RegisterSSOAuth(gRPCServer *grpc.Server, authsrv SSOAuthUseCase) {
	ssoauthpb.RegisterSSOAuthServiceServer(gRPCServer, &SSOAuthService{usecase: authsrv})
}

// SignUp
func (a *SSOAuthService) SignUp(clstream grpc.ClientStreamingServer[ssoauthpb.SignUpRequest, ssoauthpb.SignUpResponse]) error {
	slog.Error("SignUp: Implement me")

	return nil
}

// SignIn
func (a *SSOAuthService) SignIn(
	ctx context.Context,
	in *ssoauthpb.SignInRequest,
) (*ssoauthpb.SignInResponse, error) {

	slog.Error("SignIn: Implement me")

	return nil, nil
}

// SignOut
func (a *SSOAuthService) SignOut(
	ctx context.Context,
	in *ssoauthpb.SignOutRequest,
) (*ssoauthpb.SignOutResponse, error) {

	slog.Error("SignOut: Implement me")

	return nil, nil
}

// RefreshToken
func (a *SSOAuthService) RefreshToken(
	ctx context.Context,
	in *ssoauthpb.RefreshTokenRequest,
) (*ssoauthpb.RefreshTokenResponse, error) {

	slog.Error("SignOut: Implement me")

	return nil, nil
}