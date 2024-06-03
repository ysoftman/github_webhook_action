module main

go 1.22.1

require (
	github.com/go-resty/resty/v2 v2.12.0
	github.com/ysoftman/github_webhook_action v0.1.11
)

// local 에서 아직 개발중인 모듈을 테스트 할때만 사용
//replace github.com/ysoftman/github_webhook_action v0.1.11 => ../

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	google.golang.org/appengine/v2 v2.0.5 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)
