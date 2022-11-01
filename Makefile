# Prefer running recipes over files with same names
.PHONY: start

help:
	@echo
	@echo "Authware"
	@echo
	@echo "  Commands: "
	@echo
	@echo "    help - Show this message."
	@echo "    start - Start all services."

start:
	go run main.go
