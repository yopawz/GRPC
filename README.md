# GRPC

To run this program you must run the commands below:
(make sure you have a database with a table named employee and have these fields (id, name, age, salary) )
1. source .env (because i using .env as the information for the database)
2. if you want to regenerate the protobuf you can use protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    [packagename]/[filename].proto
3. finally to run the program you can run, go run server/server.go (grpc)
4. go run gql_server/gql_server.go (grpc server for gql)
5. go run gateway/main.go (grpc gateway for gql)
