package model

import "context"

type Context struct {
	base context.Context
}

func NewContext(options ...CtxOption) *Context {
	ctx := &Context{
		base: context.Background(),
	}

	for _, opt := range options {
		opt(ctx)
	}

	return ctx
}

type CtxOption func(ctx *Context)

func WithBaseCtx(base context.Context) CtxOption {
	return func(ctx *Context) {
		ctx.base = base
	}
}
