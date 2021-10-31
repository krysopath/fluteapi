test:
	go test -cover ./...

newmigration:
	migrate create -ext sql -dir db/migration -seq $(shell read -p "Enter a Name: " word; echo $$word)

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
