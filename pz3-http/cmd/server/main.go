package main

import (
	"log"
	"net/http"

	"example.com/pz3-http/internal/api"
	"example.com/pz3-http/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Вывести все таски
	mux.HandleFunc("GET /tasks", h.ListTasks)
	// Создать таск
	mux.HandleFunc("POST /tasks", h.CreateTask)

	// Вывести конкретный таск по id
	mux.HandleFunc("GET /tasks/", h.GetTask)

	// Удалить конкретный таск по id
	mux.HandleFunc("DELETE /tasks/", h.DeleteTask)

	// Подключаем логирование
	handler := api.Logging(mux)

	addr := ":8080"
	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
