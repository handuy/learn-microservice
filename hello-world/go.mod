module learn-microservice/hello-world

go 1.13

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	go.etcd.io/etcd v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
)
