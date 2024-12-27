up:
	@if [ ! -f .env ]; then \
        cp .env.example .env; \
    fi
	docker compose up -d

down:
	docker compose down

fresh:
	@if [ ! -f .env ]; then \
        cp .env.example .env; \
    fi
	docker compose down --remove-orphans
	docker compose build --no-cache
	docker compose up -d --build -V

stubs:
	cd proto && buf generate


dev: 
	go run github.com/joho/godotenv/cmd/godotenv@latest -f .env \
	go run github.com/air-verse/air@latest -c .air.toml run

migrate:
	go run github.com/joho/godotenv/cmd/godotenv@latest -f .env \
	go run ./... migrate

tidy:
	go mod tidy

logs:
	docker compose logs -f

test:
	go test -v -race -cover -count=1 -failfast ./...

lint:
	golangci-lint run -v

rename:
	find . -type f ! -path "./.git/*" ! -path "./build/*" ! -path "./Makefile" -exec sed -i.bak -e "s|github.com/ankorstore/yokai-grpc-template|github.com/$(to)|g" {} \;
	find . -type f -name "*.bak" -delete
