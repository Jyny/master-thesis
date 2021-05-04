.PHONY: run
run:
	DB="postgresql://postgres:8d137db2e7f8@127.0.0.1:5432/postgres" go run src/main.go

.PHONY: db
db:
	docker run --name pg-thesis --rm -d -p 5432:5432 -e POSTGRES_PASSWORD=8d137db2e7f8 postgres

.PHONY: clean
clean:
	docker stop pg-thesis
	git clean -Xf .
