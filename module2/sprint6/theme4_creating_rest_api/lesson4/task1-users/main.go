package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var users = []string{
	"Максим",
	"Жанет",
}

func main() {
	r := chi.NewRouter()
	r.Get("/users", getUsers)
	r.Get("/users/{id}", getUser)
	r.Post("/users", postUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

	fmt.Println("Сервер успешно запущен!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil || len(users) <= id || id < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := users[id]

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}

func postUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Неправильное тело запроса", http.StatusBadRequest)
		return
	}

	newUserId := len(users)
	users = append(users, name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(newUserId)))
}
