# Enviroment variables
GOCMD=go
PATHBIN=./bin/
PROJECTNAME=mutant
BINARY=$(PATHBIN)$(PROJECTNAME)
MAIN=cmd/main.go

build:
	$(GOCMD) build -o $(BINARY) $(MAIN)

# Run project 
run:
	$(GOCMD) run $(MAIN)

# Run covert test
coverage: 
	$(GOCMD) test -coverprofile=./coverage/cover.out ./...

cover-html:
	$(GOCMD) test -coverprofile=./coverage/cover.out ./...
	$(GOCMD) tool cover -html=./coverage/cover.out -o ./coverage/cover.html 

test:	
	$(GOCMD) test ./... -cover	

swag:
	swag init --generalInfo  $(MAIN)