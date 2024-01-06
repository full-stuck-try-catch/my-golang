import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");

  const fetchTodos = async () => {
    const res = await axios.get("http://localhost:8080/api/todos");
    setTodos(res.data);
  };

  const createTodo = async () => {
    await axios.post("http://localhost:8080/api/todos", { title });
    setTitle("");
    fetchTodos();
  };

  const updateTodo = async (todo) => {
    await axios.put(`http://localhost:8080/api/todos/${todo.ID}`, {
      ...todo,
      completed: !todo.completed,
    });
    fetchTodos();
  };

  const deleteTodo = async (id) => {
    await axios.delete(`http://localhost:8080/api/todos/${id}`);
    fetchTodos();
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h1>Todo List</h1>
      <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder="Add todo" />
      <button onClick={createTodo}>Add</button>
      <ul>
        {todos.map((todo) => (
          <li key={todo.ID}>
            <span
              style={{ textDecoration: todo.completed ? "line-through" : "none", cursor: "pointer" }}
              onClick={() => updateTodo(todo)}
            >
              {todo.title}
            </span>
            <button onClick={() => deleteTodo(todo.ID)}>❌</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
