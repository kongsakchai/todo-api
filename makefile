PHONY: dev-start
dev-start:
	@docker compose -f docker-compose.dev.yml up -d

PHONY: dev-stop
dev-stop:
	@docker compose -f docker-compose.dev.yml down

PHONY: start
start:
	@docker compose up -d

PHONY: stop
stop:
	@docker compose down
