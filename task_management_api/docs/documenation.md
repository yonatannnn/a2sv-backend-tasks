## Endpoints

### Get all tasks
- **URL**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieves all tasks.

### Get a task by ID
- **URL**: `/tasks/:id`
- **Method**: `GET`
- **Description**: Retrieves a task by its ID.

### Create a new task
- **URL**: `/tasks`
- **Method**: `POST`
- **Description**: Creates a new task.
- **Body Parameters**:
    - `title` (string): The title of the task.
    - `description` (string): The description of the task.

### Update a task
- **URL**: `/tasks/:id`
- **Method**: `PUT`
- **Description**: Updates an existing task by its ID.
- **Body Parameters**:
    - `title` (string, optional): The new title of the task.
    - `description` (string, optional): The new description of the task.
    - `completed` (boolean, optional): The completion status of the task.

### Delete a task
- **URL**: `/tasks/:id`
- **Method**: `DELETE`
- **Description**: Deletes a task by its ID.

## Example Requests

### Get all tasks
`curl -X GET http://localhost:3000/tasks`

### Get a task by ID
`curl -X GET http://localhost:3000/tasks/1`

### Create a new task
`curl -X POST http://localhost:3000/tasks -d '{"title": "New Task", "description": "This is a new task."}' -H "Content-Type: application/json"`

### Update a task
`curl -X PUT http://localhost:3000/tasks/1 -d '{"title": "Updated Task", "description": "This is an updated task.", "completed": true}' -H "Content-Type: application/json"`

### Delete a task
`curl -X DELETE http://localhost:3000/tasks/1`


https://documenter.getpostman.com/view/32072638/2sA3s4mAah