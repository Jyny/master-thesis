server/main: anc estimat-shift/main
	cd server && go build src/main.go
	DB="postgresql://postgres:8d137db2e7f8@127.0.0.1:5432/postgres" server/main

estimat-shift/main:
	cd estimat-shift && go build src/main.go

estmanc/main: anc estimat-shift/main
	cd estmanc && go build src/main.go

.PHONY: anc
anc:
	docker build -f anc/Dockerfile -t anc anc

.PHONY: db
db:
	mkdir -p db
	docker run --name pg-thesis --rm -d -p 127.0.0.1:5432:5432 -v ${PWD}/db:/var/lib/postgresql/data -e PGDATA=/var/lib/postgresql/data/pgdata -e POSTGRES_PASSWORD=8d137db2e7f8 postgres




.PHONY: client
client: aes
	cd client && pipenv run python client.py

.PHONY: pi
pi: aes
	python3 pi/pi.py

aes:
	cd server && go build aes/main.go
	cp server/main ./aes



.PHONY: clean
clean:
