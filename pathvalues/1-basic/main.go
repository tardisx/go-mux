package main

import (
	"fmt"
	"go-mux/common/db"
	"log/slog"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/user/{userId}/view", viewUser)

	slog.Info("starting web service on :8080")
	panic(http.ListenAndServe(":8080", nil))
}

func viewUser(w http.ResponseWriter, r *http.Request) {
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

	slog.Info("loaded user")
	// show the user id
	w.Write([]byte(fmt.Sprintf("loaded user %s", user)))
}
