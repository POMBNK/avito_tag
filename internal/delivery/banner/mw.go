package banner

import (
	"context"
	"net/http"
)

const (
	adminToken = "admin_token"
	userToken  = "user_token"
)

func SimpleMiddleware(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		// Middleware logic before calling the handler function
		accessToken := r.Header.Get("token")
		scopes := ctx.Value(TokenScopes).([]string)

		for _, scope := range scopes {
			if scope == accessToken || accessToken == adminToken {
				response, err = f(ctx, w, r, request)

				return response, err
			}
		}

		// Call the original handler function

		// Middleware logic after calling the handler function

		return response, err
	}
}
