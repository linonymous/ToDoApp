##### To create models
protoc --go_out=. --go_opt=paths=source_relative todo/todo.proto
##### To create server & clients
protoc --go-grpc_out=. todo/todo.proto