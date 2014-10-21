
# Export GOPATH
Should be set to project directory to make everything simple

export GOPATH=/Users/antonholmquist/Projects/gotest2


# Run
go run app.go

# Format code (fixes indentation and stuff)
go fmt

# Dependencies

## Gorilla Mux
https://github.com/gorilla/mux
URL router

## Goose
https://bitbucket.org/liamstask/goose
Database migrations

# Add dependency

go get github.com/jmoiron/sqlx

Then remove .git add commit everything in.


# Migrations

./bin/goose create <name>
./bin/goose up
./bin/goose down

# Heroku

## 1. Make sure we have buildback
heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
