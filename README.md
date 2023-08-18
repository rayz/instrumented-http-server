# Start Server
`make run`


# API

localhost:8080

GET / 

- returns all ToDos with id, description, completion status

POST /add

JSON Body
{
    "description": "..."
}

PUT /complete/{id}
- updates task with id to completed


# Go Lessons Learned
Range loop copies value. Need to index into a slice and get pointer in order to modify original value
