up:
	docker-compose -f docker-compose.yaml down -v
	docker-compose -f docker-compose.yaml up signal-server
	docker-compose -f docker-compose.yaml up -d client
	docker-compose -f docker-compose.yaml ps

down:
	docker-compose -f docker-compose.yaml down -v
