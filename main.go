package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// NOTE: Creamos el servidor y registramos el handler en su mux.
	srv := New(":8080")

	tasksHandler := NewTaskHandler()

	srv.Handle("GET /api/v1/tasks", http.HandlerFunc(tasksHandler.HandleGet))
	srv.Handle("POST /api/v1/tasks", http.HandlerFunc(tasksHandler.HandlePost))

	if err := srv.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
