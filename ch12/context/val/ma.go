package main

import (
	"context"
	"net/http"
)

func main() {

}

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = contextWithUser(ctx, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func contextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, "user", user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value("user").(string)
	return user, ok
}
