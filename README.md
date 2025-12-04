## air-verse for hot reload

1. <https://github.com/air-verse/air>
2. run air with `air -c .air.toml`

## swagger openapi

1. to update the documentation, head to cmd/api
2. move to `base-entity.go` comment `DeletedAt` and uncomment the below implementation
3. run this command `swag init`
4. after success, revert back the change of commenting `DeletedAt`

## .air.toml

1. current working is for linux because we are using docker for running this apps
   bin = "./bin/api"
   cmd = "go build -o ./bin/api ./cmd/api/"

2. if you running locally, change .air.toml line 7-8 to:
   bin = "./bin/api.exe"
   cmd = "go build -o ./bin/ ./cmd/api/"

## Open API

1. to update the documentation, remove file `/cmd/api/docs/docs.go`
2. move to folder `/cmd/api` then perform this command `swag init`
