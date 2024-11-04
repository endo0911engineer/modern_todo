package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var todos = []Todo{}

// GETリクエストに対応する関数
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POSTリクエストに対応する関数
func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo.ID = uuid.New().String()
	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}

// CORS対応のためのエントリポイント関数
func handleTodos(w http.ResponseWriter, r *http.Request) {
	// CORS設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// OPTIONSメソッドの処理（プリフライトリクエストに対応）
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// メソッドによって異なる処理を実行
	switch r.Method {
	case "GET":
		getTodos(w, r)
	case "POST":
		addTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/todos", handleTodos)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
