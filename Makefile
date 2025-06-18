run:
	go run --race main.go

run-without-race:
	go run main.go

cont-up:
	docker compose up -d
