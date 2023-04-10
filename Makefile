build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

bash:
	docker-compose exec api /bin/bash

db:
	docker exec -it db bash

fmt:
	docker-compose exec api gofmt -l -s -w .
