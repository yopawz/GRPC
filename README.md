# GRPC

To run this program you must run the commands below:

1. source .env (because i using .env as the information for the database)
2. if you want to regenerate the protobuf you can use protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    [packagename]/[filename].proto
3. finally to run the program you can run, go run server/server.go
