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

const RequestID = "RequestID"

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
	reqID := reqUUID.String()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	if ids := md[RequestID]; len(ids) > 0 {
		reqID = ids[0]
	} else {
		md.Set(RequestID, reqID)
	}

	ctx = metadata.NewIncomingContext(ctx, md)

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
		reqTime.Format(time.DateTime),
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
func RequestIDStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {

	reqUUID, err := uuid.NewRandom()
	if err != nil {
		slog.Warn(err.Error())
		return err
	}

	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		md = metadata.New(nil)
	}
	reqID := reqUUID.String()

	if ids := md[RequestID]; len(ids) > 0 {
		reqID = ids[0]
	} else {
		md.Set(RequestID, reqID)
	}

	_ = metadata.NewIncomingContext(ss.Context(), md)

	return nil
}

// RequestIDSteamInterceptor

// LoggingStreamInterceptor

// AuthSteamInterceptor
