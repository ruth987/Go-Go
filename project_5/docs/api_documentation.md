# Task Manager API Documentation

## Endpoints

### GET /tasks
Returns all tasks

### GET /tasks/:id
Returns a specific task by ID

### POST /tasks
Creates a new task
Request body:

```json
{
"title": "string",
"description": "string",
"due_date": "datetime",
"status": "string"
}
```

### PUT /tasks/:id
Updates an existing task

### DELETE /tasks/:id
Deletes a task