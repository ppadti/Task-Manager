import {
  Box,
  Button,
  ListItem,
  ListItemSecondaryAction,
  ListItemText,
  Paper,
  TextField,
} from "@mui/material";
import React, { useState } from "react";
import { Task } from "./App";

type Props = {
  task: Task;
  handleUpdate: (id: number, name: string) => void;
  handleDelete: (id: number) => void;
};

const SingleTask = ({ task, handleUpdate, handleDelete }: Props) => {
  const [editedTask, setEditedTask] = useState<string>(task.task);
  const [edit, setEdit] = useState(false);

  const handleCancel = () => {
    setEdit(false);
  };

  const handleSave = (taskId: number) => {
    setEdit(false);
    handleUpdate(taskId, editedTask);
  };

  return (
    <Paper
      elevation={3}
      sx={{
        border: "1px",
        borderRadius: "5px",
        marginBottom: "15px",
        padding: "5px",
      }}
    >
      <ListItem>
        {edit ? (
          <Box key={task.id}>
            <TextField
              sx={{ width: "330px" }}
              label="Name"
              value={editedTask}
              onChange={(e) => {
                setEditedTask(e.target.value);
              }}
            />
            <Button
              variant="contained"
              sx={{ padding: "5px", margin: "10px" }}
              onClick={() => {
                handleSave(task.id);
              }}
            >
              Save
            </Button>
            <Button variant="contained" onClick={handleCancel}>
              Cancel
            </Button>
          </Box>
        ) : (
          <>
            <ListItemText primary={task.task} />
            <ListItemSecondaryAction>
              <Button
                variant="contained"
                onClick={() => {
                  setEdit(true);
                }}
                sx={{ margin: "5px" }}
                size="small"
              >
                Edit
              </Button>
              <Button
                variant="contained"
                onClick={() => {
                  handleDelete(task.id);
                }}
                size="small"
              >
                Delete
              </Button>
            </ListItemSecondaryAction>
          </>
        )}
      </ListItem>
    </Paper>
  );
};

export default SingleTask;
