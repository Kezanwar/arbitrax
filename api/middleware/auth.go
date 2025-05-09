package middleware

import (
	"Arbitrax/constants"
	user_repo "Arbitrax/repositories/user"

	"Arbitrax/output"
	"Arbitrax/services/jwt"
	"Arbitrax/services/uuid"
	"context"
	"net/http"
)

// AuthMiddleware returns a Middleware that uses the injected user.Repository
func AuthMiddleware(repo user_repo.Repository) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(constants.AUTH_TOKEN_HEADER)
			if len(token) == 0 {
				r.Body.Close()
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token required"})
				return
			}

			parsed, err := jwt.Parse(token)
			if err != nil {
				r.Body.Close()
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: err.Error()})
				return
			}

			id, ok := parsed["uuid"].(string)
			if !ok || !uuid.Validate(id) {
				r.Body.Close()
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth token failed"})
				return
			}

			usr, err := repo.GetByUUID(r.Context(), id)
			if err != nil {
				r.Body.Close()
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth error"})
				return
			}
			if usr == nil {
				r.Body.Close()
				output.WriteJson(w, r, http.StatusForbidden, output.MessageResponse{Message: "Auth failed"})
				return
			}

			ctx := context.WithValue(r.Context(), constants.USER_CTX, usr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
