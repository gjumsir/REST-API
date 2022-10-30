module example/web-client

replace example/web-service-gin => ../web-service-gin

go 1.19

require example/web-service-gin v0.0.0-00010101000000-000000000000

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5 // indirect
)
