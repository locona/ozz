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
