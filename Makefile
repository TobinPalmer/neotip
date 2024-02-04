GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/neotip

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

run:
	go install . && heroku local web --port 5001

pull-db:
	psql $(shell whoami) -c 'DROP DATABASE "neotip-db";'
	heroku pg:pull DATABASE_URL neotip-db -a neotip

push-db:
	heroku pg:reset --confirm neotip -a neotip
	heroku pg:push neotip-db DATABASE_URL -a neotip

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
