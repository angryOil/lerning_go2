package contextWithValue

import "context"

func contextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, "user", user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value("user").(string)
	return user, ok
}
