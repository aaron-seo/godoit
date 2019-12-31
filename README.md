# godoit (wip)

Godoit is a task scheduler for UCLA students with no need for manual work to setup the quarter's schedule.

## Table of Contents

- [Technologies Used](https://github.com/aaron-seo/godoit#technologies-used)
- [Structure](https://github.com/aaron-seo/godoit#structure)
- [API](https://github.com/aaron-seo/godoit#api)
- [Getting Started](https://github.com/aaron-seo/godoit#getting-started)
- [Goals](https://github.com/aaron-seo#goals)

## Technologies Used

- [Go](https://www.golang.org)
- [go-chi](https://github.com/go-chi/chi)
- [lib-pq](https://github.com/lib/pq)
- [Docker](https://www.docker.com)

## Structure
```
.
├── Dockerfile
├── GoDoIt
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── tasks
│   ├── handlers.go
│   └── model.go
└── users
    ├── handlers.go
    └── model.go

2 directories, 11 files
```

## API

```
GET /tasks/
/tasks/
/tasks/
```

## Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/aaron-seo/godoit.git
cd godoit
```

### 2. Build docker images
```bash
docker-compose build
```

### 3. Run it
```bash
docker-compose up
```

## Goals
- Actually finish the API lmao
- Develop a React frontend
- Deploy to Kubernetes
