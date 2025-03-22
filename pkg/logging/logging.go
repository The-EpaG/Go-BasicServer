package logging

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ctxKey string

const traceIDKey ctxKey = "X-Trace-Id"

var (
	logger *log.Logger
	once   sync.Once
)

// Initialize logger
func initLogger() {
	once.Do(func() {
		logger = log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
	})
}

func GetTraceKey() ctxKey {
	return traceIDKey
}

// Generate a new request ID
func GenerateRequestID() string {
	return uuid.New().String()
}

// Inject the request ID into context
func InjectRequestID(ctx context.Context, traceID string) context.Context {
	initLogger()
	if traceID == "" {
		traceID = GenerateRequestID()
	}
	return context.WithValue(ctx, traceIDKey, traceID)
}

// Retrieve the request ID from context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(GetTraceKey()).(string); ok {
		return id
	}
	return "unknown"
}

// Log wrapper to include request ID
func Info(ctx context.Context, msg string) {
	initLogger()
	requestID := GetRequestID(ctx)
	logger.Printf("[INFO] [trace_id=%s] %s", requestID, msg)
}

func Error(ctx context.Context, msg string) {
	initLogger()
	requestID := GetRequestID(ctx)
	logger.Printf("[ERROR] [trace_id=%s] %s", requestID, msg)
}
