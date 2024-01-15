docker-compose-test: docker-compose.yml
	docker-compose build
	docker-compose up -d

docker-compose: docker-compose.yml
	docker-compose build
	docker-compose up

test: docker-compose-test
	docker exec grpc-example-server-1 go test -v -coverprofile=c.out ./...