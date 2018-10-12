.DEFAULT_GOAL := run

.PHONY: compose
compose:
	@docker-compose up -d

.PHONY: run
run:
	@go install
	@ozz

.PHONY: rule.import
rule.import:
	@oathkeeper rules import --endpoint=http://localhost:4456 config/oathkeeper/healthcheck.json

.PHONY: hydra.clients
hydra.clients:
	@hydra clients import --endpoint=http://localhost:4445 config/hydra/app.json
	@hydra clients import --endpoint=http://localhost:4445 config/hydra/oathkeeper.json
