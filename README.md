# Mutant
API Rest that determines whether a DNA sequence is from a mutant or a human

# Install Go
Download and install Go, obtain from https://go.dev/learn/

## Set GOPATH to path in bash file
```
vi source ~/.zshrc
Add this line (depends on the location of the Go binary)
export GOPATH=$HOME/Documents/go
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
source ~/.zshrc
```
# Postgres
Download Postgres from https://www.postgresql.org/download/
You must create and configure a database in Postgres.

# Create folder for Go files
```
mkdir -p ~/Documents/go/src/github.com/wilmerpino/
```
# Get source code from Github
```
cd ~/Documents/go/src/github.com/wilmerpino/
git clone https://github.com/wilmerpino/mutant.git
```
# Install package
```
cd mutant
go mod tidy
```
# Configure enviroment variables
Open file public/config.go and update yours enviroment variables
```
const (
	AppName    = "Mutant API Rest"
	LevelLog   = "debug"
	PortServer = "8088"
	DBHost     = "localhost"
	DBUser     = "postgres"
	DBName     = "mutant"
	DBPassword = ""
	DBPort     = "5432"
    SwaggerHost = "http://localhost:8088"
)
```

# Run you Mutant App
```
go run cmd/main.go 
or
make run
```

# Build your Mutant App
```
go build -o ./bin/mutant ./cmd/main.go 
or
make build
```

# Refresh Swagger Docs
```
swag init --generalInfo  cmd/main.go
or
make swag

```

# Docker
(Pending)

# Endpoint
## Check service
Check if the service is active
```
Method: GET
Path: /healthcheck
Response: 200 - OK
{
    "status": "OK"
}
```

## Check mutant DNA
Checks whether the DNA strand is from a mutant or a human
```
Methos: POST
Path: /mutant
Body: 
{
    "dna": [
        "ATGCGA",
        "CAGTGC",
        "TAATGT",
        "AGATTG",
        "CCACTA",
        "TCCCTG"
    ]
}
Response: 200 - OK
{
    "message": "DNA_HUMAN"
}
or
Response: 403 - Forbidden
{
    "message": "DNA_HUMAN"
}
or
Response: 400 - Bad Request
{
    "message": "STRAND_LENGTH_INVALID"
}
or
{
    "message": "SOME_CHARACTER_NOT_VALID"
}
```
# Statistics
Shows the statistics of the amount of DNA checked, how many are Human and how many are mutants, as well as the human/mutant ratio.
```
Methos: POST
Path: /stats
Response: 200 - Ok
{
    "count_mutant_dna": 3,
    "count_human_dna": 2,
    "ratio": 1.5
}
```

# Test
## Run coverage test
```
go test ./... -cover
or
make test
```
## Run html coverage test
```
go test -coverprofile=./coverage/cover.out ./... 
go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
or
make cover-html
```
Find coverage test in http://localhost:8088/coverage
# Postman Collector
You cant find the postman collector in public/postman/Mutant.postman_collection.json

# Contract Swagger
You cant find the contract for mutants endpoints in http://localhost:8088/swagger/index.html

# Apache JMeter
You cant find the Load Test file for JMeter in public/jmeter/MELI- Mutant.jmx

# Heroku
The application is deployed on the Heroku cloud application service.

Service: https://wp-mutant.herokuapp.com

Contract: https://wp-mutant.herokuapp.com/swagger/index.html

Coverage Test: https://wp-mutant.herokuapp.com/coverage


---
Wilmer Pino
Software Development Engineer
Github: https://github.com/wilmerpino
LinkedIn: https://www.linkedin.com/in/wilmerpino/