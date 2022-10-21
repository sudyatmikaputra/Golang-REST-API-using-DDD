## Requirement

1. Install and run PostgreSQL in local with port 5432 by exec this command 
`docker run --name postgres-docker -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 -d postgres`

2. Set secret key for "USER_DEFAULT_ROLE" with uuid value as default role id for user role. (example export USER_DEFAULT_ROLE=becdd3c3-6e9d-4fb9-9f05-6d183c87de16)

## How to run 

### To run the server:
`go run cmd/medicplus/main.go`

### Exposed port:
Listen to port 8001 by default and 9100 for profiler
