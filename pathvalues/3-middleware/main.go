package main

import (
	"context"
	"fmt"
	"go-mux/common/db"
	"log/slog"
	"net/http"
	"strconv"
)

var userContextKey = "user"

func main() {
	http.Handle("/user/{userId}/view", getUserFromPath(http.HandlerFunc(viewUser)))

	slog.Info("starting web service on :8080")
	panic(http.ListenAndServe(":8080", nil))
}

func viewUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userContextKey).(db.User)

	slog.Info("loaded user")
	// show the user id
	w.Write([]byte(fmt.Sprintf("loaded user %s", user.String())))
}

func getUserFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userId")
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad user id"))
			return
		}
		user, err := db.LoadUser(userIDint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("could not load user: %s", err.Error())))
			return
		}
		newContext := context.WithValue(r.Context(), userContextKey, user)
		r = r.Clone(newContext)
		next.ServeHTTP(w, r)
	})
}
