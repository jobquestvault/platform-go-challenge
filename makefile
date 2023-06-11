app = ak

.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build ./...

.PHONY: run
run:
	go run ./cmd/$(app)/main.go -log-level=debug -server-host=localhost -server-port=8080 -db-host=localhost -db-port=5432 -db-username=ak -db-password=ak -db-schema=ak -db-name=pgc -db-ssl=false

.PHONY: docker-compose
docker-compose:
	docker-compose -f ./deployments/doccker/docker-compose.yml up
