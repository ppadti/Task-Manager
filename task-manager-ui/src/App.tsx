import {
  Container,
  TextField,
  Button,
  List,
} from "@mui/material";
import { useEffect, useState } from "react";
import SingleTask from "./SingleTask";
import axios from "axios";

export interface Task {
  id: number;
  task: string;
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [name, setName] = useState("");

  useEffect(() => {
    fetchTask();
  }, []);

  const fetchTask = async () => {
      const response = await axios.get('http://localhost:8080/tasks')
      setTasks(response.data)
  };

  const handleAddTask = async () => {
    if (name)  { 
      setName('')
      await axios
        .post('http://localhost:8080/tasks', {
          task: name,
        })
        .then((response: any) => {
          console.log(response)
          fetchTask()
          // Handle data
        })
        .catch((error: any) => {
          console.log(error)
        })}
  };

  const handleUpdate = async (taskId: number, newTask: string) => {
     await axios
      .put(`http://localhost:8080/tasks/${taskId}`, { task: newTask })
      .then((response) => {
        console.log(response)
        fetchTask()
      })
      .catch((error) => {
        console.log(error)
      })
 
  };

  const handleDelete = async (taskId: number) => {
      await axios
      .delete(`http://localhost:8080/tasks/${taskId}`)
      .then((response) => {
        console.log(response)
        fetchTask()
      })
      .catch((error) => {
        console.log(error)
      })
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
