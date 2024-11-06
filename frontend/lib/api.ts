// Todoを取得するリクエスト
export const getTodos = async () => {
    const response = await fetch('http://localhost:8080/todos');
    return response.json();
};
  
// Todoに追加するリクエスト
export const addTodo = async (text: string) => {
    const response = await fetch('http://localhost:8080/todos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text }),
    });
    return response.json();
};

// Todoを編集するリクエスト
export const updateTodo = async (id: string, text: string) => {
    const response = await fetch(`http://localhost:8080/todos/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text }),
    });
    return response.json();
};

// Todoを削除するリクエスト
export const deleteTodo = async (id: string) => {
    await fetch(`http://localhost:8080/todos/${id}`, {
        method: 'DELETE',
    });
};
