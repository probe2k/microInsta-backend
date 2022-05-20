# INSTA_API

*Project by Yash Singh (probe2k)*

This is a GoLang REST API to handle said functions in a MicroInstagram environment. It uses MongoDB as the storage base.

## Dependencies

> mongo-driver
>
> godotenv

*Modules are ported using go's inbuilt git handler*

## API Endpoints Utilized

> /api/posts

> /api/posts/<PID>

> /api/posts/users/<id>?limit={}&lastid={}

> /api/users

> /api/users/<_id>

## Test Cases

*The said test cases can be implemented for each of the said API functionalities as :*

> go test -v

## Running

*Setup this directory at the desired location, followed by configuring go modules. Then run as follows :*

> go build main.go

> go run ./