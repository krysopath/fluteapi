test:
	go test -cover ./...

migrateup:
	migrate -path db/migration \
		-database "$$DB_SOURCE"\
	   	-verbose\
	   	up

migratedown:
	. ./app.env && migrate -path db/migration \
		-database "$$DB_SOURCE"\
	   	-verbose\
	   	down

sqlc:
	sqlc generate
