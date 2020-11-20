### Hello World microservice

Lưu ý phải chạy lệnh ```go mod init``` ở thư mục root của project (cùng cấp 2 thư mục client và server)

Lưu ý nữa là file go.mod phải có dòng sau:
```go
replace (
	github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
)
```