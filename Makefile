.DEFAULT_GOAL := compile
.PHONY: build
all: clean swagger check build  run
check:
	@echo "Checking project";
	sh scripts/check.sh

check-ci:
	@echo "Checking project";
	sh scripts/check_ci.sh

build:
	@echo "Building project";
	sh scripts/build.sh

install:
	@echo "Install dependencies";
	sh scripts/install.sh

swagger:
	@echo "Updating Swagger"
	swag init  -g cmd/wotracker-back/main.go

run:
	@echo "Running backend";
	sh scripts/run.sh

clean:
	@echo "Cleaning stuff";
	sh scripts/clean.sh

db:
	@echo "Starting DB";
	sh scripts/start_db.sh