# Makefile

# Set env and Run the server using air (hot reload)
local:
	GO_ENV=local air

staging:
	GO_ENV=staging air
# Run normally without reload
run:
	GO_ENV=production go run main.go
