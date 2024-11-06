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

// 更新リクエストに対応する関数
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// パスからIDを取得
	id := r.URL.Path[len("/todos/"):]

	// 新しいデータを取得
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Todoを更新
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Text = updatedTodo.Text
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "todo not found", http.StatusNotFound)
}

// 削除リクエストに対応する関数
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// パスからIDを取得
	id := r.URL.Path[len("/todos/"):]

	// Todoを削除
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

// CORS対応のためのエントリポイント関数
func handleTodos(w http.ResponseWriter, r *http.Request) {
	// CORS設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
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
	case "PUT":
		updateTodo(w, r)
	case "DELETE":
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodos) //特定のIDを扱うためのエンドポイント
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
