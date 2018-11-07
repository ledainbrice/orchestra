PID      = /tmp/awesome-golang-project.pid
GO_FILES = $(wildcard *.go)
APP      = ./app
serve: restart
	@fswatch -o . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@echo "actually do nothing"
$(APP): $(GO_FILES)
	@go build $? -o $@
restart: kill before $(APP)
        @app & echo $$! > $(PID)

.PHONY: serve restart kill before # let's go to reserve rules names