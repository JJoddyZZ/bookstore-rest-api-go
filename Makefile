all: run

# --- local deploy ---

mkdir-output:
	@mkdir -p output

build: mkdir-output
	@clear
	@echo "Building executable..."
	@echo
	@go build -o output/bookstore cmd/bookstore/main.go
	@echo "Generated output/bookstore"
	@echo

run: build
	@clear
	@echo "Running server..."
	@echo
	@./output/bookstore || (echo "terminated with: $$?")
	@echo

# requires github.com/vektra/mockery
gen-mocks:
	@clear
	@echo "Generating mocks..."
	@echo
	mockery --all --dir=internal/controllers --output=internal/controllers/mocks
	mockery --all --dir=internal/repositories --output=internal/repositories/mocks
	mockery --all --dir=internal/services --output=internal/services/mocks
	@echo
	@echo "Mocks generated"
	@echo

# --- docker deploy ---

docker-run:
	@clear
	@echo "Running server in container..."
	@echo
	docker-compose -f docker-compose.yaml build bookstore-api
	docker-compose -f docker-compose.yaml up

docker-down:
	@clear
	@echo "Running server in container..."
	@echo
	docker-compose -f docker-compose.yaml down

docker-reload: docker-down docker-run

# To fix 'No space left on device' docker issue
docker-prune:
	@clear
	@echo "Running docker clean up..."
	@echo
	@docker system prune -f
	@docker volume prune -f

# --- testing feats ---

test: mkdir-output
	@clear
	@echo "Running tests..."
	@echo
	@go test -coverprofile output/cover.out ./...
	@echo
	@echo "Generated output/cover.out"
	@echo

# requires: github.com/nikolaydubina/go-cover-treemap
gen-heatmap: test
	@clear
	@echo "Generating coverage heatmap..."
	@echo
	@go-cover-treemap -coverprofile output/cover.out > output/heatmap.svg
	@echo "Generated output/heatmap.svg"
	@echo

# requires: github.com/gojek/go-coverage
show-coverage-table: test
	@clear
	@echo "Generating coverage table..."
	@echo
	@go-coverage -f output/cover.out --line-filter 0
	@echo

# requires github.com/loov/goda
gen-dependencies-graph:
	@clear
	@echo "Generating output/graph.svg..."
	@echo
	@goda graph ./...| dot -Tsvg -o output/graph.svg
	@echo "Generated output/graph.svg"
	@echo
