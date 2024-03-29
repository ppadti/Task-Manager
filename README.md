## Task Manager

This project is a task manager application that utilizes React (TypeScript) for the frontend and Go for the backend. It allows users to perform CRUD (Create, Read, Update, Delete) operations on tasks.
- Frontend built with React, TypeScript, and Material UI.
- Backend built with Golang, providing endpoints for CRUD operations.
- MongoDB used as the database to store tasks.

## Features
- List all tasks
- Add a new task
- Update a task 
- Delete a task

## Technologies Used
**Backend**
- Golang
- Gorilla Mux
- MongoDB
  
**Frontend**
- React
- TypeScript
- Material UI


## Getting Started
To run the project locally:

### Frontend
- Navigate to the frontend directory.
- Run `npm install` to install dependencies.
- Run `npm start` to start the frontend server.

### Backend
- Navigate to the backend directory.
- Run `go run main.go` to start the backend server.

### API Endpoints
- `GET /tasks`: Get all tasks
- `POST /tasks`: Create a new task
- `PUT /tasks/{id}`: Update a task by ID
- `DELETE /tasks/{id}`: Delete a task by ID
