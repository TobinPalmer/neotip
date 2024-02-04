GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-starter-app

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

run:
	go install . && heroku local web --port 5001

pull-db:
	psql $(shell whoami) -c 'DROP DATABASE "go-starter-db";'
	heroku pg:pull DATABASE_URL go-starter-db -a go-starter-app

push-db:
	heroku pg:reset --confirm go-starter-app -a go-starter-app
	heroku pg:push go-starter-db DATABASE_URL -a go-starter-app

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
