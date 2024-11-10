# task-management-api
STEP 1:
create mysql schema name task_management_api

STEP 2:
set up .env follow by .env.example

STEP 3:
run file migrate at ./internal/dbsql/migration/goose/main.go
with command 
 - go run ./internal/dbsql/migration/goose/main.go 
 - cd ./internal/dbsql/migration/goose && go run main.go

STEP 4: 
run this project 
with command 
 - go run ./cmd/main.go
 - cd cmd && go run main.go

STEP 5: 
set swagger or postman with ./docs/swagger.json
 - crtl+c swagger.json then go to swagger editor crtl+v
 - crtl+c swagger.json go to postman import button then crtl+v

STEP 6:
when wanna call API please set Bearer Token but can set anything because just checked Bearer Token

# Design Relational Database for Delivery
https://drive.google.com/file/d/1CAM88xSiOEAxQAeyCOC0sjNfUoW-hh5l/view?usp=sharing