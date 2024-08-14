Task Management API Documentation

## Endpoints

### Get All Tasks
- URL: `http://localhost:3000/tasks`
- Method: GET
- Description: Retrieves a list of all tasks.
- Response:
    - 200 OK: List of tasks.

### Get Task By ID
- URL: `http://localhost:3000/tasks/:id`
- Method: GET
- Description: Retrieves a task by its ID.
- URL Parameters:
    - id: The ID of the task.
- Response:
    - 200 OK: The requested task.
    - 404 Not Found: Task not found.

### Create Task
- URL: `http://localhost:3000/tasks`
- Method: POST
- Description: Creates a new task with the specified details.
- Request Body:
    ```json
    {
        "title": "string",
        "description": "string",
        "completed": true/false
    }
    ```
- Response:
    - 201 Created: Task created successfully.
    - 400 Bad Request: Invalid input.

### Update Task
- URL: `http://localhost:3000/tasks/:id`
- Method: PUT
- Description: Updates an existing task by its ID.
- URL Parameters:
    - id: The ID of the task to be updated.
- Request Body:
    ```json
    {
        "title": "string",
        "description": "string",
        "completed": true/false
    }
    ```
- Response:
    - 200 OK: Task updated successfully.
    - 404 Not Found: Task not found.

### Delete Task
- URL: `http://localhost:3000/tasks/:id`
- Method: DELETE
- Description: Deletes a task by its ID.
- URL Parameters:
    - id: The ID of the task to be deleted.
- Response:
    - 200 OK: Task deleted successfully.
    - 404 Not Found: Task not found.

### User Login
- URL: `http://localhost:3000/login`
- Method: POST
- Description: Logs in a user with the provided username and password.
- Request Body:
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
- Response:
    - 200 OK: Login successful.
    - 401 Unauthorized: Invalid credentials.

### User Registration
- URL: `http://localhost:3000/register`
- Method: POST
- Description: Registers a new user with the provided username and password.
- Request Body:
    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```
- Response:
    - 201 Created: Registration successful.
    - 400 Bad Request: Invalid input.

### Promote User
- URL: `http://localhost:3000/promote/:id`
- Method: POST
- Description: Promotes a user to a higher role or permission level.
- URL Parameters:
    - id: The ID of the user to be promoted.
- Response:
    - 200 OK: User promoted successfully.
    - 404 Not Found: User not found.

for the tasks https://documenter.getpostman.com/view/32072638/2sA3s4mAah
for the user https://documenter.getpostman.com/view/32072638/2sA3s4mAjW
