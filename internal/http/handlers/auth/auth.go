package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=UserInteractor
type UserInteractor interface {
	RegisterUser(ctx context.Context, uname, pswrd string) error
	LoginUser(ctx context.Context, uname, pswrd string) (string, error)
}

func RegisterUserHandler(ctx context.Context, userInteractor UserInteractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("error: %v", err)
			return
		}

		// Проверка на наличие полей username и password
		if req.Login == "" || req.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		if err := userInteractor.RegisterUser(ctx, req.Login, req.Password); err != nil {
			http.Error(w, "This login is already registered", http.StatusBadRequest)
			log.Printf("error: %v", err)
			return
		}

		w.Write([]byte("You have successfully registered"))

		log.Print("success RegisterUserHandler")
		w.WriteHeader(http.StatusCreated)
	}
}

func LoginUserHandler(ctx context.Context, userInteractor UserInteractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        var req Request
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            log.Printf("error: %v", err)
            return
        }

		// Проверка на наличие полей username и password
		if req.Login == "" || req.Password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

        token, err := userInteractor.LoginUser(ctx, req.Login, req.Password)
        if err != nil {
            http.Error(w, "Invalid login or password", http.StatusUnauthorized)
            log.Printf("error: %v", err)
            return
        }

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		})

        // Возвращаем токен в ответе вместо установки cookie.
        response := map[string]string{"token": token}
        json.NewEncoder(w).Encode(response)

        log.Print("success LoginUserHandler")
    }
}