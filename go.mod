module hongbao

go 1.12

require (
	github.com/Shopify/sarama v1.25.0
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/jmoiron/sqlx v1.2.0
	github.com/kataras/iris v11.1.1+incompatible
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.6.0
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/wangsongyan/wblog v0.0.0-20190212134921-875ef9a6cda0
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5
	google.golang.org/grpc v1.26.0
	gopkg.in/ini.v1 v1.51.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190621222207-cc06ce4a13d4

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace golang.org/x/crypto v0.0.0-20190621222207-cc06ce4a13d4 => github.com/golang/crypto v0.0.0-20190621222207-cc06ce4a13d4
