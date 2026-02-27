package main

import (
	"context"
	"time"
)

type ContextKey string

const traceIdKey ContextKey = "trace-id"

var baseContext context.Context

type App struct {
	ctx context.Context
}

func NewApp(ctx context.Context) *App {
	// ctxWithTraceID := context.WithValue(ctx, traceIdKey, "app-trace-id")
	return &App{ctx}
}

func init() {
	baseContext = context.WithValue(context.Background(), traceIdKey, "initial-trace-id")
	baseContext, _ = context.WithTimeout(baseContext, 2*time.Second)
	// defer cancelFn()
	println("init")
	println(getTraceID(baseContext))
}

func main() {
	app := NewApp(baseContext)
	app.doSomething()
	println("main", getTraceID(app.ctx))

	println("main with basectx", getTraceID(baseContext))
}

func (a App) doSomething() {
	println("doSomething")
	println(getTraceID(a.ctx))
	for {
		<-a.ctx.Done()
		break
	}
	println("got cancelled")
}

func getTraceID(ctx context.Context) string {
	if value, ok := ctx.Value(traceIdKey).(string); ok {
		return value
	}
	return ""
}
