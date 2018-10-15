.DEFAULT_GOAL := run

.PHONY: compose
compose:
	@docker-compose up -d

.PHONY: client
client:
	@cd ./client & go build -o ./client/client ./client
	@./client/client

.PHONY: server
server:
	@cd ./server & go build -o ./server/server ./server
	@./server/server

.PHONY: oathkeeper.rules.import
oathkeeper.rules.import:
	-oathkeeper rules import --endpoint=http://localhost:4456 config/oathkeeper/healthcheck.json

.PHONY: hydra.clients.import
hydra.clients.import:
	-hydra clients import --endpoint=http://localhost:4445 config/hydra/app.json
	-hydra clients import --endpoint=http://localhost:4445 config/hydra/oathkeeper.json
