build: 
	@go build -o bin/arbitrax-api

run:
	@./bin/arbitrax-api

dev: 	
	@cd cmd/api && go build -o ../../bin/arbitrax-api
	@./bin/arbitrax-api

test:
	@go test -v ./... | grep -v '\[no test files\]' 
	

start-db: 
	@docker compose up -d

stop-db:
	@docker compose down 

exec-db:
	@docker exec -it arb-postgres psql -U postgres
