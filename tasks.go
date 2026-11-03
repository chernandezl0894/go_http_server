package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type TaskHandler struct{}

type taskResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type taskErrorResponse struct {
	Message string `json:"message"`
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (handle *TaskHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskResponse{
		ID:    "01-task",
		Title: "Task number 01",
	})
	w.WriteHeader(http.StatusOK)
}

func (handle *TaskHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	if err := json.NewEncoder(w).Encode(taskResponse{ID: "01-task", Title: "Task number 01"}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.Copy(w, &buf)
}
