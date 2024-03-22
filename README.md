# maas
Memes-as-a-Service

## Introduction
This project is an example of a simple REST microservice in go.  It is intended to provide "memes as a service".

## Using this code
This repo is published under an [MIT License](LICENSE) which tends to be flexible for the user while still protecting the creator.  Please feel free to use this code and also to submit Pull Requests with any fixes/improvements/suggestions/etc.  Also, if you use code from the repo I request (but do not require) that you star the repo in GitHub.

All code in this repo was written by me, but much of the code is actually reused from other projects I've done including:
- [go-server](https://github.com/NathanBak/go-server)
- [easy-cass-go](https://github.com/NathanBak/easy-cass-go)
- [cfgbuild](https://github.com/NathanBak/cfgbuild)

## Running the code

### Setup development environment
The server uses properties stored in a `.env` file.  To generate said file, from the project root run `scripts/create_env.sh`.  If you want to connect to the [DataStax Astra](https://astra.datastax.com/) database, you will need to provide a token, otherwise in-memory storage will be used.

### Running the Server
From the project root run `go run .`

### Testing the Server
Here are some example CURL commands that can be used to access the running server:
- ```curl --header "Authorization: aeca88d8-af10-4316-99cc-4ae3497ce8d0" http://localhost:8081/api/v1/memes?lat=40.730610\&lon=-73.935242\&query=food```
- ```curl --header "Authorization: aeca88d8-af10-4316-99cc-4ae3497ce8d0" http://localhost:8081/api/v1/councounts/```

### Prerequisites
- Go (see the [go.mod](https://github.com/NathanBak/maas/blob/7d2bc13307366e54e18d85cf93e705df088697d2/go.mod#L3) file for the correct version) 
- Git
- Database token (optional)