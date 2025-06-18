## go-graceful-with-context

Implement HttpServer with context graceful shutdown server in golang.

---

#### **_How to run this project_** 🏃‍♂️

```bash

### Run with makefile
make run

### Run with go-cli
go run --race main.go
```

---

#### Integration test with redis

```bash

### run container in docker compose (include redis)
make cont-up


### how to check data on redis ?
docker exec -it {your_container_name_or_id} bash

redis-cli

### data stamp at database 0 by default
HGETALL {KEY}

### Example Key
HGETALL ARTICLE|1

### curl api endpoint
curl -X POST http://localhost:2318/graceful \
     -H "Content-Type: application/json" \
     -d '{
           "author": "terry davis",
           "title": "about my temple os",
           "text": "holy c"
         }'


```
