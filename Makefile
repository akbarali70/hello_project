run:
	go run .

dev:
	~/go/bin/air
	
fmt:
	go fmt ./...

seed:
	curl -sS -X POST http://localhost:8001/seed/users
	