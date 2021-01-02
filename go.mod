module github.com/Ripolak/minict

go 1.15

require (
    github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/genproto v0.0.0-20200722002428-88e341933a54 // indirect
	google.golang.org/grpc v1.30.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/apex/log v1.9.0
	github.com/containers/image v3.0.2+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/urfave/cli v1.22.5
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20201231184435-2d18734c6014 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.26.0
replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
