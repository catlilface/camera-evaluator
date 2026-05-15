init:
	cp .env.sample .env
	cp .env ./evaluator

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f
