import {
  Container,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
  Box,
  ListItemSecondaryAction,
  IconButton,
  Paper,
} from "@mui/material";
import { useEffect, useState } from "react";
import SingleTask from "./SingleTask";

export interface Task {
  id: number;
  name: string;
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [name, setName] = useState("");

  useEffect(() => {
    fetchTask();
  }, []);

  const fetchTask = async () => {
    try {
      const response = await fetch(`http://localhost:8080/tasks`);
      const data = await response.json();
      setTasks(data);
    } catch (error) {
      console.log("Enter fetchig tasks:", error);
    }
  };
  const handleAddTask = async () => {
    if (name === "") return;
    // const id = Math.random()
    try {
      const response = await fetch(`http://localhost:8080/tasks`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name }),
      });
      if (response.ok) {
        setName("");
        // setDesc('')
        fetchTask();
      } else {
        console.log("Failed to add a task", response.statusText);
      }
    } catch (error) {
      console.log("Error adding a task", error);
    }
  };

  const handleUpdate = async (taskId: number, newTask: string) => {
    try {
      const response = await fetch(`http://localhost:8080/tasks/${taskId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "appliaction/json",
        },
        body: JSON.stringify({ name: newTask }),
      });
      if (response.ok) {
        fetchTask();
      } else {
        console.error("Failed to load tasks", response.statusText);
      }
    } catch (error) {
      console.error("error updating task", error);
    }
  };

  const handleDelete = async (taskId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/tasks/${taskId}`, {
        method: "DELETE",
      });
      if (response.ok) {
        fetchTask();
      } else {
        console.error("Failed to load");
      }
    } catch (error) {
      console.error("error deleting task", error);
    }
  };

  return (
    <Container maxWidth="sm">
      <h1>Task Manager</h1>
      <div>
        <TextField
          label="Name"
          value={name}
          onChange={(e) => {
            setName(e.target.value);
          }}
          variant="outlined"
          margin="normal"
          sx={{
            padding: "2px",
          }}
        />
      </div>
      <div>
        <Button
          variant="contained"
          onClick={handleAddTask}
          sx={{
            padding: "5px",
            marginBottom: "50px",
          }}
        >
          Add task
        </Button>
      </div>
      <List>
        {tasks &&
          tasks.map((task) => (
            <SingleTask
              task={task}
              key={task.id}
              handleUpdate={handleUpdate}
              handleDelete={handleDelete}
            />
          ))}
      </List>
    </Container>
  );
}

export default App;
