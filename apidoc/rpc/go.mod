module rpc

go 1.14

require (
	github.com/gmsec/micro v0.0.0-20210702092856-15091f6bcc57
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.29.1
