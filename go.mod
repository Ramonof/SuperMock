module SuperStub

go 1.23.3

replace google.golang.org/grpc => ./grpc-go-dynamic-service-fork

require (
	github.com/golang-migrate/migrate/v4 v4.18.2
	github.com/gorilla/mux v1.8.1
	github.com/jhump/protoreflect v1.17.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
)
