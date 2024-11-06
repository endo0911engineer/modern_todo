'use client'

import { useState, useEffect } from 'react';
import { getTodos, addTodo, updateTodo, deleteTodo } from '../lib/api';
import './globals.css';

interface Todo {
  id: string;
  text: string;
}

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState<string>('');
  const [editingTodo, setEditingTodo] = useState<Todo | null>(null);

  // 初回レンダリングのデータ取得
  useEffect(() => {
    const fetchTodos = async () => {
      const data = await getTodos();
      setTodos(data);
    };
    fetchTodos();
  }, []);

  //Todoを追加する処理
  const handleAddTodo = async () => {
    if (newTodo.trim() === '') return;
    const addedTodo = await addTodo(newTodo);
    setTodos([...todos, addedTodo]);
    setNewTodo('');
  };

  // Todoを編集する処理
  const handleUpdateTodo = async (id: string, text: string) => {
    const updatedTodo = await updateTodo(id, text);
    setTodos(todos.map((todo) => (todo.id === id ? updatedTodo : todo)));
    setEditingTodo(null);
  };

  // Todoを削除する処理
  const handleDeleteTodo = async (id: string) => {
    await deleteTodo(id);
    setTodos(todos.filter((todo) => todo.id !== id));
  };

  return (
    <div className="container">
      <h1>TODO List</h1>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            {editingTodo?.id === todo.id ? (
              <input
              type="text"
              placeholder="Add new todo"
              value={editingTodo.text}
              onChange={(e) =>
                setEditingTodo({ ...editingTodo, text: e.target.value })
              }
              />
            ) : (
              <span>{todo.text}</span>
            )}
            {editingTodo?.id === todo.id ? (
              <button onClick={() => handleUpdateTodo(todo.id, editingTodo.text)}>Save</button>
            ) : (
              <button onClick={() => setEditingTodo(todo)}>Edit</button>
            )}
            <button onClick={() => handleDeleteTodo(todo.id)}>Delete</button>
          </li>
        ))}
      </ul>
      <input
        type="text"
        value={newTodo}
        onChange={(e) => setNewTodo(e.target.value)}
      />
      <button onClick={handleAddTodo}>Add TODO</button>
    </div>
  );
}