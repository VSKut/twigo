![Go version](https://img.shields.io/github/go-mod/go-version/vskut/twigo)
[![Build Status](https://travis-ci.org/vskut/twigo.svg)](https://travis-ci.org/vskut/twigo)
[![Go Report Card](https://goreportcard.com/badge/github.com/vskut/twigo)](https://goreportcard.com/report/github.com/vskut/twigo)
[![GolangCI Report card](https://golangci.com/badges/github.com/vskut/twigo.svg)](https://golangci.com/r/github.com/VSKut/twigo)
[![codecov.io](https://codecov.io/github/vskut/twigo/branch/master/graph/badge.svg)](https://codecov.io/github/vskut/twigo)
![Swagger](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FVSKut%2Ftwigo%2Fmaster%2Fapi%2Fswagger-spec%2Fapi.json)
![License](https://img.shields.io/github/license/vskut/twigo)
# Twigo APP

### Install
1. Set your ENV
        
        # Postgres
        DB_HOST=db
        DB_DRIVER=postgres
        DB_USER=postgres
        DB_PASSWORD=postgres
        DB_NAME=postgres
        DB_SSL_MODE=disable
        
        # gRPC server
        SERVER_HOST=server
        SERVER_PORT=8081
        
        # Rest gateway
        GATEWAY_HOST=
        GATEWAY_PORT=8080
        
        # JWT secret key
        JWT_SECRET=jwtSecretKey
        
2. `docker-compose -f deployments/docker-compose.yml up` or `make dockerize`
3. Endpoints hosted on http://GATEWAY_HOST:GATEWAY_PORT/
4. Open [Swagger](http://127.0.0.1:8082/)

### Endpoints
`POST /register` - create new account with specified nick(unique in app), email, and password
	
	Payload: 
		username - some user name
		email - user email address
		password - some password 
	
	Result: 
		id - primary key
		username - username which you specified in payload 
		email   - user email address which you specified in payload 

`POST /login` - accept email and password  and return token, uses JWT
	
	Payload:
		email - user email address
		password - some password 

	Result:
		token - jwt token
		
`POST /subscribe` - add account with login to your subscription list, you start seeing his tweets in your feeds 
	
	Payload: 
		nickname - nick name for account for which you want to subscribe 

`POST /tweets` - create a tweet, account id should be found from JWT

	Payload: 
		message - some tweet message

	Result:
		id - message primary key 
		message - tweet message

`GET /tweets` - return all tweets from your subscriptions 
	
	Result:
		tweets -  all tweets from your subscriptions 