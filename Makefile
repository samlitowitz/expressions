lint: docker-compose.yaml
	docker-compose run --rm lint

test: docker-compose.yaml
	docker-compose run --rm --build test
	docker-compose down
