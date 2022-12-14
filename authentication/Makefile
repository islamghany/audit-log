

#-------------------------------- Database


# Database Container
.PHONY: db/container/run
db/container/run:
	@echo 'Running the postgreSQL container'
	docker run --name audit_log_auth_db -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_DB=auth -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

.PHONY: db/container/stop
db/container/stop:
	@echo 'Stopping the postgreSQL container'
	docker stop audit_log_auth_db

.PHONY:db/container/start
db/container/start:
	@echo 'Starting the postgreSQL container'
	docker start audit_log_auth_db

.PHONY:db/container/remove
db/container/remove:
	@echo 'Removeing the postgreSQL container'
	docker container rm audit_log_auth_db


# Database migrations 
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./db/migrations ${name}

