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


.PHONY: dockerdev
dockerdev:
	docker build . -t ak -f deployments/docker/dev/Dockerfile

.PHONY: dockercompose
dockercompose:
	rm -rf tmp/postgres-data
	mkdir -p tmp/postgres-data
	chmod -R 777 tmp
	cp scripts/sql/docker/setup.sh tmp/sql
	cp scripts/sql/docker/2023060800-initial-setup.sql tmp/sql
	docker-compose -f deployments/docker/docker-compose.yml up --build # --abort-on-container-exit --remove-orphans

.PHONY: dockercomposedown
dockercomposedown:
	docker-compose -f deployments/docker/docker-compose.yml down

.PHONY: dockerresetpg
dockerrestpg:
	sudo rm -rf tmp/postgres-data

.PHONY: dockerpsql
dockerpsql:
	sudo rm -rf ./tmp/postgres-data
	docker-compose exec app psql -U ak -h pg -d pgc -d

.PHONY: pgdockershell
pgdockershell:
	docker exec -it docker-pg-1 bash

