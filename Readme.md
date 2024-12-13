
# Welcome 🖐️ To My Go backend API with MongoDB

This project is a backend API built using Golang, leveraging MongoDB for efficient data storage. Designed for scalability and performance, it offers robust RESTful endpoints to support modern web and mobile applications.


## startup

1. clone repo using 

```bash
 https://github.com/nitinthakurdev/app-todo-backend.git
```

2. run backend using

```bash
 go run src/main.go
```


## API Reference

#### 1. user registration API - POST

```bash
   /api/v1/user/create
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your username |
| `email`    | `string` | **Required**. Your email |
| `password` | `string` | **Required**. Your password |

