package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = make(map[int]Task)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	json.NewEncoder(w).Encode(taskList)
}

func AddTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)

	Id := rand.Intn(100) + 1
	newTask.ID = Id
	tasks[newTask.ID] = newTask
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	taskId := params["id"]

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)
	id, err := strconv.Atoi(taskId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid task id")
	}
	if _, ok := tasks[id]; ok {
		tasks[id] = Task{ID: id, Name: updatedTask.Name}
		json.NewEncoder(w).Encode(tasks[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not foud")
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-TYpe", "application/json")

	params := mux.Vars(r)
	taskId := params["id"]
	id, _ := strconv.Atoi(taskId)

	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		w.WriteHeader((http.StatusOK))
		json.NewEncoder(w).Encode("Task deleted")
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("task not found")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", AddTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTasks).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	headers := handlers.AllowedHeaders([]string{"X-requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
