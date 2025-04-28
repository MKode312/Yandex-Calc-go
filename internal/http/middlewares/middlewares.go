package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"calculator_go/internal/utils/orchestrator/jwts"
)

func AuthorizeJWTToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("no cookie found")
			return
		}
						
		tokenString := cookie.Value
						
		tokenValue, err := jwts.VerifyJWTToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			log.Printf("error: %v", err)
			return
		}
						
		userID, err := strconv.ParseInt(tokenValue, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error: %v", err)
			return
		}
			
		ctx := context.WithValue(r.Context(), "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}