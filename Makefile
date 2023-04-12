build:
	docker-compose build

up:
	docker compose run --rm api

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
