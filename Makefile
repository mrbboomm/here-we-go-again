colima_start:
	colima start --arch x86_64 --memory 8

docker_start:
	docker-compose up -d oracle kafka zookeeper