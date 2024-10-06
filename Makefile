DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456789sq
DB_DBNAME=postgres
DB_SSLMODE=disable

run:
		goose -dir migrations postgres "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" up
		go run cmd/main.go
