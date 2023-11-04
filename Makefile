.PHONY: conn
conn:
	@echo "Connecting to postgres..."
	@PGPASSWORD='password' pgcli -h localhost -p 5432 -U postgres -d postgres
