Pre-reqs
Docker, Minikube, K8s, Helm


# Setup
Build server image

`docker build --rm -t server .`

Load image into minikube

`minikube image load server`

Install helm chart

https://docs.datadoghq.com/containers/kubernetes/installation/?tab=helm


Deploy server

`k apply --filename deployment.yaml `

Enable portforwarding to send requests to pod from localhost

`k port-forward <POD_NAME> 8080:8080`


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

POST /delete/{id}


# Go Lessons Learned
Range loop copies value. Need to index into a slice and get pointer in order to modify original value
