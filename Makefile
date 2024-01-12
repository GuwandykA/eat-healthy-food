# Dependencies
dependencies:
	docker-compose -f ./deployment/dependencies/docker-compose.yaml up -d

# Database
create-migration:
	migrate create -ext sql -dir db/migrations -tz Local init

# --- for Macos
#migrate:
#	migrate -database 'postgres://course_admin:course_password@localhost:5432/course_db?sslmode=disable' -path db/migrations up
#migrate-rollback:
#	migrate -database 'postgres://course_admin:course_password@localhost:5432/course_db?sslmode=disable' -path db/migrations down

# ---- for linux
migrate:
	migrate -database 'postgresql://course_admin:course_password@localhost:5454/course_db?sslmode=disable' -path db/migrations up
migrate-rollback:
	migrate -database 'postgresql://course_admin:course_password@localhost:5454/course_db?sslmode=disable' -path db/migrations down
