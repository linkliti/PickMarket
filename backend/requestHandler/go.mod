module pickmarket/requestHandler

go 1.22.0

require (
	github.com/gorilla/handlers v1.5.2
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	pmutils v0.0.0-00010101000000-000000000000
	protos v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gorilla/mux v1.8.1
)

replace pmutils => ../pmutils

replace protos => ../protos
