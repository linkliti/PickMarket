module itemsWorker

go 1.22.0

require (
	github.com/jackc/pgx/v5 v5.5.5
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	pmutils v0.0.0-00010101000000-000000000000
	protos v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace pmutils => ../pmutils

replace protos => ../protos
