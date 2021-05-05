.PHONY: server
server:
	mkdir -p upload 
	docker run --name anc-server -d --rm --network thesis -p 8080:8080 -v ${PWD}/upload:/workdir/upload -e DB="postgresql://postgres:8d137db2e7f8@pg-thesis:5432/postgres" anc-server /workdir/server


.PHONY: db
db:
	mkdir -p db
	docker run --name pg-thesis --rm -d --network thesis -p 5432:5432 -v ${PWD}/db:/var/lib/postgresql/data -e PGDATA=/var/lib/postgresql/data/pgdata -e POSTGRES_PASSWORD=8d137db2e7f8 postgres


.PHONY: network
network:
	docker network create thesis


.PHONY: build
build:
	docker build -t anc-server . 


.PHONY: kill
kill:
	docker stop anc-server
	docker stop pg-thesis
	docker network rm thesis 


.PHONY: clean
clean: kill
	git clean -Xf .
