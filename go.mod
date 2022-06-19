module github.com/zackattackz/renderer

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/micro/v3 v3.10.1
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
