all:
	pnpm start

.PHONY: backend
backend:
	go run ./backend/main.go
