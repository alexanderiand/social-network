package grpcmiddleware

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var RequestID = "RequestID"

// unary interceptors

// RequestIDUnaryInterceptor added for every request unique UUID
func RequestIDUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	// generating uuid, use google uuid v4
	reqUUID, err := uuid.NewRandom()
	if err != nil {
		slog.Warn(err.Error())
		return nil, err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, RequestID, reqUUID.String())

	return handler(ctx, req)
}

// LoggingUnaryInterceptor log every request and every response
func LoggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	// log incoming req
	reqIDs := metadata.ValueFromIncomingContext(ctx, RequestID)
	reqID := ""
	if len(reqIDs) > 0 {
		reqID = reqIDs[0]
	}

	reqTime := time.Now()
	binfo := fmt.Sprintf("req_time: %s, req_id: %s, method: %s",
		time.Now().Format(reqTime.Format(time.DateTime)),
		reqID,
		info.FullMethod,
	)
	slog.Debug(binfo)

	res, err := handler(ctx, req)

	// log response
	ainfo := fmt.Sprintf("res_time: %s, req_id: %s, res_dur: %s, err: %s",
		time.Now().Format(time.DateTime),
		reqID,
		time.Since(reqTime),
		err.Error(),
	)
	slog.Debug(ainfo)

	return res, err
}

// AuthMiddlewareUnary

// stream interceptors

// RequestIDSteamInterceptor

// LoggingStreamInterceptor

// AuthSteamInterceptor
