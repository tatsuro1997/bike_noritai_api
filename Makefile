build:
	docker-compose build

up:
	make down
	docker compose up api db

stop:
	docker-compose stop

down:
	docker-compose down

bash:
	docker-compose exec api /bin/bash

db:
	docker exec -it db bash

fmt:
	docker compose run --rm api gofmt -l -s -w .

go_test:
	ENV=test go test -v ./test/...

test_db:
	make down
	docker compose up test_db
