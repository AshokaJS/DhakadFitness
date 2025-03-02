package utils

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func GetContext(r *http.Request) context.Context {
	ctx := r.Context()
	_, ok := ctx.Value("request-id").(string)

	if !ok {
		//if value is not found in the request context then we have to create new value.
		rid := uuid.New()
		ctx = context.WithValue(ctx, "request-id", rid)
	}
	return ctx
}
