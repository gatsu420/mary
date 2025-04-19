package ctxvalue

import "context"

type ctxKey struct{}

var userCtxKey = ctxKey{}

type User struct {
	UserID string
}

func SetUser(ctx context.Context, user User) context.Context {
	ctx = context.WithValue(ctx, userCtxKey, user)
	return ctx
}

func GetUser(ctx context.Context) User {
	u := ctx.Value(userCtxKey).(User)
	return u
}
