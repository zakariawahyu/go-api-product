docker-compose:
	@echo Starting docker compose
	docker-compose -f docker-compose.yaml up -d --build

local:
	@echo Starting local
	docker-compose -f docker-compose.local.yaml up -d

run:
	@echo Run service
	go run cmd/main.go