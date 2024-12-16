Гененрация кода grpc

protoc --proto_path=..\..\api --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ..\..\api\auth.proto
