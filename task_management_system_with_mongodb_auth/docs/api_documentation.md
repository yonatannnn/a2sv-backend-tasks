# Task Management API

## Endpoints

### GET /tasks
- Retrieves all tasks.

### GET /tasks/:id
- Retrieves a task by ID.

### POST /tasks
- Creates a new task.
- Body: `{ "title": "Task Title", "description": "Task Description", "completed": false }`

### PUT /tasks/:id
- Updates a task by ID.
- Body: `{ "title": "Updated Title", "description": "Updated Description", "completed": true }`

### DELETE /tasks/:id
- Deletes a task by ID.

## User Authentication

### POST /auth/register
- Registers a new user.
- Body: `{ "username": "example", "password": "password123" }`

### POST /auth/login
- Logs in a user.
- Body: `{ "username": "example", "password": "password123" }`



for the tasks
https://documenter.getpostman.com/view/32072638/2sA3s4mAah

for the user
https://documenter.getpostman.com/view/32072638/2sA3s4mAjW

















## Endpoints

### GET /tasks
- Retrieves all tasks.

### GET /tasks/:id
- Retrieves a task by ID.

### POST /tasks
- Creates a new task.
- Body: `{ "title": "Task Title", "description": "Task Description", "completed": false }`

### PUT /tasks/:id
- Updates a task by ID.
- Body: `{ "title": "Updated Title", "description": "Updated Description", "completed": true }`

### DELETE /tasks/:id
- Deletes a task by ID.

