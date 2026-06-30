# todoListWeb

REST API for a simple Todo-list.

## Use

Base url: `http://localhost:8080`.

### New task
**POST** `/tasks`

Body:
```json
{
  "title": "string",
  "description": "string"
}
```

Success: `201 Created`  
Errors: `400`, `409`

---

### Get task
**GET** `/tasks/{title}`

Success: `200 OK`  
Error: `404`

---

### Get all tasks
**GET** `/tasks`

Success: `200 OK`

---

### Get all uncompleted tasks
**GET** `/tasks?completed=false`

Success: `200 OK`

---

### Mark completed/uncompleted
**PATCH** `/tasks/{title}`

Body:
```json
{
  "complete": true
}
```

Success: `200 OK`  
Errors: `400`, `404`

---

### Delete task
**DELETE** `/tasks/{title}`

Success: `204 No Content`  
Error: `404`

## Curl example

```bash
curl -X POST "http://localhost:8080/tasks" \
  -H "Content-Type: application/json" \
  -d '{"title":"task1","description":"desc"}'
```
```bash
curl -X GET "http://localhost:8080/tasks/task1"
```
```bash
curl -X GET "http://localhost:8080/tasks"
```
```bash
curl -X GET "http://localhost:8080/tasks?completed=false"
```
```bash
curl -X PATCH "http://localhost:8080/tasks/task1" \
  -H "Content-Type: application/json" \
  -d '{"complete":true}'
```
```bash
curl -X DELETE "http://localhost:8080/tasks/task1"
``` 