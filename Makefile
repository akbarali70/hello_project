run:
	go run .

dev:
	air

fmt:
	go fmt ./...

seed:
	curl -X POST http://localhost:8001/seed/users