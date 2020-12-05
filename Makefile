ifeq ($(strip $(PGHOST)),)
PGHOST=localhost
endif

ifeq ($(strip $(PGUSER)),)
PGUSER=postgres
endif

ifeq ($(strip $(PGPASS)),)
PGPASS=postgres
endif

ifeq ($(strip $(PGSSL)),)
PGSSL=disable
endif

ifeq ($(strip $(PGHOST)),)
PGHOST=localhost
endif

PIE=-buildmode=pie
TRFLAG=
ifneq ($(RACE),)
PIE=
TRFLAG=-race -cpu 1,2,4
export GORACE="halt_on_error=1"
endif


install:
	go install $(PIE) $(BRFLAG)

clean_test_db: 
	GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
	migrate -source file://db/postgres/migrations -database postgres://$(PGUSER):$(PGPASS)@$(PGHOST)/captable?sslmode=$(PGSSL) down -all
	migrate -source file://db/postgres/migrations -database postgres://$(PGUSER):$(PGPASS)@$(PGHOST)/captable?sslmode=$(PGSSL) up

test: lint clean_test_db
	go test -count=1 $(TRFLAG) ./...

update_db:
	GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
	migrate -source file://db/postgres/migrations -database postgres://$(PGUSER):$(PGPASS)@$(PGHOST)/captable?sslmode=$(PGSSL) up

down_db:
	GO111MODULE=off go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
	migrate -source file://db/postgres/migrations -database postgres://$(PGUSER):$(PGPASS)@$(PGHOST)/captable?sslmode=$(PGSSL) down 
