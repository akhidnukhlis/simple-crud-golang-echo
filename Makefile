migration-create:
	migrate create -ext sql -dir migrations/sql/ $(name)

migration-up:
	#migrate -path migrations/sql -verbose -database "${DATABASE_URL}" up
	migrate -path migrations/sql -verbose -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" up

migration-down:
	#migrate -path migrations/sql -verbose -database "${DATABASE_URL}" down
	migrate -path migrations/sql -verbose -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" down

run-db:
	docker compose up postgres redis

run-api:
	go run ./

build-api:
	go build -v -o ./bin/ ./

#test:
#	go test -v -cover -benchmem ./...

mock:
	mockery --all

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

#docs:
#	swag i --dir ./cmd/api/,\
#	./modules/,\
#	./pkg/wrapper,\
#	./pkg/contexts
#
#git-hooks:
#	echo "Installing hooks..." && \
#	rm -rf .git/hooks/pre-commit && \
#	ln -s ../../tools/scripts/pre-commit.sh .git/hooks/pre-commit && \
#	chmod +x .git/hooks/pre-commit && \
#	echo "Done!"

routes:
	go run ./routes/

.PHONY: routes run-api run-db build-api migration-create docs