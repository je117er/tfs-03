# Project info
* Reimplementation of an in-class exercise using RabbitMQ.
## Todo
* Integrate the [worker pattern](https://divan.dev/posts/go_concurrency_visualize/)
into the code so multiple workers can shoulder the workload.
## Directory structure
```bigquery
├── config
│   └── config.go
├── db
│   └── db.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── models.go
├── README.md
├── rmqClient
│   └── client.go
├── scheduler
│   └── scheduler.go
├── utils
│   └── utils.go
└── worker
    └── worker.go

```
