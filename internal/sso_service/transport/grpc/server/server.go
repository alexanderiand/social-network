package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"

	"social-network/internal/sso_service/app/config"
	grpcmiddleware "social-network/internal/sso_service/transport/grpc/middleware"
	"social-network/internal/sso_service/transport/http/rest/middleware"
	gconfig "social-network/pkg/config"
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// GRPCServer
type GRPCServer struct {
	*grpc.Server
	log *slog.Logger
	Host string
	Port string
	Timeouts time.Duration
}

// New create a new GRPCServer instance *GRPCServer
func New(
	ctx context.Context,
	gcfg *gconfig.Config, 
	lcfg *config.Config, 
	log *slog.Logger,
	) (
		*GRPCServer,
		error,
	) {
	fn := "transport.grpc.server.New"
	
	if ctx == nil || gcfg == nil || lcfg == nil || log == nil {
		return nil, fmt.Errorf("%s %w", fn, ErrNilStructPointer) 
	}

	// recovery options
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) error {
			log.Error("Recovered from panic ", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	// add interceptors, unary and stream interceptors
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...), // TODO: add other unary interceptors
		grpcmiddleware.RequestIDUnaryInterceptor,
		grpcmiddleware.LoggingUnaryInterceptor,
	), grpc.ChainStreamInterceptor(
		recovery.StreamServerInterceptor(recoveryOpts...), // TODO: add other unary interceptors
	),
	)

	// TODO: Register gRPC Services SSO Service (Auth, Permission, Information)
	// controller.AuthService.Register(grpcServer, authService)

	return &GRPCServer{
		Server: grpcServer,
		log:    log,
		Host: lcfg.GRPC.Host,
		Port: lcfg.GRPC.Port,
		Timeouts: gcfg.IdleTimeout,
	}, nil
}

// Run gRPC server
func (g *GRPCServer) Run() error {
	src := "ssosrv.transport.grpc.server.Run"
	lis, err := net.Listen("tcp", net.JoinHostPort(g.Host, g.Port))
	if  err != nil {
		return fmt.Errorf("%s %w", src, err) 
	}

	g.log.Info("gRPC server started", "on", net.JoinHostPort(g.Host, g.Port))

	if err := g.Serve(lis);  err != nil {
		return fmt.Errorf("%s %w", src, err)
	}

	return nil
}


// Stop gRPC server
func (g *GRPCServer) Stop() {
	slog.Info("stopping gRPC server", slog.String("addr", net.JoinHostPort(g.Host, g.Port)))

	g.GracefulStop()
}
