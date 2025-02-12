package grpc

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpcBook "github.com/serj213/bookService/internal/grpc/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	log *slog.Logger
	grpcServer *grpc.Server
	port int
}

func New(
	log *slog.Logger,
	bookService grpcBook.Book,
	port int,
) *App{


	recoveryOpt := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			log.Error("recovery from panic", slog.Any("panic", p))

			return status.Error(codes.Internal, "internal error")
		}),
	}

	loggingOpt := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpt...),
		logging.UnaryServerInterceptor(loggingInterceptor(log), loggingOpt...),
	))
 	_ = grpcServer

	return nil
}

func loggingInterceptor(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(level), msg, fields)
	})
}