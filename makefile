default:
	@echo "=============building Local DEV API============="
	docker-compose up -d --build
	docker-compose logs -f
	

up: default

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	# do something here

clean: down
	@echo "=============cleaning up============="
	docker container prune --force
	docker rmi score_averager